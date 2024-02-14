package server

import (
	"crypto/tls"
	"log"
	"net/http"
)

// Config
type ServerConfig struct {
	address            string
	port               int32
	tlsConfig          *tls.Config
	httpConfig         *http.Client
	WithPortAndAddress bool
}

// Setup Port and Address
func WithPortAndAddress(addr string, port int32) Option {
	return func(cfg *ServerConfig) {
		cfg.address = addr
		cfg.port = port
	}
}

type Option func(*ServerConfig)

// TLS configuration
func WithTLSConfig(tlsConfig *tls.Config) Option {
	return func(cfg *ServerConfig) {
		cfg.tlsConfig = tlsConfig
	}
}

// HTTP configuration
func WithHTTPConfig(httpConfig *http.Client) Option {
	return func(cfg *ServerConfig) {
		cfg.httpConfig = httpConfig
	}
}

// Create new Config
func NewConfig(opts ...Option) *ServerConfig {
	cfg := &ServerConfig{}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// Get Address
func (c *ServerConfig) Addr() string {
	c.ValidatePortAddress()
	return c.address
}

// Get Port
func (c *ServerConfig) Port() int32 {
	c.ValidatePortAddress()
	return c.port
}

func (c *ServerConfig) ValidatePortAddress() {
	if len(c.address) == 0 && c.port <= 0 {
		log.Fatal("must set address and port")
	}
}
