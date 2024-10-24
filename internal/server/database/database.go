package database

import (
	"fmt"
	"github.com/device-sec/internal/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Database struct {
	db *gorm.DB
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	DBName   string `json:"dbname"`
	Password string `json:"password"`
}

func NewDatabase(cfg DatabaseConfig) (*Database, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.DBName, cfg.Password)

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

func (d *Database) GetCommand(id string) (*models.Command, error) {
	var cmd models.Command

	err := d.db.First(&cmd, id).Error
	if err != nil {
		return nil, err // Handle database query errors
	}

	return &cmd, nil
}
