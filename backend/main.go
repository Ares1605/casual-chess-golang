package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"database/sql"

	"github.com/Ares1605/casual-chess-golang/backend/db"
	"github.com/Ares1605/casual-chess-golang/backend/apiresps"
	"github.com/Ares1605/casual-chess-golang/backend/env"
	"github.com/Ares1605/casual-chess-golang/backend/models"
	"github.com/Ares1605/casual-chess-golang/backend/user"
	"github.com/Ares1605/casual-chess-golang/backend/oauth"
	"github.com/Ares1605/casual-chess-golang/backend/security"
	"github.com/Ares1605/casual-chess-golang/backend/oauth/googleuser"
	"github.com/Ares1605/casual-chess-golang/backend/security/securityerror"
)

type cacheResponse struct {
	success bool
	user *user.User
	token string
}

func validateUsername(dbConn *sql.DB, username string) (bool, apiresps.ValidateReasonEnum, error) {
	if len(username) < 4 {
		return false, apiresps.ReasonTooShort, nil
	}
	if len(username) >= 16 {
		return false, apiresps.ReasonTooLong, nil
	}

	dbConn, err := db.Conn()
	if err != nil {
		return false, apiresps.ReasonTooLong, err
	}
	exists, err := models.UsernameExists(dbConn, username)
	if err != nil {
		return false, apiresps.ReasonTooLong, err
	}
	if exists {
		return false, apiresps.ReasonAlreadyExists, nil
	}
	return true, apiresps.ReasonAlreadyExists, nil
}
func getUser(dbConn *sql.DB, googleUser *googleuser.GoogleUser, createIfDBNone bool) (*user.User, error) {
	dbUser, err := models.GetUser(dbConn, googleUser.ID)
	if err == sql.ErrNoRows { // create user should only run if this is a did not fetch any results error...
		if createIfDBNone {
			dbUser, err = models.CreateUser(dbConn, googleUser)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else if err != nil {
		return nil, err
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

	router := gin.Default()
	router.GET("/ping/auth", security.Authenticate, func(c *gin.Context) {
		user, err := getGoogleUser(c)
    if err != nil {
      security.RejectError(c, err)
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
		security.Accept(c, gin.H{
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
			security.RejectHTML(c, err)
			return
		}

		fmt.Println("  - looking for uuid: ", uuid)
		done, ok := cache[uuid]
		if ok {
			dbConn, err := db.Conn()
			if err != nil {
				security.RejectHTML(c, err)
				return
			}
			user, err := getUser(dbConn, googleUser, true)
			if err != nil {
				security.RejectHTML(c, err)
			}
			response := &cacheResponse{
				success: true,
				user: user,
				token: googleUser.EncodedJWT,
			}
			done <- response
			c.HTML(200, "redirect.html", gin.H{})
		} else {
			err = errors.New("Never found the channel")
			security.RejectHTML(c, err)
		}
	})
	router.GET("/game/:id", func(c *gin.Context) {
		idStr := c.Param("id")

		id, err := strconv.Atoi(idStr)
    if err != nil {
    	security.Reject(c, "game id must be an integer", securityerror.Validation)
    }

		dbConn, err := db.Conn()
		if err != nil {
			security.RejectError(c, err)			
		}
		game, err := models.GetGame(dbConn, id)
		if err != nil {
			security.RejectError(c, err)
		}
		security.Accept(c, game, "")
	})
	router.GET("/user/:googleID", func(c *gin.Context) {
		googleID := c.Param("googleID")

		dbConn, err := db.Conn()
		if err != nil {
			security.RejectError(c, err)			
		}
		dbUser, err := models.GetUser(dbConn, googleID)
		if err != nil {
			security.RejectError(c, err)
		}
		c.JSON(http.StatusOK, dbUser)
	})
	router.GET("/user/id/:id", func(c *gin.Context) {
		idStr := c.Param("id")

		id, err := strconv.Atoi(idStr)
    if err != nil {
    	security.Reject(c, "user id must be an integer", securityerror.Validation)
    }

		dbConn, err := db.Conn()
		if err != nil {
			security.RejectError(c, err)			
		}
		dbUser, err := models.GetUserFromID(dbConn, id)
		if err != nil {
			security.RejectError(c, err)
		}
		security.Accept(c, dbUser, "")
	})
	router.GET("/friend/request/deny/user/:googleID", security.Authenticate, func(c *gin.Context) {
		friendGoogleID := c.Param("googleID")

		user, err := getGoogleUser(c)
		if err != nil {
			security.RejectError(c, err)
		}
		dbConn, err := db.Conn()
		if err != nil {
			security.RejectError(c, err)
		}
		pendingRow, err := models.GetPendingRow(dbConn, friendGoogleID, user.ID)
		if err != nil {
			security.Reject(c, err.Error(), securityerror.Custom)
		}
		if models.DeletePendingFriendRequest(dbConn, pendingRow.ID); err != nil {
			security.RejectError(c, err)
		}
		security.Accept(c, nil, "")
	})
	router.GET("/friend/request/accept/user/:googleID", security.Authenticate, func(c *gin.Context) {
		friendGoogleID := c.Param("googleID")

		user, err := getGoogleUser(c)
		if err != nil {
			security.RejectError(c, err)
		}
		dbConn, err := db.Conn()
		pendingRow, err := models.GetPendingRow(dbConn, friendGoogleID, user.ID)
		if err != nil {
			security.Reject(c, err.Error(), securityerror.Custom)
		}
		if models.DeletePendingFriendRequest(dbConn, pendingRow.ID); err != nil {
			security.RejectError(c, err)
		}
		err = models.AddFriend(dbConn, friendGoogleID, user.ID)
		if err != nil {
			security.RejectError(c, err)
		}
		security.Accept(c, nil, "")
	})
	router.GET("/friend/request/send/user/:googleID", security.Authenticate, func(c *gin.Context) {
		friendGoogleID := c.Param("googleID")

		googleUser, err := getGoogleUser(c)
    if err != nil {
      security.RejectError(c, err)
      return
    }
		dbConn, err := db.Conn()
		var count uint8
		err = dbConn.QueryRow("SELECT COUNT(*) FROM friends WHERE invitee_google_id=? AND invited_google_id=? OR invited_google_id=? AND invitee_google_id=? LIMIT 1", googleUser.ID, friendGoogleID, googleUser.ID, friendGoogleID).Scan(
			&count,
			)
		if err != nil {
			security.Reject(c, err.Error(), securityerror.Custom)
			return
		}
		if count != 0 {
			security.Reject(c, "You're already friends!", securityerror.Custom)
			return
		}

		err = dbConn.QueryRow("SELECT COUNT(*) FROM friends WHERE invited_google_id=? AND invitee_google_id=?", googleUser.ID, friendGoogleID).Scan(
			&count,
			)
		if err != nil {
			security.Reject(c, err.Error(), securityerror.Custom)
			return
		}
		if count != 0 {
			security.Reject(c, "You already have a pending friend request!", securityerror.Custom)
			return
		}

		_, err = dbConn.Exec("INSERT INTO pending_friends (invited_google_id, invitee_google_id) VALUES (?, ?)", googleUser.ID, friendGoogleID)
  	if err != nil {
  		security.RejectError(c, err)
  	}
  	security.Accept(c, nil, "")
	})
	router.GET("/friends", security.Authenticate, func(c *gin.Context) {
		googleUser, err := getGoogleUser(c)
    if err != nil {
      security.RejectError(c, err)
      return
    }
		dbConn, err := db.Conn()
		if err != nil {
			security.RejectError(c, err)
		}
		friends, err := models.GetFriends(dbConn, googleUser.ID)
		if err != nil {
			security.RejectError(c, err)
		}

		security.Accept(c, friends, "")
	})
	router.GET("/user/:googleID/friends", func(c *gin.Context) {
		googleID := c.Param("googleID")

		dbConn, err := db.Conn()
		if err != nil {
			security.RejectError(c, err)
		}
		friends, err := models.GetFriends(dbConn, googleID)
		if err != nil {
			security.RejectError(c, err)
		}
		security.Accept(c, friends, "")
	})
	router.GET("/user", security.Authenticate, func(c *gin.Context) {
		googleUser, err := getGoogleUser(c)
		if err != nil {
			security.RejectError(c, err)
		}
		
		dbConn, err := db.Conn()
		if err != nil {
			security.RejectError(c, err)
		}

    user, err := getUser(dbConn, googleUser, false)
    if err != nil {
      security.RejectError(c, err)
      return
    }
		security.Accept(c, user, "")
	})
	router.GET("/setup/user", security.Authenticate, func (c *gin.Context) {
		googleUser, err := getGoogleUser(c)

		username := c.Query("username")

		dbConn, err := db.Conn()
		if err != nil {
			security.RejectError(c, err)
			return
		}
		dbUser, err := models.GetUser(dbConn, googleUser.ID)
		if err != nil {
			security.RejectError(c, err)
			return
		}
		if dbUser.SetupComplete {
			security.Reject(c, "You cannot setup an already setup account!", securityerror.Custom)
			return
		}
		valid, reason, err := validateUsername(dbConn, username)
		if err != nil {
			security.RejectError(c, err)
			return
		}
		if !valid {
			security.Reject(c, string(reason), securityerror.Custom)
			return
		}

		// if username is valid, attempt to create it
		err = models.SetupUser(dbConn, googleUser, username)
		if err != nil {
			security.RejectError(c, err)
			return
		}
		security.Accept(c, &apiresps.SetupUserData{
			Username: username,
		}, "")
	})
	router.GET("/create/username/:username", security.Authenticate, func (c *gin.Context) {
		googleUser, err := getGoogleUser(c)

		username := c.Param("username")

		dbConn, err := db.Conn()
		if err != nil {
			security.RejectError(c, err)
			return
		}
		valid, reason, err := validateUsername(dbConn, username)
		if err != nil {
			security.RejectError(c, err)
			return
		}
		if !valid {
			security.Reject(c, string(reason), securityerror.Custom)
			return
		}

		// if username is valid, attempt to create it
		err = models.SetupUser(dbConn, googleUser, username)
		if err != nil {
			security.RejectError(c, err)
			return
		}
		security.Accept(c, &apiresps.SetupUserData{
			Username: username,
		}, "")
	})
	router.GET("/validate/username/:username", security.Authenticate, func(c *gin.Context) {
		username := c.Param("username")

		dbConn, err := db.Conn()
		if err != nil {
			security.RejectError(c, err)
			return
		}
		valid, reason, err := validateUsername(dbConn, username)
		if err != nil {
			security.RejectError(c, err)
			return
		}
		if valid {
			security.Accept(c, &apiresps.ValidateUsernameData{
				Valid: true,
			}, "")
		} else {
			security.Accept(c, &apiresps.ValidateUsernameData{
				Valid: false,
				Reason: reason,
			}, "")
		}
	})
	router.Run()
}
