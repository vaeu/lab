package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vaeu/lab/golang/microservices/rest/buchladen/users/model/users"
	"github.com/vaeu/lab/golang/microservices/rest/buchladen/users/utils/errors"
	"github.com/vaeu/lab/golang/microservices/rest/buchladen/users/view/services"
)

func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequest("invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, err := services.CreateUser(user)
	if err != nil {
		// handle invalid user creation
		return
	}
	c.JSON(http.StatusCreated, result)
}

func Get(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}

func Search(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}
