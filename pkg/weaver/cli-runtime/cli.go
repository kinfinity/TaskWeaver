package cli_runtime

import (
	"io"
	"log"
	"os"
	"sync"
	weavercmd "taskweaver/pkg/weaver/cmd"

	"github.com/spf13/cobra"
)

type WCli struct {
	name      string
	version   string
	IOStreams struct {
		In  io.Reader
		Out io.Writer
		Err io.Writer
	}
	logger *log.Logger
}

var _Client *WCli
var lock sync.Once

type APIClient interface {
	WCli
	RunCommand(command weavercmd.CliCommand)
}

func (c *WCli) RunCommand(cmd *cobra.Command) error {
	return nil
}

func NewCli(co *ClientOptions) (*WCli, error) {
	co.logger.Println("Start with Commands")
	/*
		setup
		- io
		- logging

	*/
	if _Client == nil {
		lock.Do(
			func() {
				_Client = &WCli{
					name:    co.Name,
					version: co.Version,
					logger:  co.GetLogger(),
				}
				_Client.IOStreams.In = os.Stdin
				_Client.IOStreams.Out = os.Stdout
				_Client.IOStreams.Err = os.Stderr

			})
	}

	return _Client, nil
}

// Get Logger
func (cli *WCli) GetLogger() (logger *log.Logger) {
	return cli.logger
}

// Client handles it's IO
// Each command doesn't know how it is ran or it's IO happens
// The client knows the IO and can handle that for each command

//  --------------------------------------
/*
// Create a config with TLS, HTTP, and Base configurations

	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	httpConfig := &http.Client{Timeout: 10}
	baseURL := "https://api.example.com"

	restClientConfig := newConfig(
		WithTLSConfig(tlsConfig),
		WithHTTPConfig(httpConfig),
		WithBaseURL(baseURL),
	)
*/
