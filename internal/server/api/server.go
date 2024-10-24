package api

import (
    "github.com/gin-gonic/gin"
    "github.com/device-sec/internal/server/database"
    "github.com/device-sec/internal/server/queue"
    "sync"
)

type Server struct {
    router      *gin.Engine
    db          *database.Database
    broker      *queue.RabbitMQ
    agents      map[string]*websocket.Conn
    agentsMutex sync.RWMutex
}

func NewServer(db *database.Database, broker *queue.RabbitMQ) *Server {
    server := &Server{
        router: gin.Default(),
        db:     db,
        broker: broker,
        agents: make(map[string]*websocket.Conn),
    }
    
    server.setupRoutes()
    return server
}

func (s *Server) Run(addr string) error {
    return s.router.Run(addr)
}
