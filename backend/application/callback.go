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

// type basicClient struct {
// 	clientID     string
// 	clientSecret string

// 	client http.Client
// }

// func newBasicClient(clientID string, clientSecret string) *basicClient {
// 	fmt.Println("11")
// 	return &basicClient{
// 		clientID:     clientID,
// 		clientSecret: clientSecret,
// 		client: http.Client{
// 			Timeout: time.Second * 5,
// 		},
// 	}
// }


func callbackHandler (rw http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL,"토큰은 여기서 받아용")
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

	//codeVerifier := resetPKCE(rw)

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	var authcode authorizationcode
	json.Unmarshal(data, &authcode)

	var opts []oauth2.AuthCodeOption
	// if isPKCE(req) {
	// 	fmt.Println(codeVerifier)
	// 	opts = append(opts, oauth2.SetAuthURLParam("code_verifier", codeVerifier)) //URa~GLE7o5p9~MF7a_5_P1XG9slFm7eywMCavZ~t8bvsbRB3nR4mGEnpyZmmvKgp
	// }
	//fmt.Println(opts,"문제의 원인. code_verifier를 못받아와..ㅠㅠ")

	token, err := c.Exchange(context.Background(), authcode.AuthorizationCode, opts...)
	if err != nil {
		fmt.Println("토큰 받는데 문제 생겼어ㅠㅠ",err.Error())
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



