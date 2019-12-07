package domain

import "fmt"

var (
	users = map[int64]*User{
		123: &User{ID: 123, FirstName: "Mina", LastName: "Gerges", Email: "mina@gmail.com"},
	}
	//UserDao for user domain methods
	UserDao usersDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type usersDaoInterface interface {
	GetUser(int64) (*User, error)
}
type userDao struct{}

// GetUser find and return user struct
func (u *userDao) GetUser(userID int64) (*User, error) {
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, fmt.Errorf("user %v not exist ", userID)
}
