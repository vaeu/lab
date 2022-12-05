package app

import "github.com/gin-gonic/gin"

var router = gin.Default()

func Start() {
	mapURLs()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Run(":8080")
}
