package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mgw2007/golang-microservices/mvc/domain"
	"github.com/mgw2007/golang-microservices/mvc/services"
	"github.com/mgw2007/golang-microservices/mvc/utils"
)

//GetUser for return users
func GetUser(c *gin.Context) {
	var err *utils.ApplicationError
	var user *domain.User
	log.Println(c.Params)

	userID, userIDErr := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if userIDErr != nil {
		err = &utils.ApplicationError{
			Message:    userIDErr.Error(),
			Code:       "user_id must be number",
			StatusCode: http.StatusBadRequest,
		}
		utils.RespondError(c, err)
		return
	}
	user, userErr := services.UsersService.GetUser(userID)
	if userErr != nil {
		utils.RespondError(c, userErr)
		return
	}
	// return user to client
	utils.Respond(c, http.StatusOK, user)

}
