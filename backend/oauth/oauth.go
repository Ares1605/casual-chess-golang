package oauth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"github.com/Ares1605/casual-chess-backend/oauth/user"
)

type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	IDToken      string `json:"id_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func GetUser(code string) *user.User {
  data := url.Values{}
	data.Set("code", code)
	data.Set("client_id", getClientID())
	data.Set("client_secret", getClientSecret())
	data.Set("redirect_uri", "http://localhost:8080/redirect")
	data.Set("grant_type", "authorization_code")

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://oauth2.googleapis.com/token", strings.NewReader(data.Encode()))
	if err != nil {
	  panic(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// send request
	resp, err := client.Do(req)
	if err != nil {
	  panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	  panic(err)
	}
	var tokenResponse tokenResponse
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
	  panic(err)
	}

  return user.New(tokenResponse.IDToken)
}
