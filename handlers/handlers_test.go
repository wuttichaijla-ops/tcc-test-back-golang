package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"test-back-golang/datasource"
	"test-back-golang/models"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupTestDB initializes an in-memory sqlite database for tests.
func setupTestDB(t *testing.T) {
	t.Helper()

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open test database: %v", err)
	}

	if err := db.AutoMigrate(&models.ProductCode{}); err != nil {
		t.Fatalf("failed to migrate test database: %v", err)
	}

	datasource.DB = db
}

// setupRouter creates a mux router with API routes registered, similar to main.go.
func setupRouter() *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	RegisterAPIRoutes(api)
	return r
}

func TestCreateProductCode_Success(t *testing.T) {
	setupTestDB(t)
	router := setupRouter()

	body := ProductCodeRequest{
		ProductName: "Test Product",
		Code:        "ABCD-1234-EFGH-5678",
	}
	payload, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/product-codes", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d, body=%s", http.StatusCreated, rr.Code, rr.Body.String())
	}

	var resp models.ProductCode
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if resp.ID == 0 {
		t.Errorf("expected non-zero ID")
	}
	if resp.ProductName != body.ProductName {
		t.Errorf("expected product_name %q, got %q", body.ProductName, resp.ProductName)
	}
	if resp.Code != body.Code {
		t.Errorf("expected code %q, got %q", body.Code, resp.Code)
	}
}

func TestCreateProductCode_InvalidCode(t *testing.T) {
	setupTestDB(t)
	router := setupRouter()

	body := ProductCodeRequest{
		ProductName: "Test Product",
		Code:        "INVALID-CODE",
	}
	payload, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/product-codes", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d, body=%s", http.StatusBadRequest, rr.Code, rr.Body.String())
	}
}

func TestCreateProductCode_MissingProductName(t *testing.T) {
	setupTestDB(t)
	router := setupRouter()

	body := ProductCodeRequest{
		ProductName: "",
		Code:        "ABCD-1234-EFGH-5678",
	}
	payload, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/api/product-codes", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d, body=%s", http.StatusBadRequest, rr.Code, rr.Body.String())
	}
}

func TestGetAndDeleteProductCodes(t *testing.T) {
	setupTestDB(t)
	router := setupRouter()

	// Create one product code first
	body := ProductCodeRequest{
		ProductName: "Test Product",
		Code:        "ZZZZ-9999-YYYY-8888",
	}
	payload, _ := json.Marshal(body)

	reqCreate := httptest.NewRequest(http.MethodPost, "/api/product-codes", bytes.NewReader(payload))
	reqCreate.Header.Set("Content-Type", "application/json")
	rrCreate := httptest.NewRecorder()
	router.ServeHTTP(rrCreate, reqCreate)

	if rrCreate.Code != http.StatusCreated {
		t.Fatalf("create expected status %d, got %d, body=%s", http.StatusCreated, rrCreate.Code, rrCreate.Body.String())
	}

	var created models.ProductCode
	if err := json.Unmarshal(rrCreate.Body.Bytes(), &created); err != nil {
		t.Fatalf("failed to unmarshal created response: %v", err)
	}

	// GET all product codes
	reqGet := httptest.NewRequest(http.MethodGet, "/api/product-codes", nil)
	rrGet := httptest.NewRecorder()
	router.ServeHTTP(rrGet, reqGet)

	if rrGet.Code != http.StatusOK {
		t.Fatalf("get expected status %d, got %d, body=%s", http.StatusOK, rrGet.Code, rrGet.Body.String())
	}

	var list []models.ProductCode
	if err := json.Unmarshal(rrGet.Body.Bytes(), &list); err != nil {
		t.Fatalf("failed to unmarshal list response: %v", err)
	}

	if len(list) != 1 {
		t.Fatalf("expected 1 product code, got %d", len(list))
	}

	// DELETE the created product code
	deleteURL := "/api/product-codes/" + strconv.Itoa(int(created.ID))
	reqDel := httptest.NewRequest(http.MethodDelete, deleteURL, nil)
	rrDel := httptest.NewRecorder()
	router.ServeHTTP(rrDel, reqDel)

	if rrDel.Code != http.StatusNoContent {
		t.Fatalf("delete expected status %d, got %d, body=%s", http.StatusNoContent, rrDel.Code, rrDel.Body.String())
	}
}


