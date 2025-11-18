package handlers

import "github.com/gorilla/mux"

// RegisterAPIRoutes registers all API routes on the given subrouter.
// Currently only product code APIs are exposed.
func RegisterAPIRoutes(api *mux.Router) {
	// Product code endpoints
	api.HandleFunc("/product-codes", GetProductCodes).Methods("GET")
	api.HandleFunc("/product-codes", CreateProductCode).Methods("POST")
	api.HandleFunc("/product-codes/{id}", DeleteProductCode).Methods("DELETE")
}

