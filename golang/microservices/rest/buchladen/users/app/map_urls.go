package app

import "github.com/vaeu/lab/golang/microservices/rest/buchladen/users/controller"

func mapURLs() {
	router.GET("/ping", controller.Ping)
}
