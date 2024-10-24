package errors

import "fmt"

type ErrorType string

const (
    ErrorTypeDatabase     ErrorType = "DATABASE_ERROR"
    ErrorTypeRabbitMQ    ErrorType = "RABBITMQ_ERROR"
    ErrorTypeWebSocket   ErrorType = "WEBSOCKET_ERROR"
    ErrorTypeValidation  ErrorType = "VALIDATION_ERROR"
)

type AppError struct {
    Type    ErrorType
    Message string
    Err     error
}

func (e *AppError) Error() string {
    if e.Err != nil {
        return fmt.Sprintf("%s: %s: %v", e.Type, e.Message, e.Err)
    }
    return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

func NewError(errorType ErrorType, message string, err error) *AppError {
    return &AppError{
        Type:    errorType,
        Message: message,
        Err:     err,
    }
}