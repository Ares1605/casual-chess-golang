package main

import (
	"context"
	"errors"
	"net/url"
	"os/exec"
	"runtime"
	"strings"

	"github.com/Ares1605/casual-chess-golang/backend/apiresps"
	"github.com/Ares1605/casual-chess-golang/backend/oauth/googlejwt"
	"github.com/Ares1605/casual-chess-golang/backend/user"
	"github.com/Ares1605/casual-chess-golang/frontend/api"
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

func awaitSignIn(customUUID uuid.UUID) (*apiresps.AwaitSignIn, error) {
	endpoint := "signin/await?uuid=" + customUUID.String()
	response, err := api.Get[apiresps.AwaitSignIn](endpoint, nil, nil)
	if err != nil {
		return nil, err
	}
	if !response.Success { // return before storing it in the kv bucket
		return response, nil
	}
	// store the jwt in a kv store
	db, err := kv.GetDB()
	defer db.Close()
	if err == nil {
		kv.Put(db, kv.JWT, []byte(response.Data.Token))
	}
	return response, nil
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
func (a *App) GetOldSession() (*user.User, error) {
	db, err := kv.GetDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	jwt, err := kv.Get(db, kv.JWT)
	if err != nil {
		return nil, err
	}
	stringified := string(jwt)
	resp, err := api.Get[apiresps.User]("user", nil, &stringified)
	if err != nil {
		return nil, err
	}
	if !resp.Success {
		return nil, errors.New(resp.Error.Message)
	}
	return &resp.Data, nil
}
func (a *App) SignIn() (*apiresps.AwaitSignIn, error) {
	customUUID := uuid.New()
	toOpenURL := "http://localhost:8080/signin?uuid=" + customUUID.String()
	if err := openURL(toOpenURL); err != nil {
		return nil, err
	}

  resp, err := awaitSignIn(customUUID)

  if err != nil {
  	return nil, err
  }
  return resp, nil
}
func (a *App) GetFriends(fullUser user.User) (*apiresps.Friends, error) {
	response, err := api.Get[apiresps.Friends]("friends", nil, &fullUser.EncodedJWT)
	if err != nil {
		return nil, err
	}
	return response, err
}
func (a *App) ServerOnline() bool {
	_, err := api.Get[apiresps.Ping]("ping", nil, nil)
	return err == nil
}
func (a *App) CreateUsername(fullUser *user.User, username string) (*apiresps.SetupUser, error) {
	endpoint := "setup/user"
	data := url.Values{}
	data.Set("username", username)
	response, err := api.Get[apiresps.SetupUser](endpoint, &data, &fullUser.EncodedJWT)
	if err != nil {
		return nil, err
	}
	// store the jwt in a kv store
	return response, nil
}
func (a *App) ValidateUsername(fullUser *user.User, username string) (*apiresps.ValidateUsername, error) {

	endpoint := "validate/username/" + url.PathEscape(username)
	response, err := api.Get[apiresps.ValidateUsername](endpoint, nil, &fullUser.EncodedJWT)
	if err != nil {
		return nil, err
	}
	// store the jwt in a kv store
	return response, nil
}
func (a *App) ForceAddRespModel() *apiresps.JunkResp {
	return &apiresps.JunkResp{}
}
