package users

import (
	"net/http"
	"strconv"

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
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func Get(c *gin.Context) {
	uID, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		err := errors.NewBadRequest("user ID is expected to be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(uID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func Search(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}

func Update(c *gin.Context) {
	uID, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		err := errors.NewBadRequest("user ID is expected to be a number")
		c.JSON(err.Status, err)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequest("invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.ID = uID

	result, nerr := services.UpdateUser(user)
	if nerr != nil {
		c.JSON(nerr.Status, nerr)
		return
	}
	c.JSON(http.StatusOK, result)
}
