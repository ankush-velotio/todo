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
type ListFunc func(model interface{}) interface{}

func ListCreateTodoView(w http.ResponseWriter, r *http.Request) {
	var allowedMethods = []string{"GET", "POST"}
	if !utils.Contains(allowedMethods, r.Method) {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.Method == "GET" {
		todos := listTodo(connect_db.DBConn.FindTodo)
		j, err := json.Marshal(todos)
		_, err = fmt.Fprintf(w, string(j))
		if err != nil {
			return
		}
	} else if r.Method == "POST" {
		createTodo(w, r, connect_db.DBConn.CreateTodo)
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
	todo.UserId = currentUser.ID
	todo.Date = time.Now()
	todo.Description = "ssd"
	err = create(&models.Todo{}, &todo)
	if err != nil {
		_, err = fmt.Fprintf(w, "CreateTodo: failed to create todo")
		if err != nil {
			return
		}
	}
}

func listTodo(find ListFunc) interface{} {
	res := find(&models.Todo{})
	return res
}
