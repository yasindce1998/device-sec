package api

import (
	"bytes"
	"encoding/json"
	"github.com/device-sec/internal/models"
	"github.com/gin-gonic/gin"
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

func setupTestServer() *Server {
	// Create a mock database (replace with your actual database setup)
	//mockDB := &MockDatabase{}

	// Create a Server instance with the mock database
	server := &Server{
		router: gin.Default(),
		//db:     mockDB, // Inject the mock database
	}

	// Define API routes
	server.setupRoutes()

	return server
}

type MockDatabase struct {
	// Store commands for testing purposes
	commands map[string]models.Command
}

func NewMockDatabase() *MockDatabase {
	return &MockDatabase{
		commands: make(map[string]models.Command),
	}
}

func (m *MockDatabase) SaveCommand(cmd *models.Command) error {
	m.commands[cmd.ID] = *cmd
	return nil
}

func (m *MockDatabase) GetCommand(id string) (*models.Command, error) {
	cmd, ok := m.commands[id]
	if !ok {
		return nil, models.ErrCommandNotFound
	}
	return &cmd, nil
}
