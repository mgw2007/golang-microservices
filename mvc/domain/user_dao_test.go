package domain

import "testing"

import "github.com/stretchr/testify/assert"

func TestGetUserNotUserFound(t *testing.T) {
	user, err := GetUser(0)
	assert.Nil(t,user,"we were not expecting a user with id 0")
	assert.NotNil(t,err,"we were  expecting an error when user with id 0")
}

func TestGetUserFound(t *testing.T) {
	user, err := GetUser(123)
	assert.NotNil(t,user,"we were  expecting a user with id 123")
	assert.Nil(t,err,"we were not expecting an error when user with id 123")
	assert.EqualValues(t,123,user.ID);
}
