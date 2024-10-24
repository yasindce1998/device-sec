package logging

import (
    "github.com/sirupsen/logrus"
    "os"
)
type LogConfig struct{
    Format string
    Level string
}
type Logger struct {
    *logrus.Logger
}

func NewLogger(config LogConfig) *Logger {
    logger := logrus.New()
    
    // Set output format
    if config.Format == "json" {
        logger.SetFormatter(&logrus.JSONFormatter{})
    } else {
        logger.SetFormatter(&logrus.TextFormatter{
            FullTimestamp: true,
        })
    }

    // Set output
    logger.SetOutput(os.Stdout)

    // Set log level
    level, err := logrus.ParseLevel(config.Level)
    if err != nil {
        level = logrus.InfoLevel
    }
    logger.SetLevel(level)

    return &Logger{logger}
}