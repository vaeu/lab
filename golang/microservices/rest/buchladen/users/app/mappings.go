package app

import "github.com/vaeu/lab/golang/microservices/rest/buchladen/users/controller"

func mapURLs() {
	router.GET("/ping", controller.Ping)
	router.POST("/users", controller.CreateUser)
	router.GET("/users/:user_id", controller.GetUser)
	router.GET("/users/search", controller.SearchUser)
}
