package models

import (
	"gorm.io/gorm"
)

type InOut struct {
    gorm.Model
    UserID        uint      `json:"user_id"` // Foreign key
    CreatedAt     string    `json:"created_at"`
    LatestCommand string    `json:"latest_command"`
}