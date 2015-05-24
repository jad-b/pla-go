package auth

import (
	// "html/template"
	"net/http"
)

// CreateUserHandler creates a user.
func CreateUserHandler(w http.ResponseWriter, r *http.Request) *User {
	name, password, email := r.PostFormValue("username"),
		r.PostFormValue("password"),
		r.PostFormValue("email")
	return &User{name, password, email}
}
