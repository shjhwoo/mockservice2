package application

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func ssologinHandler(c *gin.Context){
	var w http.ResponseWriter = c.Writer
	var r *http.Request = c.Request
	fmt.Println("요청확인")
	//클라이언트 정보에 따라 통합 로그인 페이지 접속 url을 만들어 준 후 그쪽으로 사용자를 보내준다.
	//0.클라이언트 설정
	con := oauth2.Config{
		ClientID: "vegas",
		ClientSecret: "foobar",
		RedirectURL: "http://localhost:3006/callback",
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

	fmt.Println(ssoLoginURL,"요청url")
	//2.사용자 리디렉션
	http.Redirect(w, r, ssoLoginURL, http.StatusSeeOther)
}
