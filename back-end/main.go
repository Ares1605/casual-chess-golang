package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Ares1605/casual-chess-backend/env"
)
func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		routeTo := "https://accounts.google.com/o/oauth2/v2/auth?client_id=" + env.Get("OAUTH_CLIENT_ID") + "&redirect_uri=localhost:8080/redirect&response_type=code&scope=https://www.googleapis.com/auth/userinfo.email openid"
		c.HTML(
			200,
			"<html><body><button onclick=\"x()\">click</button><script>function x() {window.location.href=\"" + routeTo + "\"}</script></body></html>",
			nil
		)
	})
	router.Run() // listen and serve on 0.0.0.0:8080
}
