package models

import (
    "time"
    "github.com/google/uuid"
)

type CommandStatus string

const (
    StatusPending CommandStatus = "pending"
    StatusSent    CommandStatus = "sent"
    StatusDone    CommandStatus = "done"
)

type Command struct {
    ID        string        `json:"id" gorm:"primaryKey"`
    CreatedAt time.Time     `json:"created_at"`
    Type      string        `json:"type"`
    Payload   string        `json:"payload"` // plist payload
    Status    CommandStatus `json:"status"`
    AgentID   string       `json:"agent_id"`
}

func NewCommand(agentID, commandType, payload string) *Command {
    return &Command{
        ID:        uuid.New().String(),
        CreatedAt: time.Now(),
        Type:      commandType,
        Payload:   payload,
        Status:    StatusPending,
        AgentID:   agentID,
    }
}