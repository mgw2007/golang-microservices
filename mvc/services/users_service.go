package services

import "github.com/mgw2007/golang-microservices/mvc/domain"

import "github.com/mgw2007/golang-microservices/mvc/utils"

import "net/http"

type usersService struct {
}

var (
	//UsersService for user service operations
	UsersService usersService
)

// GetUser find and return user struct
func (u *usersService) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	user, err := domain.UserDao.GetUser(userID)
	if err != nil {
		return nil, &utils.ApplicationError{
			Message:    err.Error(),
			Code:       "user_not_exist",
			StatusCode: http.StatusNotFound,
		}
	}
	return user, nil
}
