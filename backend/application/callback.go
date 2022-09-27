package application

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
)

type authorizationcode struct {
	AuthorizationCode       string `json:"authorizationcode"`
}

func callbackHandler (rw http.ResponseWriter, req *http.Request) {
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

	//1. 프론트에서 보내준 authorization code 추출
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	var authcode authorizationcode
	json.Unmarshal(data, &authcode)

	//2. 코드를 가지고 토큰을 받아온다
	var opts []oauth2.AuthCodeOption
	token, err := c.Exchange(context.Background(), authcode.AuthorizationCode, opts...)
	if err != nil {
		rw.Write([]byte(fmt.Sprintf(`<p>I tried to exchange the authorize code for an access token but it did not work but got error: %s</p>`, err.Error())))
		return
	}

	fmt.Println(token.AccessToken,token.RefreshToken,token)

	// //해당 토큰을 다시 IDP로 전송하여 사용자 정보를 받아온다.
	// userinfo, err := http.NewRequest(http.MethodGet,"http://localhost:8080/user",nil)

	// if err != nil {
	// 	fmt.Fprint(rw, "사용자 정보조회 실패")
	// }

	// res, err := http.DefaultClient.Do(userinfo)

	// //받아온 사용자 정보로 특정 서비스에서만 사용 가능한 토큰으로 재발급해서 클라이언트에게 쿠키로 전송한다.
	// fmt.Println(res,"응답")
}



