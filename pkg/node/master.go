package node

import (
	"TaskWeaver/node/core"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type MasterAgent struct {
	Node
	api          *gin.Engine
	ln           net.Listener
	svr          *grpc.Server
	nodeSvr      *core.NodeServiceGrpcServer
	caPrivateKey *rsa.PrivateKey
}

var agent *MasterAgent

func NewMaster(addr string, api *gin.Engine) (*MasterAgent, error) {
	//
	if agent == nil {
		agent = &MasterAgent{}

		// initialize agent
		if err := agent.Init(); err != nil {
			panic(err)
		}
	}
	return agent, nil
}

func (ma *MasterAgent) Start() {
	// grpc server
	go ma.svr.Serve(ma.ln)
	// start api server ?
	// ---

	// CleanUp
	defer ma.svr.Stop()
}

func (ma *MasterAgent) Init() (err error) {
	// create listener
	ma.ln, err = net.Listen("tcp", ":"+ma.Config().Port())
	if err != nil {
		return err
	}

	ma.svr = grpc.NewServer()

	ma.nodeSvr = core.GetNodeServiceGrpcServer()
	core.RegisterNodeServiceServer(ma.svr, ma.nodeSvr)

	// init API
	go func() {
		ma.api = gin.Default()
		// register API setup elsewhere?
	}()

	return nil
}

func (n *MasterAgent) SetupSecurity() (err error) {
	if n == nil {
		return fmt.Errorf("Node is not initialized")
	}
	// Generate CA private key
	caPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Error generating CA private key:", err)
		return
	}

	// Create CA certificate template
	caTemplate := x509.Certificate{
		SerialNumber: big.NewInt(2023),
		Subject: pkix.Name{
			Organization: []string{""},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0),
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	caCertBytes, err := x509.CreateCertificate(rand.Reader, &caTemplate, &caTemplate, &caPrivateKey.PublicKey, caPrivateKey)
	if err != nil {
		fmt.Println("Error creating CA certificate:", err)
		return
	}

	caCertFile, err := os.Create("ca.crt")
	if err != nil {
		fmt.Println("Error creating CA certificate file:", err)
		return
	}
	pem.Encode(caCertFile, &pem.Block{Type: "CERTIFICATE", Bytes: caCertBytes})
	caCertFile.Close()

	masterPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Error generating master node private key:", err)
		return
	}

	masterTemplate := x509.Certificate{
		SerialNumber: big.NewInt(2024),
		Subject: pkix.Name{
			Organization: []string{"MyCompany Master Node"},
		},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().AddDate(1, 0, 0),
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}

	// Create master node certificate
	masterCertBytes, err := x509.CreateCertificate(rand.Reader, &masterTemplate, &caTemplate, &masterPrivateKey.PublicKey, caPrivateKey)
	if err != nil {
		fmt.Println("Error creating master node certificate:", err)
		return
	}

	// Write master node certificate to file
	masterCertFile, err := os.Create("master.crt")
	if err != nil {
		fmt.Println("Error creating master node certificate file:", err)
		return
	}
	pem.Encode(masterCertFile, &pem.Block{Type: "CERTIFICATE", Bytes: masterCertBytes})
	masterCertFile.Close()

	// Write master node private key to file
	masterKeyFile, err := os.Create("master.key")
	if err != nil {
		fmt.Println("Error creating master node private key file:", err)
		return
	}
	pem.Encode(masterKeyFile, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(masterPrivateKey)})
	masterKeyFile.Close()

	return nil

}

// process csr

func (n *MasterAgent) ProcessCSR(csrBytes []byte) (certBytes []byte, err error) {
	if n == nil {
		return nil, fmt.Errorf("Master Agent is not initialized")
	}

	// Load CA private key (for signing)
	caPrivateKey, err := os.Open("ca.key")
	if err != nil {
		log.Fatalf("Error opening CA private key: %v", err)
	}
	defer caPrivateKey.Close()
	// Get info to determine size
	fileInfo, err := caPrivateKey.Stat()

	// Load content into buffer
	var caPrivBuf := make([]byte, fileInfo.Size())
	_, err = caPrivateKey.Read(caPrivBuf)
	if err != nil && err != io.EOF {
		log.Fatalf("Error reading file: %v", err)
	}
	// Decode the PEM block
	caPrivateKeyBlock, _ := pem.Decode(caPrivBuf)
	caPrivateKeyParsed, err := x509.ParsePKCS1PrivateKey(caPrivateKeyBlock.Bytes)
	if err != nil {
		log.Fatalf("Error parsing CA private key: %v", err)
	}

	//
	csrBlock, _ := pem.Decode(csrBytes)
	csr, err := x509.ParseCertificateRequest(csrBlock.Bytes)
	if err != nil {
		log.Fatalf("Error parsing worker node CSR: %v", err)
	}

	// Create and sign a certificate based on the CSR
	template := x509.Certificate{
		SerialNumber: big.NewInt(time.Now().Unix()), // Unique serial number
		Subject:      csr.Subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(1, 0, 0), // Valid for 1 year
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}, // For client authentication
	}

	_certBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &csr.PublicKey, caPrivateKeyParsed)
	if err != nil {
		log.Fatalf("Error creating signed certificate: %v", err)
		return nil, err
	}

	return _certBytes, nil
}
