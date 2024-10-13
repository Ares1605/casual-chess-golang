package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"log"

	pkgerrors "github.com/pkg/errors"
	"github.com/Ares1605/casual-chess-backend/db"
	"github.com/Ares1605/casual-chess-backend/env"
	"github.com/Ares1605/casual-chess-backend/models"
	"github.com/Ares1605/casual-chess-backend/oauth"
	"github.com/Ares1605/casual-chess-backend/security"
	"github.com/Ares1605/casual-chess-backend/oauth/googleuser"
	"github.com/Ares1605/casual-chess-backend/security/securityerror"
	"github.com/gin-gonic/gin"
)

type cacheResponse struct {
	success bool
	user *models.User
	token string
	firstTimeUser bool
}

func getGoogleUser(c *gin.Context) (*googleuser.GoogleUser, error) {
  value, exists := c.Get("googleuser")
  if !exists {
    return nil, errors.New("user not found in context")
  }

  user, ok := value.(*googleuser.GoogleUser)
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
		user, err := getGoogleUser(c)
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
				token: "",
				user: &models.User{},
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
			"user": response.user,
			"token": response.token,
			"first_time_user": response.firstTimeUser,
		})
	})
	router.LoadHTMLGlob("templates/*")
	router.GET("/signin", func(c *gin.Context) {
		uuid := c.Query("uuid")
		if uuid == "" {
			c.HTML(200, "error.html", gin.H{})
			return
		}

		routeTo := "https://accounts.google.com/o/oauth2/v2/auth?client_id=" + env.Get("OAUTH_CLIENT_ID") + "&redirect_uri=http://localhost:8080/redirect&response_type=code&scope=https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/userinfo.profile openid&state=" + uuid
		c.HTML(200, "signin.html", gin.H{
			"routeTo": routeTo,
			"uuid": uuid,
		})
	})
	router.GET("/redirect", func(c *gin.Context) {
		fmt.Println("Redirect called...")

		code := c.Query("code")
		uuid := c.Query("state")
		googleUser, err := oauth.GetGoogleUser(code)
		if err != nil {
			c.HTML(200, "error.html", gin.H{})
			wrappedErr := pkgerrors.WithStack(err)
    	log.Printf("Error occurred: %+v", wrappedErr)
			return
		}

		fmt.Println("  - looking for uuid: ", uuid)
		done, ok := cache[uuid]
		if ok {
			dbConn, err := db.Conn()
			if err != nil {
				c.HTML(200, "error.html", gin.H{})
				wrappedErr := pkgerrors.WithStack(err)
    		log.Printf("Error occurred: %+v", wrappedErr)
				return
			}
			dbUser, err := models.GetUserFromUUID(dbConn, googleUser.UUID)
			firstTimeUser := false
			if err != nil {
				firstTimeUser = true
				// db user doesn't exist, create one
				dbUser, err = models.CreateUser(dbConn, googleUser)
				if err != nil {
					c.HTML(200, "error.html", gin.H{})
					wrappedErr := pkgerrors.WithStack(err)
    			log.Printf("Error occurred: %+v", wrappedErr)
					return
				}
			}
			response := &cacheResponse{
				success: true,
				user: dbUser,
				token: googleUser.EncodedJWT,
				firstTimeUser: firstTimeUser,
			}
			done <- response
			c.HTML(200, "redirect.html", gin.H{})
		} else {
			c.HTML(200, "error.html", gin.H{})
			wrappedErr := pkgerrors.WithStack(errors.New("Never found the channel"))
    	log.Printf("Error occurred: %+v", wrappedErr)
		}
	})
	router.GET("/game/:id", func(c *gin.Context) {
		idStr := c.Param("id")

		id, err := strconv.Atoi(idStr)
    if err != nil {
    	securityMnger.Reject(c, "game id must be an integer", securityerror.Validation)
    }

		dbConn, err := db.Conn()
		if err != nil {
			securityMnger.Reject(c, err.Error(), securityerror.Internal)			
		}
		game, err := models.GetGame(dbConn, id)
		if err != nil {
			securityMnger.Reject(c, err.Error(), securityerror.Internal)
		}
		c.JSON(http.StatusOK, game)
	})
	router.GET("/user/uuid/:uuid", func(c *gin.Context) {
		uuid := c.Param("uuid")

		dbConn, err := db.Conn()
		if err != nil {
			securityMnger.Reject(c, err.Error(), securityerror.Internal)			
		}
		dbUser, err := models.GetUserFromUUID(dbConn, uuid)
		if err != nil {
			securityMnger.Reject(c, err.Error(), securityerror.Internal)
		}
		c.JSON(http.StatusOK, dbUser)
	})
	router.GET("/user/id/:id", func(c *gin.Context) {
		idStr := c.Param("id")

		id, err := strconv.Atoi(idStr)
    if err != nil {
    	securityMnger.Reject(c, "game id must be an integer", securityerror.Validation)
    }

		dbConn, err := db.Conn()
		if err != nil {
			securityMnger.Reject(c, err.Error(), securityerror.Internal)			
		}
		dbUser, err := models.GetUserFromID(dbConn, id)
		if err != nil {
			securityMnger.Reject(c, err.Error(), securityerror.Internal)
		}
		c.JSON(http.StatusOK, dbUser)
	})
	router.Run() // listen and serve on 0.0.0.0:8080
}
