package main

import (
	"github.com/device-sec/config"
	"github.com/device-sec/internal/server/api"
	"github.com/device-sec/internal/server/database"
	"github.com/device-sec/internal/server/queue"
	"log"
)

func main() {

	var cfg config.Config

	// Create a DatabaseConfig object from the nested fields
	dbConfig := database.DatabaseConfig{
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		DBName:   cfg.Database.DBName,
	}

	// Load configuration
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	// Initialize database
	db, err := database.NewDatabase(dbConfig)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	// Initialize message broker
	broker, err := queue.NewRabbitMQ(cfg.RabbitMQ.URL)
	if err != nil {
		log.Fatal("Cannot connect to RabbitMQ:", err)
	}
	defer broker.Close()

	// Initialize and start server
	server := api.NewServer(db, broker)
	if err := server.Run(cfg.Server.Port); err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
