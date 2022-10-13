package application

import (
	"net/http"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type Claims struct {
	DN string `json:"dn"`
	Uid string `json:"uid"`
	Employeenumber string `json:"employeenumber"`
	Cn string `json:"cn"`
	Sn string `json:"sn"`
	Mobile string `json:"mobile"`
	Departments []string 	`json:"departments"`
	Hospitalcode string `json:"hospitalcode"`
	Services []string `json:"services"`
	jwt.StandardClaims
}

func checksso(c *gin.Context) {
	//sso 쿠키가 살아있는지 확인해야지...
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
	c.JSON(http.StatusOK, gin.H{
		"message": "SSO 쿠키를 확인합니다",
		"redirectionURL": ssoLoginURL,
	})
}

	//아래는 배포 환경에서 쓸 코드임.

	// cookie, err := req.Cookie("hanchartAccessToken")
	// if err != nil {
	// 	fmt.Println("서비스 쿠키가 없어",err)
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"message": "no service cookie",
	// 	})
	// 	return
	// 	//sso 쿠키가 있는지 자동으로 확인을 해서 재발급 시도를 하도록 하자
	// 	//sso 쿠키 확인하는 곳으로 사용자 이동시키기
	// 	//withCredentials헤더 설정. 
	// 	//resp, err := http.Get("http://localhost:8080/api/oauth2/auth")
	// 	//sso 쿠키마저도 없음
	// 	//sso 쿠키 살아있음

	// }
	
	// tknStr := strings.Split(hanchartCookie.Cookie,"=")[1]
	// // tknStr := cookie.Value
	// fmt.Println(tknStr)

	// claims := &Claims{}

	// tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
	// 	return jwtKey, nil
	// })

	// if err != nil {
	// 	if err == jwt.ErrSignatureInvalid {
	// 		rw.WriteHeader(http.StatusUnauthorized)
	// 		fmt.Println("권한 없는 이상한 토큰")
	// 		return
	// 	}
	// 	rw.WriteHeader(http.StatusBadRequest)
	// 	fmt.Println("아무튼 잘못된 요청")
	// 	return
	// }

	// if !tkn.Valid {
	// 	rw.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }

	// //유효한 토큰이다. 즉 이미 로그인이 되어 있다.
	// //사용자를 원래 서비스로 되돌려주자. 
	// c.JSON(http.StatusOK, gin.H{
	// 	"message":"has login cookie",
	// })
