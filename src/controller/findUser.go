package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lyalbat/crud-golang/src/configurations/handle_errors"
)

func FindUserById(c *gin.Context) {
	err := handle_errors.GenerateBadRequestErr("caiu na pegadinha do malandro")
	c.JSON(err.Code, err)
}

func FindUserByEmail(c *gin.Context) {}
