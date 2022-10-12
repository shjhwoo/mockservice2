package application

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
)

type AuthTokenClaims struct {
	Name               string `json:"name"` // 유저 이름
	UserId             string `json:"userid"`
	OrgId              string `json:"orgid"`
	Department         string `json:"department"`
	Services           string `json:"services"`
	Roles              string `json:"roles"` // 유저 역할
	jwt.StandardClaims        // 표준 토큰 Claims
}

func GetAccessToken(name string, userId string, orgId string, department string, roles string, services string) (string, error) {
	at := AuthTokenClaims{
		Name:       name,
		UserId:     userId,
		OrgId:      orgId,
		Department: department,
		Roles:      roles,
		Services:   services,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Minute * 15)), // 만료시간 15분
		},
	}

	atoken := jwt.NewWithClaims(jwt.SigningMethodHS256, &at)
	signedAuthToken, err := atoken.SignedString([]byte("accessToken"))
	return signedAuthToken, err
}

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
