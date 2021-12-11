package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"testing"
	"time"
	auth "todo/internal/auth/jwt"
)

func TestGenerateJWT(t *testing.T) {
	cases := []string{
		"test@email.com",
		"user@example.com",
	}

	for _, email := range cases {
		token, err := auth.GenerateJWT(email)
		if err != nil {
			t.Errorf("JWT token generation failed")
		}
		signingKey := []byte("ehe76bsgd7")
		tkn, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) { return signingKey, nil })
		if claims, ok := tkn.Claims.(jwt.MapClaims); ok && tkn.Valid {
			if claims["exp"].(float64) < float64(time.Now().Unix()) {
				t.Errorf("Token is expired")
			}
		}
	}
}

func TestIsAuthorized(t *testing.T) {
	cases := []http.HandlerFunc{
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Test token auth")
		},
	}

	for _, handlrFunc := range cases {
		_, err := auth.IsAuthorized(handlrFunc)
		if err != nil {
			t.Errorf("Error in authorization %s", err.Error())
		}
	}
}
