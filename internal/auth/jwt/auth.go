package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSigningKey = []byte(os.Getenv("SECRET_KEY"))

func GenerateJWT(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["aud"] = "todo.io"
	claims["iss"] = "jwtgo.io"
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	tokenString, err := token.SignedString(jwtSigningKey)

	if err != nil {
		_ = fmt.Errorf("something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

// IsAuthorized Middleware for verifying the JWT token
func IsAuthorized(endpoint http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {

			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("invalid Signing Method. error in parsing token")
				}
				// verify audience claim
				aud := "todo.io"
				checkAudience := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
				if !checkAudience {
					return nil, fmt.Errorf("invalid audience")
				}
				// verify issuer claim
				iss := "jwtgo.io"
				checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
				if !checkIss {
					return nil, fmt.Errorf("invalid issuer")
				}

				return jwtSigningKey, nil
			})
			if err != nil {
				_, err := fmt.Fprintf(w, err.Error())
				if err != nil {
					return
				}
			}

			if token.Valid {
				endpoint(w, r)
			}
		} else {
			_, err := fmt.Fprintf(w, "No Authorization Token provided")
			if err != nil {
				return
			}
		}
	})
}
