package route

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"go_web_starter/config"
	"go_web_starter/controller"
)

var Application *gin.Engine

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     ":443",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)
		if err != nil {
			return
		}
		c.Next()
	}
}

func Init() {
	router := gin.New()
	if config.Config.AppConfig.Https == true {
		router.Use(TlsHandler())
	}

	Application = router

	userController := new(controller.UserController)

	index := router.Group("/")
	{
		index.GET("/ping",userController.GetUser)
	}
}
