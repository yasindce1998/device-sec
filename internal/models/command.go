package models

import (
    "time"
)

type CommandStatus string

const (
    StatusPending CommandStatus = "pending"
    StatusSent    CommandStatus = "sent"
    StatusDone    CommandStatus = "done"
)

type Command struct {
    ID        string        `json:"id"`
    CreatedAt time.Time     `json:"created_at"`
    Type      string        `json:"type"`
    Payload   string        `json:"payload"` // plist payload
    Status    CommandStatus `json:"status"`
}