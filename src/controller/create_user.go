package controller

import (
	"fmt"

	"github.com/diogoalvesf/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/diogoalvesf/meu-primeiro-crud-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError("There are some incorrect fields")
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
