package user

import (
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"todo/cmd/connect_db"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	pgConf := connect_db.PostgresConn
	conn := pgConf.ConnectDB()
	defer func(pgConf connect_db.PostgreSQL, conn *gorm.DB) {
		err := pgConf.CloseDB(conn)
		if err != nil {
			log.Println("SignUp: Cannot close current database")
		}
	}(pgConf, conn)

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		err = errors.New("error in reading body")
		setHeader(w, err, http.StatusBadRequest)
		return
	}
	var dbUser User
	conn.Where("email = ?", user.Email).First(&dbUser)

	//checks if email is already register or not
	if dbUser.Email != "" {
		err = errors.New("email already in use")
		setHeader(w, err, http.StatusConflict)
		return
	}

	user.Password, err = GenerateHashPassword([]byte(user.Password))
	if err != nil {
		log.Fatalln("error in password hash")
	}

	//insert user details in database
	conn.Create(&user)
	user.Password = ""
	setHeader(w, user, http.StatusOK)
}

func GenerateHashPassword(password []byte) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(password, 14)
	return string(bytes), err
}

func setHeader(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println("SignUp: Error in encoding the data")
	}
}
