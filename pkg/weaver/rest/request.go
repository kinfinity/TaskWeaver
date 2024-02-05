package rest_client

import (
	"encoding/json"
	"time"

)

type Header struct {
	Content   ContentType
	UserAgent string
}

func NewHeader() *Header {
	return &Header{
		Content:   "application/json",
		UserAgent: "weaver/1.0",
	}
}

// An interface for request objects.
type IRequest interface {
	Normalize()
	Build()
}

// Base Request across rest calls
type clientRequest struct {
	Data            any
	JSONData        string
	Head            Header
	Timeout         time.Duration
	MaxRetries      uint16
	RetryFunc       func(error) bool
	RetryAfterSecs  uint16
	RequestBody     []byte
	QueryString     string
	QueryParameters map[string]string
	Path            RequestPath
	Verb            HTTPVerb
}

// Normalize Request data
func (cr *clientRequest) Normalize() {
	// Unmarshal the response JSON directly into the existing instance
	_data, err := json.Marshal(cr.Data)
	if err != nil {
		// Handle unmarshal error
		return
	}
	cr.JSONData = string(_data)
}

func (cr *clientRequest) Header() *Header {
	return &cr.Head
}

// Build Request
func (cr *clientRequest) Build() {}

func (cr *clientRequest) setContentType() *clientRequest {
	return cr
}
func (cr *clientRequest) AddData(Body any) *clientRequest {
	cr.verifyData()
	return cr
}
func (cr *clientRequest) verifyData() {

}
func (cr *clientRequest) SetQueryString() *clientRequest            { return cr }
func (cr *clientRequest) SetQueryParams() *clientRequest            { return cr }
func (cr *clientRequest) SetPath(requestPath string) *clientRequest { return cr }

// Request Method
func (cr *clientRequest) SetVerb(verb HTTPVerb) {
	cr.Verb = verb
}

// GetJson returns the Json representation of this request's Data field.
func (cr *clientRequest) GetJson() string { return cr.JSONData }

// Build from query params ...
func (cr *clientRequest) BuildUrl() string { return cr.JSONData }

// Initialize ClientRequest
func NewClientRequest(data interface{}, timeout time.Duration, maxRetries uint16, retryFunc func(error) bool, retryAfterSecs uint16) *clientRequest {
	return &clientRequest{
		Data:           data,
		Timeout:        timeout,
		MaxRetries:     maxRetries,
		RetryFunc:      retryFunc,
		RetryAfterSecs: retryAfterSecs,
		Head: Header{
			Content:   UseContentType(ContentTypeJSON),
			UserAgent: "",
		},
	}
}

// set content length
// request size calculation
