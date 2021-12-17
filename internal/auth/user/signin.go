package user

import (
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"net/http"
	"todo/cmd/connect_db"
	auth "todo/internal/auth/jwt"
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
	pgConf := connect_db.PostgresConn
	conn := pgConf.ConnectDB()
	defer func(pgConf connect_db.PostgreSQL, conn *gorm.DB) {
		err := pgConf.CloseDB(conn)
		if err != nil {
			log.Println("SignIn: Cannot close current database")
		}
	}(pgConf, conn)

	var authDetails Authentication
	err := json.NewDecoder(r.Body).Decode(&authDetails)
	if err != nil {
		err = errors.New("error in reading body")
		setHeader(w, err, http.StatusBadRequest)
		return
	}
	b, _ := ioutil.ReadAll(r.Body)
	log.Println(b)

	if authDetails.Email == "" || authDetails.Password == "" {
		err = errors.New("email and password both are required fields")
		setHeader(w, err, http.StatusBadRequest)
		return
	}

	var authUser models.User
	conn.Where("email = ?", authDetails.Email).First(&authUser)
	if authUser.Email == "" {
		err = errors.New("username or password is incorrect")
		setHeader(w, err, http.StatusUnauthorized)
		return
	}

	check := CheckPasswordHash([]byte(authUser.Password), []byte(authDetails.Password))

	if !check {
		err = errors.New("username or password is incorrect")
		setHeader(w, err, http.StatusUnauthorized)
		return
	}

	validToken, err := auth.GenerateJWT(authUser.Email)
	if err != nil {
		err = errors.New("failed to generate token")
		setHeader(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("auth-token", validToken)
}

func CheckPasswordHash(hash []byte, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}
