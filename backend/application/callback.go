package application

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/kataras/jwt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type authorizationcode struct {
	AuthorizationCode       string `json:"authorizationcode"`
}

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

	codeVerifier := resetPKCE(rw)

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	var authcode authorizationcode
	json.Unmarshal(data, &authcode)

	var opts []oauth2.AuthCodeOption
	if isPKCE(req) {
		fmt.Println(codeVerifier)
		opts = append(opts, oauth2.SetAuthURLParam("code_verifier", codeVerifier)) 
	}
	fmt.Println(opts,"문제의 원인. code_verifier를 못받아와..ㅠㅠ")

	token, err := c.Exchange(context.Background(), authcode.AuthorizationCode, opts...)
	if err != nil {
		fmt.Println("토큰 받는데 문제 생겼어ㅠㅠ",err.Error())
		return
	}

	//fmt.Println("Access:",token.AccessToken,"Refresh:",token.RefreshToken,"ID_Token:",token.Extra("id_token"))

	//해당 토큰을 다시 IDP로 전송하여 사용자 정보를 받아온다.
	var appClientInfo = clientcredentials.Config{
		ClientID:     "vegas",
		ClientSecret: "foobar",
		Scopes:       []string{"openid","offline"},
		TokenURL:     "http://localhost:8080/oauth2/token",
	}
	
	type requestBody struct {
		AppClientConfig clientcredentials.Config `json:"app_client_config"`
		IDToken string `json:"id_token"`
	}

	var reqBody = requestBody{
		AppClientConfig: appClientInfo,
		IDToken: fmt.Sprintf("%v", token.Extra("id_token")),
	}

	reqBodyJSON, _ := json.Marshal(reqBody)

	fmt.Println(reqBody,"***")

	resp, err := http.Post("http://localhost:8080/resource?token="+token.AccessToken, "application/json", bytes.NewBuffer(reqBodyJSON))
	if err != nil {
		fmt.Fprint(rw, "사용자객체 받기 실패")
	}

	//받아온 사용자 정보로 특정 서비스에서만 사용 가능한 토큰으로 재발급해서 클라이언트에게 쿠키로 전송한다.
	type employeeInfo struct {
		DN string `json:"dn"`
		Uid string `json:"uid"`
		Employeenumber string `json:"employeenumber"`
		Cn string `json:"cn"`
		Sn string `json:"sn"`
		Mobile string `json:"mobile"`
		// Departments []string 	`json:"departments"`
		Hospitalcode string `json:"hospitalcode"`
		// Services []string `json:"services"`
	}

	data, rerr := ioutil.ReadAll(resp.Body)
	if rerr != nil {
		fmt.Println("응답바디 읽어오기 실패!")
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	defer req.Body.Close()

	var userinfo employeeInfo
	if err := json.Unmarshal(data, &userinfo); err != nil {
		fmt.Println("에러발생", err)
	}
	fmt.Println(userinfo,"사용자정보응답")

	//jwt를 생성한다. 페이로드엔 사용자 식별정보와 간단한 부서, 소속정보를 담는다. 
	var sharedKey = []byte("sercrethatmaycontainch@r$32chars") //환경변수처리.
	serviceAccessToken, err := jwt.Sign(jwt.HS256, sharedKey, userinfo, jwt.MaxAge(15*time.Minute))
	if err != nil {
		panic(err)
	}

	fmt.Println("토큰 확인하세용",string(serviceAccessToken[:]))
	http.SetCookie(rw, &http.Cookie{
		Name:   "vegasAccessToken",
		Value:  "Bearer " + string(serviceAccessToken[:]),
		Domain: "localhost:3006",
		Path: "/",
		MaxAge: 15*60,
	},
	)
	fmt.Fprint(rw,"베가스 액세스 토큰이 발급되었습니다:)")
}



