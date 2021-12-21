package user

import (
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"todo/cmd/connect_db"
	"todo/internal/common/utils"
	"todo/internal/models"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	pgConf := connect_db.DBConn
	conn := pgConf.ConnectDB()
	defer func(pgConf connect_db.DB, conn *gorm.DB) {
		err := pgConf.CloseDB(conn)
		if err != nil {
			log.Println("SignUp: Cannot close current database")
		}
	}(pgConf, conn)

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		err = errors.New("error in reading body")
		utils.SetHeader(w, err, http.StatusBadRequest)
		return
	}
	var dbUser models.User
	conn.Where("email = ?", user.Email).First(&dbUser)

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

	//insert user details in database
	conn.Create(&user)
	user.Password = ""
	utils.SetHeader(w, user, http.StatusOK)
}

func GenerateHashPassword(password []byte) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(password, 14)
	return string(bytes), err
}
