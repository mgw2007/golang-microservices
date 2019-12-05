package domain

import "fmt"

var (
	users = map[int64]*User{
		123: &User{ID: 123, FirstName: "Mina", LastName: "Gerges", Email: "mina@gmail.com"},
	}
)

// GetUser find and return user struct
func GetUser(userID int64) (*User, error) {
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, fmt.Errorf("user %v not exist ", userID)
}
