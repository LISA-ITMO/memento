package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"memento/internal/services/storage"
	"net/http"
)

type Server struct {
	storage storage.Storage
	router  *gin.Engine
}

func New(storage storage.Storage) *Server {
	server := Server{
		storage,
		gin.Default(),
	}
	server.setupRoutes()
	return &server
}

func (s *Server) MustRun() {
	err := http.ListenAndServe(":8080", s.router)
	if err != nil {
		log.Fatalf("Невозможно запустить сервер: %v", err.Error())
	}
}

