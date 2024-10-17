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

	"github.com/Ares1605/casual-chess-golang/frontend/apiresps"
	"github.com/Ares1605/casual-chess-golang/frontend/kv"
	"github.com/Ares1605/casual-chess-golang/backend/oauth/googlejwt"
	"github.com/Ares1605/casual-chess-golang/backend/oauth/googleuser"
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

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
func awaitSignIn(customUUID uuid.UUID) (*googleuser.GoogleUser, error) {
	client := &http.Client{}
  data := url.Values{}
	req, err := http.NewRequest("GET", "http://localhost:8080/signin/await?uuid=" + customUUID.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return &googleuser.GoogleUser{}, err
	}

	req.Header.Add("Content-Type", "application/json")

	// send request
	resp, err := client.Do(req)
	if err != nil {
		return &googleuser.GoogleUser{}, err
	}
	defer resp.Body.Close()


	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &googleuser.GoogleUser{}, err
	}
	var parsed apiresps.AwaitSignIn
	err = json.Unmarshal(body, &parsed)
	if err != nil {
		return &googleuser.GoogleUser{}, err
	}

	// store the jwt in a kv store
	db, err := kv.GetDB()
	defer db.Close()
	if err == nil {
		kv.Put(db, kv.JWT, []byte(parsed.Token))
	}
	return googleuser.New(parsed.Token)
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
func (a *App) SignIn() *googleuser.GoogleUser {
	customUUID := uuid.New()
	url := "http://localhost:8080/signin?uuid=" + customUUID.String()

	if err := openURL(url); err != nil {
		panic(err)
	}

  googleUser, err := awaitSignIn(customUUID)
  if err != nil {
  	panic(err)
  }
  return googleUser
}
func (a *App) GetFriends(googleUser *googleuser.GoogleUser) apiresps.Friends {
	client := &http.Client{}
  data := url.Values{}
	req, err := http.NewRequest("GET", "http://localhost:8080/friends", strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer " + googleUser.EncodedJWT)
	fmt.Println(googleUser.EncodedJWT)

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
	fmt.Println(body)
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
