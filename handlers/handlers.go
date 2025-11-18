package handlers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"test-back-golang/datasource"
	"test-back-golang/models"

	"github.com/gorilla/mux"
)

// ===== Product Code CRUD (16 chars, format XXXX-XXXX-XXXX-XXXX) =====

var productCodePattern = regexp.MustCompile(`^[A-Z0-9]{4}(-[A-Z0-9]{4}){3}$`)

type ProductCodeRequest struct {
	ProductName string `json:"product_name"`
	Code        string `json:"code"`
}

// GetProductCodes returns all product codes
func GetProductCodes(w http.ResponseWriter, r *http.Request) {
	if datasource.DB == nil {
		http.Error(w, "database not initialized", http.StatusInternalServerError)
		return
	}

	var codes []models.ProductCode
	if err := datasource.DB.Order("id asc").Find(&codes).Error; err != nil {
		http.Error(w, "failed to fetch product codes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(codes)
}

// CreateProductCode creates a new product code
func CreateProductCode(w http.ResponseWriter, r *http.Request) {
	if datasource.DB == nil {
		http.Error(w, "database not initialized", http.StatusInternalServerError)
		return
	}

	var req ProductCodeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(req.ProductName) == "" {
		http.Error(w, "product_name is required", http.StatusBadRequest)
		return
	}

	code := req.Code
	code = strings.ToUpper(code)

	if !productCodePattern.MatchString(code) {
		http.Error(w, "code must be 16 characters in format XXXX-XXXX-XXXX-XXXX (A-Z, 0-9)", http.StatusBadRequest)
		return
	}

	pc := models.ProductCode{
		ProductName: req.ProductName,
		Code:        code,
	}
	if err := datasource.DB.Create(&pc).Error; err != nil {
		http.Error(w, "failed to create product code", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pc)
}

// DeleteProductCode deletes a product code by ID
func DeleteProductCode(w http.ResponseWriter, r *http.Request) {
	if datasource.DB == nil {
		http.Error(w, "database not initialized", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := datasource.DB.Delete(&models.ProductCode{}, id).Error; err != nil {
		http.Error(w, "failed to delete product code", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
