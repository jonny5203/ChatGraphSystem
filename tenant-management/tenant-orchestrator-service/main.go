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
        fmt.Fprintf(w, "Tenant Orchestrator Service is running")
    })
    
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Tenant Orchestrator Service is healthy")
    })
    
    logger.Info("Tenant Orchestrator Service starting on port 8080")
    http.ListenAndServe(":8080", nil)
}