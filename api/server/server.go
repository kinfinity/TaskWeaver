package server

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"taskweaver/api/middleware"
	"taskweaver/api/node"
	"taskweaver/api/rest"
	"time"

	"github.com/gorilla/mux"
)

type server struct {
	baseRouter *mux.Router // base Router
	subRouters []rest.Router
	config     *ServerConfig
	httpServer *http.Server
	logger     *log.Logger
	port       int32
}

const (
	api_version = "/api/v1"
)

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
		"[ "+AppName+" ]: ",
		0,
	)
	// Create Router Logger Middleware
	routerLogger := middleware.NewRequestLogger(logger)
	router.Use(routerLogger.HandleMiddleware())
	// Create Auth  Middleware
	authM := middleware.AuthMiddleware{}

	// Initialize all Routers and Routes w Middleware
	router.NewRoute().Methods("GET").PathPrefix(api_version + "/").HandlerFunc(HealthCheck)
	nodeRouter := node.NewNodeRouter(router, api_version+"/nodes", authM)

	// append nodeRouter to rest.Router list
	routerList := make([]rest.Router, 10)
	routerList = append(routerList, nodeRouter)

	logger.Printf("API Server Starting... %v ", time.Now().Format("2006-01-02 15:04:05"))
	rest.PrintMuxRoutes(logger, router)

	// close log file?
	// defer logFile.Close()

	return &server{
		baseRouter: router,
		subRouters: routerList,
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
		Addr:           s.config.address + ":" + strconv.Itoa(int(s.config.port)),
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
