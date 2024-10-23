package api

import (
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
)

type Server struct {
    router *gin.Engine
    db     Database
    broker MessageBroker
}

func NewServer(db Database, broker MessageBroker) *Server {
    router := gin.Default()
    server := &Server{
        router: router,
        db:     db,
        broker: broker,
    }
    
    server.setupRoutes()
    return server
}

func (s *Server) setupRoutes() {
    s.router.POST("/commands", s.createCommand)
    s.router.GET("/ws", s.handleWebSocket)
}

func (s *Server) createCommand(c *gin.Context) {
    // Handle command creation and queueing
}

func (s *Server) handleWebSocket(c *gin.Context) {
    // Handle WebSocket connections from agents
}

type Database interface {
    SaveCommand(cmd *Command) error
    UpdateCommandStatus(id string, status CommandStatus) error
    GetPendingCommands() ([]*Command, error)
}

type MessageBroker interface {
    PublishCommand(cmd *Command) error
    SubscribeToCommands() (<-chan *Command, error)
}