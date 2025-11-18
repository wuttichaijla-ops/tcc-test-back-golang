package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"test-back-golang/models"

	"github.com/gorilla/mux"
)

// In-memory storage (in production, use a database)
var (
	items    []models.Item
	comments []models.Comment
	mu       sync.RWMutex
	itemID   int
	commentID int
)

func init() {
	// Initialize with sample data
	items = []models.Item{
		{ID: 1, Name: "Sample Item 1", Description: "This is a sample item"},
		{ID: 2, Name: "Sample Item 2", Description: "Another sample item"},
	}
	itemID = 3

	comments = []models.Comment{
		{ID: 1, Author: "Blend 285", Text: "have a good day", Avatar: "B"},
	}
	commentID = 2
}

// GetItems returns all items
func GetItems(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// AddItem adds a new item
func AddItem(w http.ResponseWriter, r *http.Request) {
	var newItem models.Item
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	mu.Lock()
	newItem.ID = itemID
	itemID++
	items = append(items, newItem)
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
}

// DeleteItem deletes an item by ID
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	found := false
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// AddComment adds a new comment
func AddComment(w http.ResponseWriter, r *http.Request) {
	var newComment models.Comment
	if err := json.NewDecoder(r.Body).Decode(&newComment); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	mu.Lock()
	newComment.ID = commentID
	commentID++
	comments = append(comments, newComment)
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newComment)
}

// GetComments returns all comments
func GetComments(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comments)
}

