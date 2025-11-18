package main

import (
	"log"
	"net/http"
	"test-back-golang/handlers"
	"test-back-golang/middleware"

	"github.com/gorilla/mux"
)

func main() {
	// Create router
	r := mux.NewRouter()

	// API routes
	api := r.PathPrefix("/api").Subrouter()
	
	// Items endpoints
	api.HandleFunc("/items", handlers.GetItems).Methods("GET")
	api.HandleFunc("/items", handlers.AddItem).Methods("POST")
	api.HandleFunc("/items/{id}", handlers.DeleteItem).Methods("DELETE")

	// Comments endpoints
	api.HandleFunc("/addComment", handlers.AddComment).Methods("POST")
	api.HandleFunc("/comments", handlers.GetComments).Methods("GET")

	// Apply CORS middleware
	handler := middleware.CORS(r)

	// Start server
	port := ":8080"
	log.Printf("Server starting on port %s", port)
	log.Println("API endpoints:")
	log.Println("  GET    /api/items")
	log.Println("  POST   /api/items")
	log.Println("  DELETE /api/items/{id}")
	log.Println("  POST   /api/addComment")
	log.Println("  GET    /api/comments")
	
	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

