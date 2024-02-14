package server

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"taskweaver/api/middleware"
	"taskweaver/api/rest"
	"time"

	"github.com/gorilla/mux"
)

type server struct {
	baseRouter *mux.Router // base Router
	config     *ServerConfig
	httpServer *http.Server
	logger     *log.Logger
	port       int32
}

// Add Router w Routes
func (s *server) AddRoutersWithRoutes(routers ...rest.Router) {

}

// New Server with Config
func NewServer(config *ServerConfig, router *mux.Router) *server {
	// Open or create a log file
	logFile, _ := CreateAndOpen("api.log")

	// Create a logger
	logger := log.New(
		io.MultiWriter(
			os.Stdout,
			NewFileWriter(logFile),
		),
		"[ TaskWeaver ]: ",
		0,
	)
	// Create Router Logger Middleware
	routerLogger := middleware.NewRequestLogger(logger)
	// Initialize all Routers and Routes w Middleware
	router.Use(routerLogger.HandleMiddleware())
	logger.Printf("API Server Starting... %v ", time.Now().Format("2006-01-02 15:04:05"))

	// close logfile?
	// defer logFile.Close()

	return &server{
		baseRouter: router,
		config:     config,
		logger:     logger,
		port:       config.Port(),
	}
}

// Start the HTTP Server
func (s *server) ListenAndServe() error {
	if s.baseRouter == nil {
		panic("No Base Handler provided")
	}

	s.httpServer = &http.Server{
		Addr:           s.config.address,
		Handler:        s.baseRouter,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    60 * time.Second,
		MaxHeaderBytes: 1024 * 8, // 8 KB
		TLSConfig:      s.config.tlsConfig,
	}

	errChan := make(chan error)
	go func() {
		s.logger.Printf("%s listening on :%d", AppName, s.config.port)
		errChan <- s.httpServer.ListenAndServe()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = s.httpServer.Shutdown(ctx)
	close(errChan)
	return <-errChan
}
