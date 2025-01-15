package main

import (
	"forum/internal/db"
	"forum/internal/handlers"
	"log"
	"net/http"
)

func main() {
	// Initialize the database
	db.Init()

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Set up routes
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.HandleFunc("/register", handlers.RegisterHandler)
	// mux.HandleFunc("/post", handlers.PostHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Server started at http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
