package route

import (
	"github.com/gin-gonic/gin"
	"go_web_starter/config"
	"go_web_starter/controller"
	"go_web_starter/middleware"
)

var Application *gin.Engine

func Init() {
	router := gin.New()

	if config.Config.AppConfig.Mode == "debug" {
		gin.SetMode(gin.DebugMode)
		router.Use(middleware.CostTime())
	}

	if config.Config.AppConfig.Https == true {
		router.Use(middleware.TlsHandler())
	}

	Application = router

	userController := new(controller.UserController)

	index := router.Group("/")
	{
		index.GET("/ping", userController.GetUser)
	}
}
