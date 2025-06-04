package logging

import (
    "io"
    "log"
    "os"
    "path/filepath"
)

// Setup initializes logging to both file and console
func Setup(logDir string) (*os.File, error) {
    logFilePath := filepath.Join(logDir, "app.log")
    logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        return nil, err
    }

    multiWriter := io.MultiWriter(os.Stdout, logFile)
    log.SetOutput(multiWriter)
    log.SetFlags(log.LstdFlags | log.Lshortfile)
    return logFile, nil
}