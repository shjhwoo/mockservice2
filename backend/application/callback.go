package application

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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
	// 	opts = append(opts, oauth2.SetAuthURLParam("code_verifier", codeVerifier)) 
	// }
	//fmt.Println(opts,"문제의 원인. code_verifier를 못받아와..ㅠㅠ")

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
		TokenURL:     "http://localhost:3846/oauth2/token",
	}
	
	type requestBody struct {
		AppClientConfig clientcredentials.Config `json:"app_client_config"`
		IDToken string `json:"id_token"`
	}

	var reqBody = requestBody{
		AppClientConfig: appClientInfo,
		IDToken: fmt.Sprintf("%v", token.Extra("id_token")),
	}

	reqBodyJSON, err := json.Marshal(reqBody)

	fmt.Println(reqBody,"***")

	resp, err := http.Post("http://localhost:8080/resource?token="+token.AccessToken, "application/json", bytes.NewBuffer(reqBodyJSON))


	if err != nil {
		fmt.Fprint(rw, "사용자객체 받기 실패")
	}

	//받아온 사용자 정보로 특정 서비스에서만 사용 가능한 토큰으로 재발급해서 클라이언트에게 쿠키로 전송한다.
	fmt.Println(resp,"응답")

	//apply jwt...

	// http.SetCookie(rw, &http.Cookie{
	// 	Name:   "vegas",
	// 	Value:  "Bearer " + token.AccessToken,
	// 	Domain: "localhost:3006",
	// 	Path: "/",
	// },
	// )
}



