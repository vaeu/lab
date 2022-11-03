package users

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vaeu/lab/golang/microservices/rest/buchladen/users/model/users"
	"github.com/vaeu/lab/golang/microservices/rest/buchladen/users/view/services"
)

func Create(c *gin.Context) {
	var user users.User
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		// handle invalid body request
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		// handle invalid JSON request
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
