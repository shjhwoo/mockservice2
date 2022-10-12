package application

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sloHandler(c *gin.Context) {
	var rw http.ResponseWriter = c.Writer
	var req *http.Request = c.Request

	// data, err := ioutil.ReadAll(req.Body)
	// if err != nil {
	// 	fmt.Println("요청 바디 자체를 확인할 수 없습니다")
	// 	return
	// }
	// var cookie cookie
	// if err := json.Unmarshal(data, &cookie); err != nil {
	// 	fmt.Println("제이슨 형식으로 쿠키 파싱하는 데 실패함")
	// 	return
	// }
	// fmt.Println(strings.Split(cookie.Cookie, "; "), "쿠키 값 확인하세요")
	// cookiename := strings.Split(cookie.Cookie, "=")[0]
	// if cookiename != "vegasAccessToken" {
	// 	fmt.Println("베가스 서비스 쿠키가 아닙니다.")
	// 	return
	// }
	// fmt.Println(rw)
	// var tknStr string
	// for _, cookieValue := range strings.Split(cookie.Cookie, "; ") {
	// 	if strings.Split(cookieValue, "=")[0] == "vegasAccessToken" {
	// 		tknStr = strings.Split(cookieValue, "=")[1]
	// 	}
	// }

	fmt.Println(req.Body, "--------------------------")

	c.SetCookie("test1", "test1", 15*60, "/", "localhost", false, false)

	// fmt.Println(tknStr)

	c.JSON(http.StatusCreated, rw)
}
