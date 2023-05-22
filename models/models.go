package models

import (
	"time"
)

type Blog struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Body        string    `json:"body,required"`
	Created_At  time.Time `json:"created_at,omitempty"`
	Updated_At  time.Time `json:"updated_at,omitempty"`
}

type BlogPatch struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Body        *string `json:"body,omitempty"`
}
