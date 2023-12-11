package api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeeo/pack-management/internal/config"
	"github.com/rs/cors"
)

type Server struct {
	httpServer *http.Server
	config     config.ServerConfig
}

type ServerHandler interface {
	Configure(*mux.Router)
}

func NewRestServer(config config.ServerConfig, handlers ...ServerHandler) *Server {
	server := &Server{
		config: config,
	}

	address := fmt.Sprintf("0.0.0.0:%d", config.Port)
	mux := mux.NewRouter()

	for _, handler := range handlers {
		handler.Configure(mux)
	}

	corsHandler := cors.AllowAll().Handler(mux)

	server.httpServer = &http.Server{
		Addr:         address,
		Handler:      corsHandler,
		WriteTimeout: server.config.WriteTimeout,
		ReadTimeout:  server.config.ReadTimeout,
	}

	return server
}

func (s Server) Run() error {
	log.Printf("listening on %d \n", s.config.Port)
	return s.httpServer.ListenAndServe()
}

func (s Server) Shutdown() error {
	return s.httpServer.Shutdown(context.TODO())
}
