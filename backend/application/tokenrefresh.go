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

type refreshToken struct {
	AccessToken string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func tokenRefreshHandler(c *gin.Context) {
	fmt.Println("액세스 토큰 리프레시 요청이 들어왔습니다")
	var req *http.Request = c.Request

	//가지고 온 리프레시 토큰을 파싱해서 받아온다. 
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("요청 바디를 읽을 수 없습니다",err);
	}

	var token refreshToken
	if err := json.Unmarshal(data, &token); err != nil {
		fmt.Println("제이슨 형식으로 파싱 실패",err);
	}

	refreshToken := token.RefreshToken
	accessToken := token.AccessToken

	//액세스 토큰이 비었어도 검증하자
	if accessToken == "" {
		refToken, err := jwt.Verify(jwt.HS256,jwtKey, []byte(refreshToken))
		if err != nil {
			fmt.Println(err.Error(),"리프레시토큰검증")
			if err != jwt.ErrTokenSignature {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "토큰생성에 사용한 서명이 잘못되었습니다",
				})
				return
			}
			if err.Error() == "jwt: token is empty" || err == jwt.ErrExpired {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "리프레시 토큰이 만료되었습니다",
				})
				return
			}
		}
		//유효한 리프레시 토큰임이 증명됨. 이를 이용해 새로운 액세스 토큰을 발급하고 돌려준다
		var sharedKey = []byte("sercrethatmaycontainch@r$32chars")
		serviceAccessToken, err := jwt.Sign(jwt.HS256,sharedKey,refToken.Payload,jwt.MaxAge(15*time.Minute))
		if err != nil {
			fmt.Println(err,"???")
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "액세스 토큰 생성 실패",
			})
			return
		}

		//액세스 토큰 정상적으로 발급완료
		c.JSON(http.StatusCreated, gin.H{
			"message": "액세스 토큰이 재발급되었습니다",
			"accessToken": string(serviceAccessToken[:]),
		})
		return
	}
	
	//만료되었다고 증명되면 해당 리프레시 토큰을 검증한다. 
	accToken, err := jwt.Verify(jwt.HS256, jwtKey,[]byte(accessToken))
	if err == nil {
		fmt.Println("액세스 토큰 정상적으로 쓸수있습니다. 그냥 쓰세요.",string(accToken.Token[:]))
		c.JSON(http.StatusCreated, gin.H{
			"message": "이미 사용가능한 액세스 토큰인데 또 리프레시 요청을 보내셨군요. 그냥 쓰세요.",
			"accessToken": string(accToken.Token[:]),
		})
		return
	}

	if err.Error() == "jwt: token is empty" || err == jwt.ErrExpired {
		//액세스 토큰이 만료되었다는 뜻. 여기서 리프레시 토큰을 검증
		refToken, err := jwt.Verify(jwt.HS256,jwtKey, []byte(refreshToken))
		if err != nil {
			fmt.Println(err.Error(),"리프레시토큰검증")
			if err != jwt.ErrTokenSignature {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "토큰생성에 사용한 서명이 잘못되었습니다",
				})
				return
			}
			if err.Error() == "jwt: token is empty" || err == jwt.ErrExpired {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "리프레시 토큰이 만료되었습니다",
				})
				return
			}
		}
		//유효한 리프레시 토큰임이 증명됨. 이를 이용해 새로운 액세스 토큰을 발급하고 돌려준다
		var sharedKey = []byte("sercrethatmaycontainch@r$32chars")
		serviceAccessToken, err := jwt.Sign(jwt.HS256,sharedKey,refToken.Payload,jwt.MaxAge(15*time.Minute))
		if err != nil {
			fmt.Println(err,"???")
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "액세스 토큰 생성 실패",
			})
			return
		}

		//액세스 토큰 정상적으로 발급완료
		c.JSON(http.StatusCreated, gin.H{
			"message": "액세스 토큰이 재발급되었습니다",
			"accessToken": string(serviceAccessToken[:]),
		})
		return
	}

	 c.JSON(http.StatusInternalServerError, gin.H{
		"message": "서버에 문제가 발생했습니다.",
	 })
}