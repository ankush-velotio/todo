package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type Status string

const (
	Completed  Status = "completed"
	Incomplete Status = "incomplete"
	Canceled   Status = "canceled"
)

/*
	_Todo model
*/
type Todo struct {
	Id    int    `orm:"pk; auto"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Date         time.Time `json:"date"`
	IsBookmarked bool      `json:"isBookmarked"`
	UserId    int
	Owner        *User   `json:"owner" orm:"rel(fk)"`
	Editors      []*User `json:"editors" orm:"rel(m2m);rel_table(todo/internal/models.TodoUsers)"`
	Status       Status `json:"status"`
}

func init() {
	orm.RegisterModel(new(Todo))
}
