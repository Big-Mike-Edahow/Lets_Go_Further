// models.go

package main

import (
	"time"
)

type Movie struct {
    Id        int64     `json:"id"`
    Created   time.Time `json:"-"` // Use the - directive to always hide the field
    Title     string    `json:"title"`
    Year      int32     `json:"year,omitempty"`    // Add the omit if empty directive
    Runtime Runtime     `json:"runtime,omitempty"`
    Genres    []string  `json:"genres,omitempty"`
    Version   int32     `json:"version"`
}
