package apiserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nurtaims/golang-mongodb/app/auth"
	"github.com/sirupsen/logrus"
)

// APIServer ...
type APIServer struct {
	config *Config
	loger  *logrus.Logger
	router *mux.Router
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		loger:  logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start ...
func (server *APIServer) Start() error {
	if err := server.configureLogger(); err != nil {
		return err
	}
	server.loger.Info("Starting API server")
	return nil
}

// configureLogger ...
func (server *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(server.config.LogLevel)
	if err != nil {
		return err
	}
	server.configureRouter()
	server.loger.SetLevel(level)
	return http.ListenAndServe(server.config.BindAddr, server.router)
}

func (server *APIServer) configureRouter() {
	server.router.HandleFunc("/api/login", auth.LoginFunc())
}
