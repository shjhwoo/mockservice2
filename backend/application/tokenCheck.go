package application

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kataras/jwt"
)

type Token struct {
	AccessToken string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

var jwtKey = []byte("sercrethatmaycontainch@r$32chars")

func tokenCheckHandler(c *gin.Context) {
	fmt.Println("서비스 토큰의 유효성을 검증합니다")
	//var rw http.ResponseWriter = c.Writer
	var req *http.Request = c.Request
	//액세스 토큰부터 검증한다 .
	data, err := ioutil.ReadAll(req.Body);
	if err != nil {
		fmt.Println("요청 바디를 읽어오지 못함")
	}
	var token Token
	if err := json.Unmarshal(data, &token); err != nil {
		fmt.Println("제이슨 파싱 실패")
	}
	fmt.Println(token,"토큰확인")

	//액세스 토큰 검증
	tknStr := token.AccessToken
	tkn, err := jwt.Verify(jwt.HS256, jwtKey, []byte(tknStr))
	if err != nil {
		fmt.Println(err.Error(),"액세스토큰검증")
		if err == jwt.ErrTokenSignature {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "토큰생성에 사용한 서명이 잘못되었습니다",
			})
			return
		}

		if err.Error() == "jwt: token is empty" || err == jwt.ErrExpired {
			//만약에 만료된 에러라면 보내준 리프레시 토큰을 검증
			tknStr := token.RefreshToken
			tkn, err := jwt.Verify(jwt.HS256, jwtKey, []byte(tknStr))
			if err != nil {
				fmt.Println(err,"리프레시쿠키검증")
				if err == jwt.ErrTokenSignature {
					c.JSON(http.StatusUnauthorized, gin.H{
						"message": "토큰생성에 사용한 서명이 잘못되었습니다",
					})
					return
				}
				//리프레시 토큰도 없거나 만료됨.
				if err.Error() == "jwt: token is empty" || err == jwt.ErrExpired {
					c.Redirect(http.StatusSeeOther,"/checksso") //303
					return
				}
			}

			fmt.Println("리프레시 토큰이 아직 살아 있으니, access token을 새로 만들어 준다.")
			serviceAccessToken, err := jwt.Sign(jwt.HS256, jwtKey, tkn.Payload, jwt.MaxAge(15*time.Minute))
			c.JSON(http.StatusCreated, gin.H{
				"message":"액세스 토큰이 만료되어 새로 발급했습니다",
				"accessToken": string(serviceAccessToken[:]),
			})
			return
			// cookie, err := req.Cookie("vegasRefreshToken");
			// if err != nil {
			// 	fmt.Println(err,"리프레시쿠키검증")
			// 	if err.Error() == "http: named cookie not present" {
			// 		c.Redirect(http.StatusSeeOther,"/checksso") //303
			// 		return
			// 	}
			// }
			// refreshToken := strings.Split(cookie.Value,"=")[1];
			// tkn, err := jwt.Verify(jwt.HS256, jwtKey, []byte(refreshToken))
			// if err != nil {
			// 	fmt.Println(err,"리프레시토큰검증")
			// }
			// if err == jwt.ErrExpired{
			// 	//리프레시 토큰마저도 만료됨.
			// 	//sso 확인하러 간다.
			// 	c.Redirect(http.StatusSeeOther,"/checksso") //303
			// 	return
			// }
			// fmt.Println("리프레시 토큰이 아직 살아 있으니, access token을 새로 만들어 준다.")
			// serviceAccessToken, err := jwt.Sign(jwt.HS256, jwtKey, tkn.Payload, jwt.MaxAge(15*time.Minute))
			// c.JSON(http.StatusCreated, gin.H{
			// 	"message":"액세스 토큰이 만료되어 새로 발급했습니다",
			// 	"accessToken": string(serviceAccessToken[:]),
			// })
			// return
		}
	}
	fmt.Println("액세스 토큰",tkn.Token)
	c.JSON(http.StatusOK,gin.H{
		"message":"유효한 액세스 토큰입니다",
	})
}

