package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func sloHandler(c *gin.Context) {

	con := oauth2.Config{
		ClientID:     "hanchart",
		ClientSecret: "foobar",
		RedirectURL:  "http://localhost:4006/callback",
		Scopes:       []string{"openid", "offline"},
		Endpoint: oauth2.Endpoint{
			TokenURL: "http://localhost:8080/api/oauth2/token",
			AuthURL:  "http://localhost:8080/",
		},
	}

	pkceCodeVerifier := generateCodeVerifier(64)
	pkceCodeChallenge = generateCodeChallenge(pkceCodeVerifier)

	ssoLogoutURL := con.AuthCodeURL("nuclear-tuna-plays-piano") + "&nonce=some-random-nonce&code_challenge=" + pkceCodeChallenge + "&code_challenge_method=S256&logout=true"

	c.SetCookie("hanchartRefreshToken", "", 0, "/", "localhost", false, false)

	c.JSON(http.StatusOK, gin.H{
		"redirectionURL": ssoLogoutURL,
	})
}
