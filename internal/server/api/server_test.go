package api

import (
    "bytes"
    "encoding/json"
    "github.com/device-sec/internal/models"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestCreateCommand(t *testing.T) {
    // Setup test server
    server := setupTestServer()
    
    command := models.Command{
        Type:    "install-app",
        Payload: "<plist>test</plist>",
        AgentID: "test-agent",
    }
    
    jsonData, _ := json.Marshal(command)
    
    req := httptest.NewRequest("POST", "/api/v1/commands", bytes.NewBuffer(jsonData))
    req.Header.Set("Content-Type", "application/json")
    
    w := httptest.NewRecorder()
    server.router.ServeHTTP(w, req)
    
    if w.Code != http.StatusCreated {
        t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
    }
}