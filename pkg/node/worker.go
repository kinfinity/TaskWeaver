package node

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"taskWeaver/pkg/node"
	"taskWeaver/pkg/node/core"
	"taskWeaver/pkg/task"
	"time"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Worker struct {
}

type WorkerAgent struct {
	Node
	conn      *grpc.ClientConn
	c         core.NodeServiceClient
	Queue     queue.Queue
	Db        map[uuid.UUID]task.Task
	TaskCount int
}

func (w *Worker) CollectStats() {
	fmt.Println("Stats!")
}

var workerAgent *WorkerAgent

func GetWorkerAgent() *WorkerAgent {
	if workerAgent == nil {
		workerAgent = &WorkerAgent{}

		// initialize on Node
		if err := workerAgent.Init(); err != nil {
			panic(err)
		}
	}
	return workerAgent
}

func (wa *WorkerAgent) Init() (err error) {
	// Get Node Information
	opts, err := node.GetGrpcOpts()

	// Setup certificates
	// Generate node private key

	// connect to master agent
	conn, err := grpc.Dial(*masterAddr, grpc.WithInsecure())
	if err != nil {
		return err
	}

	wa.c = NewNodeServiceClient(wa.conn)
	return nil
}

func (wa *WorkerAgent) Start() error {
	_, _ = wa.c.ReportStatus(context.Background(), &Request{})

	for {
	}
}

func (n *WorkerAgent) GenerateCSR() (err error) {
	if n == nil {
		return fmt.Errorf("Node is not initialized")
	}
	// Generate Node private key
	nodePrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Failed generating RSA Key: %s", err)
	}

	// node certificate template
	nodeTemplate := x509.Certificate{
		SerialNumber:          big.NewInt(2025),
		Subject:               pkix.Name{Organization: []string{""}},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(0, 0, 30), // valid for 30 days
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
	}

	// node certificate signing request template
	nodeCSRTemplate := x509.CertificateRequest{
		Subject: pkix.Name{CommonName: "localhost"},
	}

	// create CSR and get certificate from master
	csrBytes, err := x509.CreateCertificateRequest(rand.Reader, &nodeCSRTemplate, nodePrivateKey)
	if err != nil {
		log.Fatal("Error creating CSR:", err)
	}
	// Encode CSR to PEM format
	csrPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csrBytes})

	// write certificate to file
	certificatePath := []string{"/etc/ssl/certs/", n.Name, ".crt"}
	certificateFile, err := os.Create(strings.Join(certificatePath, ""))
	if err != nil {
		log.Fatalf("Error creating node certificate file: %v", err)
	}
	pem.Encode(certificateFile, &pem.Block{Type: "CERTIFICATE", Bytes: _}) // pick workerCertBytes from return stream
	certificateFile.Close()

	// Load worker node certificate and private key for gRPC server/client configuration
	workerCert, err := tls.LoadX509KeyPair("worker.crt", "worker.key")
	if err != nil {
		log.Fatalf("Error loading worker node certificate and private key: %v", err)
	}

	// Create a gRPC credentials object with the worker node's certificate and private key
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{workerCert},
	})

	return nil

}

// Node join should be a REST API call
// Subsequent communication is via GRPC
