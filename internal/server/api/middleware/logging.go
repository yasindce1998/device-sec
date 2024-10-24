package middleware

import (
    "github.com/gin-gonic/gin"
    "github.com/device-sec/internal/pkg/logging"
	"github.com/sirupsen/logrus"
    "time"
)

func LoggingMiddleware(logger *logging.Logger) gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        path := c.Request.URL.Path
        
        c.Next()
        
        end := time.Now()
        latency := end.Sub(start)
        
        logger.WithFields(logrus.Fields{
            "status":     c.Writer.Status(),
            "method":     c.Request.Method,
            "path":       path,
            "latency":    latency,
            "client_ip":  c.ClientIP(),
            "user_agent": c.Request.UserAgent(),
        }).Info("Request completed")
    }
}