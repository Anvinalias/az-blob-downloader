package logging

import (
    "io"
    "log"
    "os"
    "path/filepath"
    "time"
    "fmt"
)

// Setup initializes logging to both file and console
func Setup(logDir string) (*os.File, error) {
    timestamp := time.Now().Format("150405") // HHMMSS (24-hour format)
    logFileName := fmt.Sprintf("app_%s.log", timestamp)
    logFilePath := filepath.Join(logDir, logFileName)
    logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        return nil, err
    }

    multiWriter := io.MultiWriter(os.Stdout, logFile)
    log.SetOutput(multiWriter)
    log.SetFlags(log.LstdFlags)
    return logFile, nil
}