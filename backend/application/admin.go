package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getchartservice(c * gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":"2022-09-08진료기록부",
		"isDoctor": true,
	})
}