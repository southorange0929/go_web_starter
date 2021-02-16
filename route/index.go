package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var Application *gin.Engine

func Init() {
	router := gin.New()

	Application = router

	router.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})
}
