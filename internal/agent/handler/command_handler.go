package handler

import (
    "encoding/json"
    "github.com/device-sec/internal/models"
    "log"
)

type CommandHandler struct {
    logger *log.Logger
}

func NewCommandHandler(logger *log.Logger) *CommandHandler {
    return &CommandHandler{
        logger: logger,
    }
}

func (h *CommandHandler) HandleCommand(payload []byte) error {
    var cmd models.Command
    if err := json.Unmarshal(payload, &cmd); err != nil {
        return err
    }

    h.logger.Printf("Received command: %s with payload: %s\n", cmd.ID, cmd.Payload)
    return nil
}