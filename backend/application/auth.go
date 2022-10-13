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

type cookie struct {
	Cookie string `json:"cookie"`
}

//acctoken과 같이 들어온 요청이 유효한지를 판단하는 미들웨어
func checkAcctoken(c *gin.Context){
	fmt.Println("미들웨어: 베가스 쿠키가 유효한지 검증합니다")
	var rw http.ResponseWriter = c.Writer
	var req *http.Request = c.Request

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("요청 바디 자체를 못 읽었음")
	}

	var hanchartCookie cookie
	if err := json.Unmarshal(data,&hanchartCookie); err != nil {
		fmt.Println(err,"제이쓴 파싱에 실패했습니다")
	}

	if hanchartCookie.Cookie == "" || hanchartCookie.Cookie == "isPKCE=true" {
		fmt.Println(err,"베가스 쿠키가 없습니다")
		con := oauth2.Config{
			ClientID: "hanchart",
			ClientSecret: "foobar",
			RedirectURL: "http://localhost:4006/callback",
			Scopes: []string{"openid", "offline"},
			Endpoint: oauth2.Endpoint{
				TokenURL: "http://localhost:8080/api/oauth2/token",
				AuthURL:  "http://localhost:8080/",
			},
		}

		pkceCodeVerifier := generateCodeVerifier(64)
		pkceCodeChallenge = generateCodeChallenge(pkceCodeVerifier)
		//1.sso통합 로그인 페이지 생성
		ssoLoginURL := con.AuthCodeURL("nuclear-tuna-plays-piano")+"&nonce=some-random-nonce&code_challenge="+pkceCodeChallenge+"&code_challenge_method=S256"
		//이 주소를 다시 프론트로 보내서 리디렉션 시켜줌
		fmt.Println(err,"SSO 쿠키가 있는지 확인하러 갑니다...")
		c.JSON(http.StatusOK, gin.H{
			"redirectionURL": ssoLoginURL,
		})
		return
	}

	//쿠키가 존재. jwt 파싱해서 유효한 토큰인지를 확인
	fmt.Println("액세스 토큰을 파싱하여 검증합니다")
	var jwtKey = []byte("sercrethatmaycontainch@r$32chars")
	tknStr := strings.Split(hanchartCookie.Cookie,"=")[1]
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			rw.WriteHeader(http.StatusUnauthorized)
			fmt.Println("권한 없는 이상한 토큰입니다.")
			return
		}
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Println("아무튼 잘못된 요청입니다.")
		return
	}

	if !tkn.Valid {
		rw.WriteHeader(http.StatusUnauthorized)
		fmt.Println("권한이 없는 요청입니다.")
		return
	}

	fmt.Println("유효한 토큰입니다. api에 요청을 보냅니다...")
	c.Next()
}