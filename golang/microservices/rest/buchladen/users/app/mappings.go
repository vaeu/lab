package app

import (
	"github.com/vaeu/lab/golang/microservices/rest/buchladen/users/controller/ping"
	"github.com/vaeu/lab/golang/microservices/rest/buchladen/users/controller/users"
)

func mapURLs() {
	router.GET("/ping", ping.Query)
	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	router.GET("/users/search", users.Search)
	router.PUT("/users/:user_id", users.Update)
}
