package main

import (
"encoding/json"
"fmt"
"log"
"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
fmt.Fprintln(w, "Hello, 世界")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")

response := map[string]string{
    "status":      "healthy",
    "application": "Go CI/CD Project",
}

if err := json.NewEncoder(w).Encode(response); err != nil {
    log.Println("Failed to encode health response:", err)
}
}

func main() {
http.HandleFunc("/", handler)
http.HandleFunc("/health", healthHandler)

fmt.Println("Running demo app. Press Ctrl+C to exit...")
log.Fatal(http.ListenAndServe(":8888", nil))
}
