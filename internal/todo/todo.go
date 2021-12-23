package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"todo/cmd/connect_db"
	"todo/internal/common/utils"
	"todo/internal/models"
)

type CreateFunc func(model, value interface{}) error
type ListFunc func() interface{}

func ListCreateTodoView(w http.ResponseWriter, r *http.Request) {
	var allowedMethods = []string{"GET", "POST"}
	if !utils.Contains(allowedMethods, r.Method) {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.Method == "GET" {
		todos := listTodo(connect_db.DBConn.GetAllTodo)
		utils.SetHeader(w, todos, http.StatusOK)
	} else if r.Method == "POST" {
		createTodo(w, r, connect_db.DBConn.Create)
	}
}

func createTodo(w http.ResponseWriter, r *http.Request, create CreateFunc) {
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		err = errors.New("error in reading body")
		utils.SetHeader(w, err, http.StatusBadRequest)
		return
	}
	if todo.Title == "" || todo.Status == "" {
		err = errors.New("title and status are required fields")
		utils.SetHeader(w, err, http.StatusBadRequest)
		return
	}

	currentUser := utils.GetRequestUser(w, r)
	todo.UserId = currentUser.Id
	todo.Date = time.Now()
	todo.Description = "ssd"
	todo.Owner = &currentUser
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()
	err = create(models.Todo{}, &todo)
	if err != nil {
		_, err = fmt.Fprintf(w, "CreateTodo: failed to create todo")
		if err != nil {
			return
		}
	}
	utils.SetHeader(w, todo, http.StatusCreated)
}

func listTodo(find ListFunc) interface{} {
	res := find()
	return res
}
