package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Query(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
