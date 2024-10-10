package main

import (
	"context"
	"fmt"
	"os/exec"
  "runtime"
  "strings"
	"net/http"
	"io/ioutil"
	"net/url"
	"log"
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
func awaitSignIn() string {
	client := &http.Client{}
  data := url.Values{}
	req, err := http.NewRequest("GET", "http://localhost:8080/signin/await", strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")

	fmt.Println("\n\n\n    Sneding!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n\n\n")
	// send request
	resp, err := client.Do(req)
	if err != nil {
	  log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	  panic(err)
	}
  return "response:" + string(body)
}
func (a *App) SignIn() string {
	customUUID := uuid.New()
	url := "http://localhost:8080/signin?uuid=" + customUUID.String()

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
  err := cmd.Start()
  if err != nil {
  	panic(err)
  }

  return awaitSignIn()
}
