package item

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type CreateItemRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type UpdateItemRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type ListItemsResponse struct {
	Items  []*Item `json:"items"`
	Total  int64   `json:"total"`
	Limit  int     `json:"limit"`
	Offset int     `json:"offset"`
}