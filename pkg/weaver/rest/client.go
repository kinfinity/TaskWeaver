/*
 */
package rest_client

import (
	"net/http"
	"net/url"
	"strings"
)

type Client interface {
	Post() *IResponse
	Get() *IResponse
	Execute(Req IRequest) IResponse
	ExecuteMultiple(ReqList []IRequest) <-chan IResponse
}

type RESTClient struct {
	base         *url.URL
	Client       *http.Client
	ResponseList []IResponse
	// ? how to handle each response
}

func NewClient(config ClientConfig, client *http.Client) (*RESTClient, error) {
	base, _ := url.Parse(config.BaseURL)
	if !strings.HasSuffix(base.Path, "/") {
		base.Path += "/"
	}
	base.RawQuery = ""
	base.Fragment = ""

	return &RESTClient{
		base:   base,
		Client: client,
	}, nil
}

func (rc *RESTClient) withMiddleware() *RESTClient

func (rc *RESTClient) Execute(Req IRequest) IResponse {
	// Build Request
	Req.Build()
	Req.Normalize()

	// Execute Request
	return nil
}

func (rc *RESTClient) ExecuteMultiple(ReqList []IRequest) <-chan IResponse {
	// Create a channel to receive responses
	responseChan := make(chan IResponse, len(ReqList))

	// Start goroutines for all requests
	for _, req := range ReqList {
		// spawn multiple processes to handle each request
		go func(r IRequest) {
			resp := rc.Execute(r)
			// responses should be synced back into a channel
			responseChan <- resp
		}(req)
	}

	// Spawn a goroutine to close the channel once all responses are received
	go func() {
		// Wait for all responses to be received
		for i := 0; i < len(ReqList); i++ {
			<-responseChan
		}
		close(responseChan)
	}()

	return responseChan
}

/*
	 USAGE:
	 // Sleep for a short duration before starting to receive responses
	time.Sleep(500 * time.Millisecond)

	// Now, start receiving responses from the channel
	for response := range responseChan {
		// Process each response as it arrives
		fmt.Println("Received response:", response)
	}
*/

// since Request holds just the path of resource we
// need to build URL with baseURL if not provided in the request itself
func NormalizeRequestURL(Req IRequest) {}

func SetToken(token string) {}
