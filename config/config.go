package config

import (
    "github.com/spf13/viper"
)

type Config struct {
    Server struct {
        Port    string `mapstructure:"port"`
        BaseURL string `mapstructure:"base_url"`
    }
    Database struct {
        Host     string `mapstructure:"host"`
        Port     string `mapstructure:"port"`
        User     string `mapstructure:"user"`
        Password string `mapstructure:"password"`
        DBName   string `mapstructure:"dbname"`
    }
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