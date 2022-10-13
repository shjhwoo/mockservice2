package application

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kataras/jwt"
)

type accessToken struct {
	AccessToken string `json:"accessToken"`
}

func getchartservice(c * gin.Context) {
	fmt.Println("차트 열람 요청이 들어왔습니다")
	var req *http.Request = c.Request

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("요청 바디를 읽을 수 없습니다",err);
		return
	}

	var accessToken accessToken
	if err := json.Unmarshal(data, &accessToken); err != nil {
		fmt.Println("제이슨 형식으로 파싱 실패",err);
		return
	}
	fmt.Println(accessToken.AccessToken)
	accToken, err := jwt.Verify(jwt.HS256, jwtKey,[]byte(accessToken.AccessToken))

	fmt.Println(accToken,"토큰확인")

	if err !=nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "액세스 토큰 권한이 없습니다",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":"2022-09-08진료기록부",
		"isDoctor": true,
	})
}