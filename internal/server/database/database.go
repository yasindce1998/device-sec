package database

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/device-sec/internal/models"
)

type Database struct {
    db *gorm.DB
}

func NewDatabase(config DatabaseConfig) (*Database, error) {
    dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
        config.Host, config.Port, config.User, config.DBName, config.Password)
    
    db, err := gorm.Open("postgres", dsn)
    if err != nil {
        return nil, err
    }

    // Auto migrate the schema
    db.AutoMigrate(&models.Command{})

    return &Database{db: db}, nil
}

func (d *Database) SaveCommand(cmd *models.Command) error {
    return d.db.Create(cmd).Error
}

func (d *Database) UpdateCommandStatus(id string, status models.CommandStatus) error {
    return d.db.Model(&models.Command{}).Where("id = ?", id).Update("status", status).Error
}

func (d *Database) GetPendingCommands() ([]*models.Command, error) {
    var commands []*models.Command
    err := d.db.Where("status = ?", models.StatusPending).Find(&commands).Error
    return commands, err
}