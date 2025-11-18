package models

// Item represents an item in the system
type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Comment represents a comment in the system
type Comment struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Text   string `json:"text"`
	Avatar string `json:"avatar"`
}

