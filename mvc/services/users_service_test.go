package services

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/mgw2007/golang-microservices/mvc/domain"
	"github.com/stretchr/testify/assert"
)

var (
	userDoaMock     usersDaoMock
	getUserFunction func(userID int64) (*domain.User, error)
)

func init() {
	domain.UserDao = &usersDaoMock{}
}

type usersDaoMock struct{}

func (u usersDaoMock) GetUser(userID int64) (*domain.User, error) {
	return getUserFunction(userID)
}
func TestGetUserNotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userID int64) (*domain.User, error) {
		return nil, fmt.Errorf("user %v not exist ", userID)
	}
	user, err := UsersService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user_not_exist", err.Code)
}

func TestGetUserInDatabase(t *testing.T) {
	getUserFunction = func(userID int64) (*domain.User, error) {
		return &domain.User{
			ID: uint64(userID),
		}, nil
	}
	testUserID := int64(121)
	user, err := UsersService.GetUser(testUserID)
	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.EqualValues(t, testUserID, user.ID)
}
