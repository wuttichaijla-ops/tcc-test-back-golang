package main

import (
	"log"
	"net/http"
	"test-back-golang/datasource"
	"test-back-golang/handlers"
	"test-back-golang/middleware"

	"github.com/gorilla/mux"
)

func main() {
	// Init database connection (GORM)
	datasource.InitDatabase()

	// Create router
	r := mux.NewRouter()

	// API routes
	api := r.PathPrefix("/api").Subrouter()
	handlers.RegisterAPIRoutes(api)

	// Apply CORS middleware
	handler := middleware.CORS(r)

	// Start server
	port := ":8080"
	log.Printf("Server starting on port %s", port)
	log.Println("API endpoints:")
	log.Println("  GET    /api/product-codes")
	log.Println("  POST   /api/product-codes")
	log.Println("  DELETE /api/product-codes/{id}")

	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
