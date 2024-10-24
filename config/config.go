package config

import (
    "github.com/spf13/viper"
    "github.com/device-sec/internal/server/database"
)

type Config struct {
    Server struct {
        Port    string `mapstructure:"port"`
        BaseURL string `mapstructure:"base_url"`
    }
    Database database.DatabaseConfig
    RabbitMQ struct {
        URL string `mapstructure:"url"`
    }
}

func LoadConfig(path string) (config Config, err error) {
    viper.AddConfigPath(path)
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")

    viper.AutomaticEnv()

    err = viper.ReadInConfig()
    if err != nil {
        return
    }

    err = viper.Unmarshal(&config)
    return
}