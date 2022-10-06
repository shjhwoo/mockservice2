package application

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", healthCheck)
	router.GET("/sso/login", ssologinHandler)
	router.GET("/callback", callbackHandler)
	router.POST("/callback", callbackHandler)
	router.GET("/admin", getadminservice)
	router.GET("/all", getnormalservice)
	return router
}


