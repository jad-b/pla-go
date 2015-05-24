package auth

import (
	"reflect"
)

// User represents a user in the system
type User struct {
	Name     string
	Password string
	Email    string
}

// NewUser creates a new User
func NewUser(name, password, email string) *User {
	return &User{name, password, email}
}

func (user *User) equal(other *User) bool {
	return user.Name == other.Name &&
		reflect.DeepEqual(user.Password, other.Password) &&
		user.Email == other.Email
}
