package main

import (
    "fmt"
    "net/http"
    
    "go.uber.org/zap"
)

func main() {
    logger, _ := zap.NewProduction()
    defer logger.Sync()
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Snowflake Gateway Service is running")
    })
    
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Snowflake Gateway Service is healthy")
    })
    
    logger.Info("Snowflake Gateway Service starting on port 8080")
    http.ListenAndServe(":8080", nil)
}