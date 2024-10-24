package api

import (
    "github.com/gin-gonic/gin"
    "github.com/device-sec/internal/models"
    "net/http"
    "github.com/gorilla/websocket"
)
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
      return true // Allow all origins for now (consider security implications)
    },
  }
func (s *Server) createCommand(c *gin.Context) {
    var cmd models.Command
    if err := c.ShouldBindJSON(&cmd); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := s.db.SaveCommand(&cmd); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if err := s.broker.PublishCommand(&cmd); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, cmd)
}
func (s *Server) getCommand(c *gin.Context) {
    id := c.Param("id")
  
    // Retrieve command by ID from database
    cmd, err := s.db.GetCommand(id)
    if err != nil {
      // Handle potential errors (e.g., not found, database error)
      switch err {
      case models.ErrCommandNotFound:
        c.JSON(http.StatusNotFound, gin.H{"error": "Command not found"})
      default:
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      }
      return
    }
  
    // Return the retrieved command in the response
    c.JSON(http.StatusOK, cmd)
  }
func (s *Server) handleWebSocket(c *gin.Context) {
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    agentID := c.Query("agent_id")
    s.agentsMutex.Lock()
    s.agents[agentID] = conn
    s.agentsMutex.Unlock()

    // Handle WebSocket connection
    go s.handleWebSocketConnection(agentID, conn)
}