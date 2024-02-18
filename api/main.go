package main

import (
	"crypto/tls"
	"net/http"
	"taskweaver/api/server"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	server := server.NewServer(
		server.NewConfig(
			server.WithPortAndAddress("localhost:", 3000),
			server.WithHTTPConfig(&http.Client{}),
			server.WithTLSConfig(&tls.Config{}),
		),
		router,
	)

	server.ListenAndServe()
}
