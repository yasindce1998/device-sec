package handler

import (
    "encoding/json"
    "github.com/device-sec/internal/models"
    "testing"
)

func TestHandleCommand(t *testing.T) {
    logger := setupTestLogger()
    handler := NewCommandHandler(logger.Logger)
    
    cmd := &models.Command{
        ID:      "test-id",
        Type:    "install-app",
        Payload: "<plist>test</plist>",
        Status:  models.StatusPending,
    }
    
    payload, _ := json.Marshal(cmd)
    
    err := handler.HandleCommand(payload)
    if err != nil {
        t.Errorf("Failed to handle command: %v", err)
    }
}