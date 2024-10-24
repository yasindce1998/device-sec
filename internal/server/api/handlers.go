package api

import (
    "github.com/gin-gonic/gin"
    "github.com/device-sec/internal/models"
    "net/http"
)

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