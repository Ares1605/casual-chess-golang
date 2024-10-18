package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"log"

	pkgerrors "github.com/pkg/errors"
	"github.com/Ares1605/casual-chess-golang/backend/db"
	"github.com/Ares1605/casual-chess-golang/backend/apiresps"
	"github.com/Ares1605/casual-chess-golang/backend/env"
	"github.com/Ares1605/casual-chess-golang/backend/models"
	"github.com/Ares1605/casual-chess-golang/backend/user"
	"github.com/Ares1605/casual-chess-golang/backend/oauth"
	"github.com/Ares1605/casual-chess-golang/backend/security"
	"github.com/Ares1605/casual-chess-golang/backend/oauth/googleuser"
	"github.com/Ares1605/casual-chess-golang/backend/security/securityerror"
	"github.com/gin-gonic/gin"
	"database/sql"
)

type cacheResponse struct {
	success bool
	user *user.User
	token string
}

func getUser(dbConn *sql.DB, googleUser *googleuser.GoogleUser, createIfDBNone bool) (*user.User, error) {
	dbUser, err := models.GetUser(dbConn, googleUser.ID)
	if err != nil { // create user should only run if this is a did not fetch any results error...
		if createIfDBNone {
			dbUser, err = models.CreateUser(dbConn, googleUser)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	fullUser, err := user.MergeUsers(googleUser, dbUser)
	if err != nil {
		return nil, err
	}
	return fullUser, nil
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
				user: &user.User{},
			}
			oldDone <- response
		}

		done := make(chan *cacheResponse)
		// create caching item 
		fmt.Println("  - attaching uuid: ", uuid)
		cache[uuid] = done

		// wait for channel to finish
		response := <- done
		securityMnger.Accept(c, gin.H{
			"user": response.user,
			"token": response.token,
		}, "Sign in complete!")
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
			user, err := getUser(dbConn, googleUser, true)
			if err != nil {
				c.HTML(200, "error.html", gin.H{})
				wrappedErr := pkgerrors.WithStack(err)
    		log.Printf("Error occurred: %+v", wrappedErr)
			}
			response := &cacheResponse{
				success: true,
				user: user,
				token: googleUser.EncodedJWT,
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
		securityMnger.Accept(c, game, "")
	})
	router.GET("/user/:googleID", func(c *gin.Context) {
		googleID := c.Param("googleID")

		dbConn, err := db.Conn()
		if err != nil {
			securityMnger.Reject(c, err.Error(), securityerror.Internal)			
		}
		dbUser, err := models.GetUser(dbConn, googleID)
		if err != nil {
			securityMnger.Reject(c, err.Error(), securityerror.Internal)
		}
		c.JSON(http.StatusOK, dbUser)
	})
	router.GET("/user/id/:id", func(c *gin.Context) {
		idStr := c.Param("id")

		id, err := strconv.Atoi(idStr)
    if err != nil {
    	securityMnger.Reject(c, "user id must be an integer", securityerror.Validation)
    }

		dbConn, err := db.Conn()
		if err != nil {
			securityMnger.Reject(c, err.Error(), securityerror.Internal)			
		}
		dbUser, err := models.GetUserFromID(dbConn, id)
		if err != nil {
			securityMnger.Reject(c, err.Error(), securityerror.Internal)
		}
		securityMnger.Accept(c, dbUser, "")
	})
	router.GET("/friend/request/deny/user/:googleID", securityMnger.Authenticate, func(c *gin.Context) {
		friendGoogleID := c.Param("googleID")

		user, err := getGoogleUser(c)
		if err != nil {
			securityMnger.Reject(c, err.Error(), securityerror.Internal)
		}
		dbConn, err := db.Conn()
		if err != nil {
			securityMnger.Reject(c, err.Error(), securityerror.Internal)
		}
		pendingRow, err := models.GetPendingRow(dbConn, friendGoogleID, user.ID)
		if err != nil {
			securityMnger.Reject(c, err.Error(), securityerror.Custom)
		}
		if models.DeletePendingFriendRequest(dbConn, pendingRow.ID); err != nil {
			securityMnger.Reject(c, err.Error(), securityerror.Internal)
		}
		securityMnger.Accept(c, nil, "")
	})
	router.GET("/friend/request/accept/user/:googleID", securityMnger.Authenticate, func(c *gin.Context) {
		friendGoogleID := c.Param("googleID")

		user, err := getGoogleUser(c)
		if err != nil {
			securityMnger.Reject(c, err.Error(), securityerror.Internal)
		}
		dbConn, err := db.Conn()
		pendingRow, err := models.GetPendingRow(dbConn, friendGoogleID, user.ID)
		if err != nil {
			securityMnger.Reject(c, err.Error(), securityerror.Custom)
		}
		if models.DeletePendingFriendRequest(dbConn, pendingRow.ID); err != nil {
			securityMnger.Reject(c, err.Error(), securityerror.Internal)
		}
		err = models.AddFriend(dbConn, friendGoogleID, user.ID)
		if err != nil {
			securityMnger.Reject(c, err.Error(), securityerror.Internal)
		}
		securityMnger.Accept(c, nil, "")
	})
	router.GET("/friend/request/send/user/:googleID", securityMnger.Authenticate, func(c *gin.Context) {
		friendGoogleID := c.Param("googleID")

		user, err := getGoogleUser(c)
    if err != nil {
      securityMnger.Reject(c, err.Error(), securityerror.Internal)
      return
    }
		dbConn, err := db.Conn()
		var count uint8
		err = dbConn.QueryRow("SELECT COUNT(*) FROM friends WHERE invitee_google_id=? AND invited_google_id=? OR invited_google_id=? AND invitee_google_id=? LIMIT 1", user.ID, friendGoogleID, user.ID, friendGoogleID).Scan(
			&count,
			)
		if err != nil {
			securityMnger.Reject(c, err.Error(), securityerror.Custom)
			return
		}
		if count != 0 {
			securityMnger.Reject(c, "You're already friends!", securityerror.Custom)
			return
		}

		err = dbConn.QueryRow("SELECT COUNT(*) FROM friends WHERE invited_google_id=? AND invitee_google_id=?", user.ID, friendGoogleID).Scan(
			&count,
			)
		if err != nil {
			securityMnger.Reject(c, err.Error(), securityerror.Custom)
			return
		}
		if count != 0 {
			securityMnger.Reject(c, "You already have a pending friend request!", securityerror.Custom)
			return
		}

		_, err = dbConn.Exec("INSERT INTO pending_friends (invited_google_id, invitee_google_id) VALUES (?, ?)", user.ID, friendGoogleID)
  	if err != nil {
  		securityMnger.Reject(c, err.Error(), securityerror.Internal)
  	}
  	securityMnger.Accept(c, nil, "")
	})
	router.GET("/friends", securityMnger.Authenticate, func(c *gin.Context) {
		googleUser, err := getGoogleUser(c)
    if err != nil {
      securityMnger.Reject(c, err.Error(), securityerror.Internal)
      return
    }
		dbConn, err := db.Conn()
		if err != nil {
			securityMnger.Reject(c, err.Error(), securityerror.Internal)
		}
		friends, err := models.GetFriends(dbConn, googleUser.ID)
		if err != nil {
			securityMnger.Reject(c, err.Error(), securityerror.Internal)
		}

		securityMnger.Accept(c, friends, "")
	})
	router.GET("/user/:googleID/friends", func(c *gin.Context) {
		googleID := c.Param("googleID")

		dbConn, err := db.Conn()
		if err != nil {
			securityMnger.Reject(c, err.Error(), securityerror.Internal)
		}
		friends, err := models.GetFriends(dbConn, googleID)
		if err != nil {
			securityMnger.Reject(c, err.Error(), securityerror.Internal)
		}
		securityMnger.Accept(c, friends, "")
	})
	router.GET("/user", securityMnger.Authenticate, func(c *gin.Context) {
		googleUser, err := getGoogleUser(c)
		
		dbConn, err := db.Conn()
		if err != nil {
			securityMnger.Reject(c, err.Error(), securityerror.Internal)
		}

    user, err := getUser(dbConn, googleUser, false)
    if err != nil {
      securityMnger.Reject(c, err.Error(), securityerror.Internal)
      return
    }
		securityMnger.Accept(c, user, "")
	})
	router.GET("/validate/username/:username", securityMnger.Authenticate, func(c *gin.Context) {

		username := c.Param("username")

		if len(username) < 4 {
			securityMnger.Accept(c, &apiresps.ValidateUsernameData {
				Valid: false,
				Reason: apiresps.ReasonTooShort,
			}, "")
		}
		if len(username) >= 16 {
			securityMnger.Accept(c, &apiresps.ValidateUsernameData {
				Valid: false,
				Reason: apiresps.ReasonTooLong,
			}, "")
		}

		dbConn, err := db.Conn()
		if err != nil {
			securityMnger.Reject(c, err.Error(), securityerror.Internal)
		}
		exists, err := models.UsernameExists(dbConn, username)
		if err != nil {
			securityMnger.Reject(c, err.Error(), securityerror.Internal)
		}
		if exists {
			securityMnger.Accept(c, &apiresps.ValidateUsernameData {
				Valid: false,
				Reason: apiresps.ReasonAlreadyExists,
			}, "")
		}
		securityMnger.Accept(c, &apiresps.ValidateUsernameData {
			Valid: true,
		}, "")
	})
	router.Run()
}
