package main

import (
	"fmt"
	"log"
	"net/http"
	"errors"
	"github.com/Ares1605/casual-chess-backend/env"
	"github.com/Ares1605/casual-chess-backend/oauth"
	"github.com/Ares1605/casual-chess-backend/oauth/user"
	"github.com/Ares1605/casual-chess-backend/security"
	"github.com/gin-gonic/gin"
  "github.com/Ares1605/casual-chess-backend/security/securityerror"
)

type cacheResponse struct {
	success bool
	user *user.User 
}

func getUser(c *gin.Context) (*user.User, error) {
  value, exists := c.Get("user")
  if !exists {
    return nil, errors.New("user not found in context")
  }

  user, ok := value.(*user.User)
  if !ok {
    return nil, errors.New("invalid user type in context")
  }

  return user, nil
}

func main() {
	cache := make(map[string]chan *cacheResponse)
	securityMnger := security.New()

	router := gin.Default()
	router.GET("/ping/auth", securityMnger.Authenticate, func(c *gin.Context) {
		user, err := getUser(c)
    if err != nil {
      securityMnger.Reject(c, err.Error(), securityerror.Internal)
      return
    }

		c.String(http.StatusOK, "Hi" + user.Email)
	})
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	router.GET("/signin/await", func(c *gin.Context) {
		uuid := c.Query("uuid")
		oldDone, ok := cache[uuid]
		if ok {
			response := &cacheResponse{
				success: false,
				user: &user.User{}, // faulty user
			}
			oldDone <- response
		}

		done := make(chan *cacheResponse)
		// create caching item 
		fmt.Println("  - attaching uuid: ", uuid)
		cache[uuid] = done

		// wait for channel to finish
		response := <- done
		c.JSON(200, gin.H{
			"success": response.success,
			"email": response.user.Email,
			"uuid": response.user.UUID,
			"token": response.user.Token,
		})
	})
	router.LoadHTMLGlob("templates/*")
	router.GET("/signin", func(c *gin.Context) {
		uuid := c.Query("uuid")
		if uuid == "" {
			c.HTML(200, "error.html", gin.H{})
			return
		}

		routeTo := "https://accounts.google.com/o/oauth2/v2/auth?client_id=" + env.Get("OAUTH_CLIENT_ID") + "&redirect_uri=http://localhost:8080/redirect&response_type=code&scope=https://www.googleapis.com/auth/userinfo.email openid&state=" + uuid
		c.HTML(200, "signin.html", gin.H{
			"routeTo": routeTo,
			"uuid": uuid,
		})
	})
	router.GET("/redirect", func(c *gin.Context) {
		fmt.Println("Redirect called...")
		c.HTML(200, "redirect.html", gin.H{})

		code := c.Query("code")
		uuid := c.Query("uuid")
		user := oauth.GetUser(code)

		fmt.Println("  - looking for uuid: ", uuid)
		done, ok := cache[uuid]
		if ok {
			response := &cacheResponse{
				success: true,
				user: user,
			}
			done <- response
		} else {
			log.Fatal("errmm, never found channel")
		}
	})
	router.Run() // listen and serve on 0.0.0.0:8080
}
