package main

import (
    "github.com/device-sec/config"
    "github.com/device-sec/internal/agent/handler"
    "github.com/device-sec/internal/agent/websocket"
    "log"
    "os"
    "os/signal"
    "syscall"
	"time"
)

func main() {
    // Load configuration
    cfg, err := config.LoadConfig(".")
    if err != nil {
        log.Fatal("Cannot load config:", err)
    }

    logger := log.New(os.Stdout, "AGENT: ", log.LstdFlags)

    // Initialize WebSocket client
    client, err := websocket.NewClient(cfg.Server.BaseURL, os.Getenv("AGENT_ID"))
    if err != nil {
        logger.Fatal("Cannot create WebSocket client:", err)
    }

    // Initialize command handler
    cmdHandler := handler.NewCommandHandler(logger)

    // Handle graceful shutdown
    signalChan := make(chan os.Signal, 1)
    signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

    go func() {
        for {
            if err := client.Connect(); err != nil {
                logger.Println("Connection error:", err)
                time.Sleep(5 * time.Second)
                continue
            }

            for {
                _, message, err := client.conn.ReadMessage()
                if err != nil {
                    logger.Println("Read error:", err)
                    break
                }

                if err := cmdHandler.HandleCommand(message); err != nil {
                    logger.Println("Handle error:", err)
                }
            }
        }
    }()

    <-signalChan
    logger.Println("Shutting down agent...")
}