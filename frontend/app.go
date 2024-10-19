package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
	"strings"

	"github.com/Ares1605/casual-chess-golang/backend/user"
	"github.com/Ares1605/casual-chess-golang/backend/oauth/googlejwt"
	"github.com/Ares1605/casual-chess-golang/backend/oauth/googleuser"
	"github.com/Ares1605/casual-chess-golang/backend/apiresps"
	"github.com/Ares1605/casual-chess-golang/frontend/kv"
	"github.com/google/uuid"
)

// App struct
type App struct {
	ctx context.Context
}


// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func awaitSignIn(customUUID uuid.UUID) (*user.User, error) {
	client := &http.Client{}
  data := url.Values{}
	req, err := http.NewRequest("GET", "http://localhost:8080/signin/await?uuid=" + customUUID.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return &user.User{}, err
	}

	req.Header.Add("Content-Type", "application/json")

	// send request
	resp, err := client.Do(req)
	if err != nil {
		return &user.User{}, err
	}
	defer resp.Body.Close()


	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &user.User{}, err
	}
	var parsed apiresps.AwaitSignIn
	err = json.Unmarshal(body, &parsed)
	if err != nil {
		return &user.User{}, err
	}

	// store the jwt in a kv store
	db, err := kv.GetDB()
	defer db.Close()
	if err == nil {
		kv.Put(db, kv.JWT, []byte(parsed.Data.Token))
	}
	return &parsed.Data.User, nil
}
func openURL(url string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
  case "windows":
    cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
  case "darwin":
    cmd = exec.Command("open", url)
  default: // "linux", "freebsd", "openbsd", "netbsd"
    // Check if running in WSL
    out, _ := exec.Command("uname", "-r").Output()
    if strings.Contains(strings.ToLower(string(out)), "microsoft") {
      // Running in WSL, use powershell.exe to open URL
      cmd = exec.Command("powershell.exe", "Start", url)
    } else {
      cmd = exec.Command("xdg-open", url)
    }
  }
  return cmd.Start()
}
func (a *App) GetSession() string {
	db, err := kv.GetDB()
	if err != nil {
		panic(err)
	}
	jwt, err := kv.Get(db, kv.JWT)
	if err != nil {
		panic(err)
	}
	googlejwt.New(string(jwt[:]))
	return string(jwt[:])
}
func (a *App) SignIn() (*user.User, error) {
	customUUID := uuid.New()
	url := "http://localhost:8080/signin?uuid=" + customUUID.String()
	if err := openURL(url); err != nil {
		return nil, err
	}

  user, err := awaitSignIn(customUUID)
  if err != nil {
  	return nil, err
  }
  return user, nil
}
func (a *App) GetFriends(fullUser user.User, googleID string) apiresps.Friends {
	client := &http.Client{}
  data := url.Values{}
	req, err := http.NewRequest("GET", "http://localhost:8080/friends", strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer " + fullUser.EncodedJWT)

	// send request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()


	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var parsed apiresps.Friends
	if err := json.Unmarshal(body, &parsed); err != nil {
		panic(err)
	}

	// store the jwt in a kv store
	return parsed
}
func (a *App) ServerOnline() bool {
	resp, err := http.Get("http://localhost:8080/ping")
  if err != nil {
      return false
  }
  defer resp.Body.Close()
  return true
}
func (a *App) GetUser(googleUser *googleuser.GoogleUser) *user.User {
	client := &http.Client{}
  data := url.Values{}
	req, err := http.NewRequest("GET", "http://localhost:8080/user", strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer " + googleUser.EncodedJWT)

	// send request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()


	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var parsed apiresps.User
	if err := json.Unmarshal(body, &parsed); err != nil {
		panic(err)
	}

	// store the jwt in a kv store
	return &parsed.Data
}
func (a *App) ValidateUsername(fullUser *user.User, username string) (*apiresps.ValidateUsernameData, error) {
	fmt.Println(username)
	client := &http.Client{}
  data := url.Values{}
	fmt.Println("http://localhost:8080/validate/username/" + url.PathEscape(username))
	req, err := http.NewRequest("GET", "http://localhost:8080/validate/username/" + url.PathEscape(username), strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer " + fullUser.EncodedJWT)

	// send request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()


	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var parsed apiresps.ValidateUsername
	if err := json.Unmarshal(body, &parsed); err != nil {
		return nil, err
	}

	// store the jwt in a kv store
	return &parsed.Data, nil
}
