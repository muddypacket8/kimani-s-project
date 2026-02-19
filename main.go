package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// --- STEP 1: Hello World Handler ---
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World! Welcome to the Go Web Toolkit 🚀")
}

// --- STEP 2: JSON API Response struct ---
type ApiResponse struct {
	Message   string `json:"message"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

// --- STEP 3: REST API Handler ---
func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := ApiResponse{
		Message:   "Welcome to the Go Web Toolkit API!",
		Status:    "success",
		Timestamp: time.Now().Format(time.RFC3339),
	}

	json.NewEncoder(w).Encode(response)
}

// --- STEP 4: About page handler ---
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Go Web Toolkit - Moringa AI Capstone Project")
	fmt.Fprintln(w, "Built with Go standard library. No external dependencies needed!")
}

// --- STEP 5: 404 Not Found handler ---
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "404 - Page Not Found")
}

func main() {
	// --- ROUTING ---
	mux := http.NewServeMux()

	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/api", apiHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/404", notFoundHandler)

	// --- START SERVER ---
	port := ":8080"
	fmt.Printf("🚀 Go Web Toolkit server running at http://localhost%s\n", port)
	fmt.Println("Routes available:")
	fmt.Println("  GET /        → Hello World")
	fmt.Println("  GET /api     → JSON API Response")
	fmt.Println("  GET /about   → About page")

	log.Fatal(http.ListenAndServe(port, mux))
}
