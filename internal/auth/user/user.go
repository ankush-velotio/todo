package user

import "github.com/jinzhu/gorm"

// User model
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password []byte `json:"-"`
	Active   bool   `json:"active"`
}
