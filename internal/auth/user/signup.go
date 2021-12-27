package user

import (
	"encoding/json"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
	db "todo/cmd/connect_db"
	"todo/internal/common/utils"
	"todo/internal/models"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		err = errors.New("error in reading body")
		utils.SetHeader(w, err, http.StatusBadRequest)
		return
	}
	var dbUser models.User
	dbUser = db.DBConn.FindUser(user.Email).(models.User)

	//checks if email is already register or not
	if dbUser.Email != "" {
		err = errors.New("email already in use")
		utils.SetHeader(w, err, http.StatusConflict)
		return
	}

	user.Password, err = GenerateHashPassword([]byte(user.Password))
	if err != nil {
		log.Fatalln("error in password hash")
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	//insert user details in database
	err = db.DBConn.Create(models.User{}, &user)
	if err != nil {
		err = errors.New("unable to create your account")
		utils.SetHeader(w, err, http.StatusInternalServerError)
	}
	data, err := utils.GetCustomJSON(user, "password")
	if err != nil {
		return
	}
	utils.SetHeader(w, data, http.StatusOK)
}

func GenerateHashPassword(password []byte) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(password, 14)
	return string(bytes), err
}
