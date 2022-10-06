package application

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthCheck(c *gin.Context) {
	var rw http.ResponseWriter = c.Writer
	fmt.Fprint(rw, "OK")
}