package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/mgw2007/golang-microservices/mvc/domain"
	"github.com/mgw2007/golang-microservices/mvc/services"
	"github.com/mgw2007/golang-microservices/mvc/utils"
)

//GetUser for return users
func GetUser(w http.ResponseWriter, r *http.Request) {
	var err *utils.ApplicationError
	var user *domain.User

	userID, userIDErr := strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 64)

	if userIDErr != nil {
		err = &utils.ApplicationError{
			Message:    userIDErr.Error(),
			Code:       "user_id must be number",
			StatusCode: http.StatusBadRequest,
		}
		jsonerr, _ := json.Marshal(err)
		w.WriteHeader(err.StatusCode)
		w.Write(jsonerr)
		return
	}
	user, userErr := services.GetUser(userID)
	if userErr != nil {
		jsonerr, _ := json.Marshal(userErr)
		w.WriteHeader(userErr.StatusCode)
		w.Write(jsonerr)
		return
	}
	// return user to client
	jsonvalue, _ := json.Marshal(user)
	w.Header().Set("content-type", "application/json")
	w.Write(jsonvalue)

}
