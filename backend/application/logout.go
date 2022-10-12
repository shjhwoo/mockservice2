package application

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func logoutHandler(c *gin.Context) {
	fmt.Println("로그아웃 요청이 들어왔습니다")
	var rw http.ResponseWriter = c.Writer
	var req *http.Request = c.Request

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("요청 바디 자체를 확인할 수 없습니다")
		return
	}
	var cookie cookie
	if err := json.Unmarshal(data, &cookie); err != nil {
		fmt.Println("제이슨 형식으로 쿠키 파싱하는 데 실패함")
		return
	}

	tknStr := strings.Split(cookie.Cookie, "=")[1]

	fmt.Println("유효한 토큰인지 검증 시작합니다...")
	claims := &Claims{}
	var jwtKey = []byte("sercrethatmaycontainch@r$32chars")

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			rw.WriteHeader(http.StatusUnauthorized)
			fmt.Println("권한 없는 이상한 토큰")
			return
		}
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Println("아무튼 잘못된 요청", err)
		return
	}

	if !tkn.Valid {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	fmt.Println("받은 쿠키는 유효합니다. 이제 진짜로 SSO쿠키를 파괴할 수 있는 url을 제공합니다")

	//그런 다음에 SSO 세션을 체크하는 url을 사용자에게 되돌려준다.
	//sso 쿠키가 살아있는지 확인해야지...
	con := oauth2.Config{
		ClientID:     "vegas",
		ClientSecret: "foobar",
		RedirectURL:  "http://localhost:3006/callback",
		Scopes:       []string{"openid", "offline"},
		Endpoint: oauth2.Endpoint{
			TokenURL: "http://localhost:8080/api/oauth2/token",
			AuthURL:  "http://localhost:8080/",
		},
	}

	pkceCodeVerifier := generateCodeVerifier(64)
	pkceCodeChallenge = generateCodeChallenge(pkceCodeVerifier)

	//로그아웃 요청
	ssoLogoutURL := con.AuthCodeURL("nuclear-tuna-plays-piano") + "&nonce=some-random-nonce&code_challenge=" + pkceCodeChallenge + "&code_challenge_method=S256&logout=true"

	//이 주소를 다시 프론트로 보내서 리디렉션 시켜줌
	c.JSON(http.StatusOK, gin.H{
		"redirectionURL": ssoLogoutURL,
	})
}
