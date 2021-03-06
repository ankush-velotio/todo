package models

import "github.com/jinzhu/gorm"

// User model
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"primaryKey" json:"email"`
	Password string `json:"password"`
	Active   bool   `json:"active"`
}
