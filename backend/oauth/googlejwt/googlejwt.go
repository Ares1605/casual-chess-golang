package googlejwt

import (
  "encoding/json"
  "github.com/golang-jwt/jwt"
  "errors"
)

type GoogleJWT struct {
	Iss           string `json:"iss"`
	Azp           string `json:"azp"`
	Aud           string `json:"aud"`
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	AtHash        string `json:"at_hash"`
	Iat           int64  `json:"iat"`
	Exp           int64  `json:"exp"`
}

func New(idToken string) (*GoogleJWT, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(idToken, jwt.MapClaims{})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Claims does not exist in token")
	}

	jsonClaims, err := json.Marshal(claims)
	if err != nil {
		return nil, err
	}

	var parsedJWT GoogleJWT
	if err := json.Unmarshal(jsonClaims, &parsedJWT); err != nil {
	  panic(err)
	}

	return &parsedJWT, nil
}

