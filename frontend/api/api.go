package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func Get[T any](endpoint string, data *url.Values, token *string) (*T, error) {
	client := &http.Client{}
	if data == nil {
  	data = &url.Values{}
	}
	apiUrl := "http://localhost:8080/" + endpoint
	if data != nil && len(*data) > 0 {
		apiUrl += "?" + data.Encode()
	}

	req, err := http.NewRequest("GET", apiUrl, strings.NewReader(data.Encode()))
	if err != nil {
	  return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
  if token != nil { // add bearer token if required
    req.Header.Add("Authorization", "Bearer " + *token)
  }

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
	var parsed T
	err = json.Unmarshal(body, &parsed)
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}
