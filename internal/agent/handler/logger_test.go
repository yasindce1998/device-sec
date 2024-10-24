// handler/logger_test.go
package handler

import (
    "io"
    "log"
    "os"
)

type TestLogger struct {
    *log.Logger
    logFile *os.File
}

// setupTestLogger creates a logger suitable for testing.
func setupTestLogger() *TestLogger {
    // Create a multi-writer that writes to both a file and discard
    logFile, err := os.CreateTemp("", "test-agent-*.log")
    if err != nil {
        // If we can't create a file, just use discard
        return &TestLogger{
            Logger: log.New(io.Discard, "TEST: ", log.LstdFlags),
        }
    }

    // Create multi-writer for both file and discard
    writer := io.MultiWriter(logFile, io.Discard)

    // Return new logger with test prefix
    return &TestLogger{
        Logger:  log.New(writer, "TEST: ", log.LstdFlags|log.Lshortfile),
        logFile: logFile,
    }
}

// Cleanup should be called after the test is done
func (tl *TestLogger) Cleanup() {
    if tl.logFile != nil {
        tl.logFile.Close()
        os.Remove(tl.logFile.Name())
    }
}