package user

import (
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"todo/cmd/connect_db"
	auth "todo/internal/auth/jwt"
	"todo/internal/common/utils"
	"todo/internal/models"
)

// Authentication type is just for getting login credentials
type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	pgConf := connect_db.DBConn
	conn := pgConf.ConnectDB()
	defer func(pgConf connect_db.DB, conn *gorm.DB) {
		err := pgConf.CloseDB(conn)
		if err != nil {
			log.Println("SignIn: Cannot close current database")
		}
	}(pgConf, conn)

	var authDetails Authentication
	err := json.NewDecoder(r.Body).Decode(&authDetails)
	if err != nil {
		err = errors.New("error in reading body")
		utils.SetHeader(w, err, http.StatusBadRequest)
		return
	}

	if authDetails.Email == "" || authDetails.Password == "" {
		err = errors.New("email and password both are required fields")
		utils.SetHeader(w, err, http.StatusBadRequest)
		return
	}

	var authUser models.User
	conn.Where("email = ?", authDetails.Email).First(&authUser)
	if authUser.Email == "" {
		err = errors.New("username or password is incorrect")
		utils.SetHeader(w, err, http.StatusUnauthorized)
		return
	}

	check := CheckPasswordHash([]byte(authUser.Password), []byte(authDetails.Password))

	if !check {
		err = errors.New("username or password is incorrect")
		utils.SetHeader(w, err, http.StatusUnauthorized)
		return
	}

	validToken, err := auth.GenerateJWT(authUser.Email)
	if err != nil {
		err = errors.New("failed to generate token")
		utils.SetHeader(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("auth-token", validToken)
}

func CheckPasswordHash(hash []byte, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}
