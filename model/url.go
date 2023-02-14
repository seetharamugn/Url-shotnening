package model

import (
	"github.com/jinzhu/gorm"
)

type URL struct {
	gorm.Model
	Original  string `json:"original"`
	Shortened string `json:"shortened"`
	ExpiresAt int64  `json:"expires_at"`
}
