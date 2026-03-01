package main

import (
    "fmt"
    "net/http"
)

func main() {
    // TODO: Initialize database connection (PostgreSQL/MySQL)
    // TODO: Set up environment variables for config

    http.HandleFunc("/api/register", registerHandler)
    http.HandleFunc("/api/login", loginHandler)

    fmt.Println("Server starting on :9090...")
    // TODO: Add graceful shutdown
    if err := http.ListenAndServe(":9090", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    // TODO: Parse JSON body
    // TODO: Validate user input (email format, password strength)
    // TODO: Hash password using bcrypt
    // TODO: Save user to database
    fmt.Fprintf(w, "Register endpoint hit")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    // TODO: Parse JSON body
    // TODO: Find user in database
    // TODO: Compare password hash
    // TODO: Generate JWT token and return it
    fmt.Fprintf(w, "Login endpoint hit")
}
