package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/gorilla/mux"
)

type RestError struct {
	StatusCode uint32
	Message    string `json:"message,omitempty"`
}

type HTTPVerb string

const (
	GET    HTTPVerb = "GET"
	POST   HTTPVerb = "POST"
	DELETE HTTPVerb = "DELETE"
	PUT    HTTPVerb = "PUT"
)

// Check Verb Validity
func isValidHTTPVerb(verb HTTPVerb) bool {
	switch verb {
	case GET, POST, DELETE, PUT:
		return true
	default:
		return false
	}
}

func NewHTTPVerb(value string) (HTTPVerb, error) {
	newVerb := HTTPVerb(strings.ToUpper(value))
	if !isValidHTTPVerb(newVerb) {
		return "", errors.New("invalid HTTP verb")
	}
	return newVerb, nil
}

type ContentType string

const (
	ContentTypeJSON           ContentType = "application/json"
	ContentTypeXML            ContentType = "application/xml"
	ContentTypeFormURLEncoded ContentType = "application/x-www-form-urlencoded"
)

func isValidContentType(contentType ContentType) bool {
	switch contentType {
	case ContentTypeJSON, ContentTypeXML, ContentTypeFormURLEncoded:
		return true
	default:
		return false
	}
}

func UseContentType(contentType ContentType) ContentType {
	if !isValidContentType(contentType) {
		err := errors.New("invalid content type")
		fmt.Println(err)
		return ""
	}
	return contentType
}

type RequestPath string

func isValidRequestPath(path RequestPath) bool {
	pathRegex := regexp.MustCompile(`^(/[\w-]+(/[a-zA-Z0-9-_]+)*)?$`)
	return pathRegex.MatchString(string(path))
}

func PrintMuxRoutes(logger *log.Logger, router *mux.Router) {
	logger.Println("Registered Routes:")
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err == nil {
			methods, _ := route.GetMethods()
			logger.Printf("Path: %s, Methods: %v\n", path, methods)
		}
		return err
	})
}

func encode[T any](w http.ResponseWriter, r *http.Request, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

func decode[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil
}
