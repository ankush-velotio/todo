package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// User model
type User struct {
	Id    int    `orm:"pk; auto"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Active   bool   `json:"active"`
}

func init() {
	orm.RegisterModel(new(User))
}
