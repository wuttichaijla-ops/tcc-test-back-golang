package models

// ProductCode represents a product and its 16-character code (XXXX-XXXX-XXXX-XXXX)
type ProductCode struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	ProductName string `json:"product_name" gorm:"size:255;not null"`
	Code        string `json:"code" gorm:"size:19;uniqueIndex;not null"` // e.g. ABCD-1234-EFGH-5678
}

