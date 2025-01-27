package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email       string `gorm:"unique;not null;index"`
	Password    string
	Name        string
	ResetToken  string
	TokenExpiry time.Time
}
