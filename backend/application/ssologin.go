package application

import (
	"net/http"

	"golang.org/x/oauth2"
)

func ssologinHandler(w http.ResponseWriter, r *http.Request){
	//클라이언트 정보에 따라 통합 로그인 페이지 접속 url을 만들어 준 후 그쪽으로 사용자를 보내준다.
	//0.클라이언트 설정
	c := oauth2.Config{
		ClientID: "vegas",
		ClientSecret: "foobar",
		RedirectURL: "http://localhost:3006/callback",
		Scopes: []string{"openid", "offline"},
		Endpoint: oauth2.Endpoint{
			TokenURL: "http://localhost:8080/oauth2/token",
			AuthURL:  "http://localhost:8080/oauth2/auth",
		},
	}

	//1.sso통합 로그인 페이지 생성
	ssoLoginURL := c.AuthCodeURL("nuclear-tuna-plays-piano")+"&nonce=some-random-nonce"

	//2.사용자 리디렉션
	http.Redirect(w, r, ssoLoginURL, http.StatusSeeOther)
}