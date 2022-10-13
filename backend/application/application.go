package application

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(
		cors.Config{
			AllowOrigins:     []string{"http://localhost:3006"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Accept", "Content-Type", "X-Requested-With", "withCredentials"},
			AllowCredentials: true,
			MaxAge:           0,
		},
	))

	router.GET("/", healthCheck)
	router.GET("/checksso", checksso)
	router.GET("/sso/login", ssologinHandler)
	router.GET("/callback", callbackHandler)
	router.POST("/callback", callbackHandler)
	router.POST("/api/chart", getchartservice)
	router.POST("/logout", logoutHandler)
	router.POST("/slo", sloHandler)
	router.POST("/checkservicetkn", tokenCheckHandler)
	router.POST("/refresh", tokenRefreshHandler)
	// 미들웨어 추가
	//acctoken과 같이 들어온 요청이 유효한지를 판단하는 미들웨어 작성해야함
	router.GET("/all", getnormalservice)
	return router
}
