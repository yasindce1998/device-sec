package database

// import (
//     "github.com/device-sec/internal/models"
//     "github.com/jinzhu/gorm"
//     "testing"
//     "time"
// )

// func TestSaveCommand(t *testing.T) {
//     db := setupTestDB()
    
//     cmd := &models.Command{
//         ID:        "test-id",
//         CreatedAt: time.Now(),
//         Type:      "install-app",
//         Payload:   "<plist>test</plist>",
//         Status:    models.StatusPending,
//         AgentID:   "test-agent",
//     }
    
//     err := db.SaveCommand(cmd)
//     if err != nil {
//         t.Errorf("Failed to save command: %v", err)
//     }
    
//     // Verify command was saved
//     saved, err := db.GetCommand(cmd.ID)
//     if err != nil {
//         t.Errorf("Failed to get command: %v", err)
//     }
    
//     if saved.ID != cmd.ID {
//         t.Errorf("Expected ID %s, got %s", cmd.ID, saved.ID)
//     }
// }

// func setupTestDB() *gorm.DB {
//     db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
//     if err != nil {
//       panic("failed to connect to test database")
//     }
//     return db
//   }