package application

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(
		cors.Config{
			AllowOrigins:   []string{"http://localhost:3006"},
			AllowMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
			AllowHeaders:   []string{"Origin", "Accept", "Content-Type", "X-Requested-With","withCredentials"},
			AllowCredentials: true,
			MaxAge:           0,
		},
	))

	router.GET("/", healthCheck)
	router.POST("/checkcookie", checksso)
	router.GET("/sso/login", ssologinHandler)
	router.GET("/callback", callbackHandler)
	router.POST("/callback", callbackHandler)
	router.GET("/admin", getadminservice)
	router.GET("/all", getnormalservice)
	return router
}

