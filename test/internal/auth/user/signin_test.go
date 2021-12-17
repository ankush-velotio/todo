package user

import (
	"gotest.tools/v3/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todo/internal/auth/user"
)

func TestSignIn(t *testing.T) {
	request, _ := http.NewRequest("GET", "/signup", nil)
	response := httptest.NewRecorder()

	user.SignUp(response, request)

	assert.Assert(t, response.Code == http.StatusMethodNotAllowed)

	// Signin with valid email and password
	var data = `{
		"email": "test@email.com",
		"password": "pass"
	}`

	request, _ = http.NewRequest("POST", "/signin", strings.NewReader(data))
	response = httptest.NewRecorder()
	user.SignIn(response, request)

	assert.Assert(t, response.Code == http.StatusOK)

	// Signin with insufficient credentials
	data = `{
		"email": "test@email.com",
	}`

	request, _ = http.NewRequest("POST", "/signin", strings.NewReader(data))
	response = httptest.NewRecorder()
	user.SignIn(response, request)

	assert.Assert(t, response.Code == http.StatusBadRequest)

	// Signin with wrong email
	data = `{
		"email": "testdemo@email.com",
		"password": "pass"
	}`

	request, _ = http.NewRequest("POST", "/signin", strings.NewReader(data))
	response = httptest.NewRecorder()
	user.SignIn(response, request)

	assert.Assert(t, response.Code == http.StatusUnauthorized)

	// Signin with wrong password
	data = `{
		"email": "test@email.com",
		"password": "wrongPass"
	}`

	request, _ = http.NewRequest("POST", "/signin", strings.NewReader(data))
	response = httptest.NewRecorder()
	user.SignIn(response, request)

	assert.Assert(t, response.Code == http.StatusUnauthorized)
}
