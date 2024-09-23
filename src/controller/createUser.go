package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lyalbat/crud-golang/src/configurations/handle_errors"
	"github.com/lyalbat/crud-golang/src/controller/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		handle_error := handle_errors.GenerateBadRequestErr(
			fmt.Sprintf("Incorrect fields. error=%s", err.Error()))
		c.JSON(handle_error.Code, handle_error)
		return
	}
	fmt.Println(userRequest)
}
