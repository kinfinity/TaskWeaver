/*
 */
package rest_client

import (
	"crypto/tls"
	"net/http"
	"time"
)

type RequestOptions struct {
	TimeOut time.Duration
}

// Config
type ClientConfig struct {
	TLSConfig       *tls.Config
	HTTPConfig      *http.Client
	BaseURL         string
	DRequestOptions *RequestOptions
}

type Option func(*ClientConfig)

// TLS configuration
func WithTLSConfig(tlsConfig *tls.Config) Option {
	return func(cfg *ClientConfig) {
		cfg.TLSConfig = tlsConfig
	}
}

// HTTP configuration
func WithHTTPConfig(httpConfig *http.Client) Option {
	return func(cfg *ClientConfig) {
		cfg.HTTPConfig = httpConfig
	}
}

// base URL
func WithBaseURL(baseURL string) Option {
	return func(cfg *ClientConfig) {
		cfg.BaseURL = baseURL
	}
}

// Create new Config
func newConfig(opts ...Option) *ClientConfig {
	cfg := &ClientConfig{}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}
