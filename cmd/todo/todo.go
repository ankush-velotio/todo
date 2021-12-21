package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	auth "todo/internal/auth/jwt"
	"todo/internal/auth/user"
	"todo/internal/todo"
)

func authorizeRequest(view http.HandlerFunc) http.Handler {
	handler, err := auth.IsAuthorized(view)
	if err != nil {
		return nil
	}
	return handler
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/signup", user.SignUp)
	router.HandleFunc("/signin", user.SignIn)
	router.Handle("/todos", authorizeRequest(todo.ListCreateTodoView)).Methods("GET", "POST")
	log.Fatal(http.ListenAndServe(":8081", router))
}
