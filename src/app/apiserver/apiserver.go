package apiserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nurtaims/golang-mongodb/src/app/auth"
	"github.com/sirupsen/logrus"
)

// APIServer type
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

func (server *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(server.config.LogLevel)
	if err != nil {
		return err
	}
	server.configureRouter()
	server.loger.SetLevel(level)
	return http.ListenAndServe(server.config.BindAddr, server.router)
}

// API METHODS
func (server *APIServer) configureRouter() {
	server.router.HandleFunc("/api/getall", auth.GetAllFunc()).Methods("GET")
	server.router.HandleFunc("/api/create", auth.CreateFunc()).Methods("POST")
	server.router.HandleFunc("/api/getone/{id}", auth.GetOneFunc()).Methods("GET")
}
