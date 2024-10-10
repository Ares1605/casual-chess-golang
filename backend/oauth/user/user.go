package user

import (
	"github.com/golang-jwt/jwt"
	"encoding/json"
)

type User struct {
  UUID string
  Email string
	Token string
	Expiry int64
}

func New(idToken string) *User {
  googleJWT := decodeJWT(idToken)
  return &User {
    UUID: googleJWT.Sub,
    Email: googleJWT.Email,
		Expiry: googleJWT.Exp,
		Token: idToken,
  }
}

type googleJWT struct {
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

func decodeJWT(idToken string) *googleJWT {
	token, _, err := new(jwt.Parser).ParseUnverified(idToken, jwt.MapClaims{})
	if err != nil {
	  panic(err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
	  panic(ok)
	}

	jsonClaims, err := json.Marshal(claims)
	if err != nil {
	  panic(err)
	}

	var googleJWT googleJWT
	if err := json.Unmarshal(jsonClaims, &googleJWT); err != nil {
	  panic(err)
	}

	return &googleJWT
}
