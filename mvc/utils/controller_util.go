package utils

import "github.com/gin-gonic/gin"

//Respond func to send response with checking type
func Respond(c *gin.Context, status int, body interface{}) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(status, body)
	} else {
		c.JSON(status, body)
	}
}

//RespondError func to send response with checking type
func RespondError(c *gin.Context, err *ApplicationError) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(err.StatusCode, err)
	} else {
		c.JSON(err.StatusCode, err)
	}
}
