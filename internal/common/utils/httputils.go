package utils

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"os"
)

var jwtSigningKey = []byte(os.Getenv("SECRET_KEY"))

func SetHeader(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println("SignUp: Error in encoding the data")
	}
}

func GetDataFromAuthToken(w http.ResponseWriter, r *http.Request, key string) interface{} {
	if r.Header["Token"] != nil {

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			return jwtSigningKey, nil
		})
		if err != nil {
			_, err := fmt.Fprintf(w, err.Error())
			if err != nil {
				return nil
			}
		}

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			return claims[key]
		}
	} else {
		_, err := fmt.Fprintf(w, "No Authorization Token provided")
		if err != nil {
			return nil
		}
	}
	return nil
}
