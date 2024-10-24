package middleware

import (
    "github.com/gin-gonic/gin"
    "github.com/device-sec/internal/pkg/logging"
    "net/http"
)

func RecoveryMiddleware(logger *logging.Logger) gin.HandlerFunc {
    return func(c *gin.Context) {
        defer func() {
            if err := recover(); err != nil {
                logger.WithField("error", err).Error("Server panic recovered")
                c.JSON(http.StatusInternalServerError, gin.H{
                    "error": "Internal server error",
                })
            }
        }()
        c.Next()
    }
}