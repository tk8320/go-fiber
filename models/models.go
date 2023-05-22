package models

import (
	"time"
)

// swagger:model Blog
type Blog struct {
	// The ID of Post

	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Body        string    `json:"body,required"`
	Created_At  time.Time `json:"created_at,omitempty"`
	Updated_At  time.Time `json:"updated_at,omitempty"`
}

// swagger:model BlogPatch
type BlogPatch struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Body        *string `json:"body,omitempty"`
}
