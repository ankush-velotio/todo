package user

import (
	"bytes"
	"gotest.tools/v3/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/internal/auth/user"
)

func TestSignUp(t *testing.T) {
	request, _ := http.NewRequest("GET", "/signup", nil)
	response := httptest.NewRecorder()

	user.SignUp(response, request)

	assert.Assert(t, response.Code == http.StatusMethodNotAllowed)

	// Register new user
	var data = []byte(`{
		"name": "Ankush",
		"email": "ankush@gmail.com",
		"password": "mypass",
		"active": true
	}`)

	request, _ = http.NewRequest("POST", "/signup", bytes.NewBuffer(data))
	response = httptest.NewRecorder()
	user.SignUp(response, request)

	assert.Assert(t, response.Code == http.StatusOK)

	// Register new user with the already registered email ID
	request, _ = http.NewRequest("POST", "/signup", bytes.NewBuffer(data))
	response = httptest.NewRecorder()
	user.SignUp(response, request)

	assert.Assert(t, response.Code == http.StatusConflict)

	// Register user with invalid payload - should return 400 Bad Request
	data = []byte(`{
		"name": "Ankush",
		"email": "ankush@gmail.com",
	}`)

	request, _ = http.NewRequest("POST", "/signup", bytes.NewBuffer(data))
	response = httptest.NewRecorder()
	user.SignUp(response, request)

	assert.Assert(t, response.Code == http.StatusBadRequest)
}
