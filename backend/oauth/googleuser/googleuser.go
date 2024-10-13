package googleuser

import (
	"github.com/Ares1605/casual-chess-backend/oauth/googlejwt"
)

type GoogleUser struct {
  UUID string
  Email string
  Name string
  Profile string
  DecodedJWT *googlejwt.GoogleJWT
  EncodedJWT string
}

func New(token string) (*GoogleUser, error) {
  decodedJWT, err := googlejwt.New(token)
  if err != nil {
    return nil, err
  }
  return &GoogleUser {
    UUID: decodedJWT.Sub,
    Email: decodedJWT.Email,
    Profile: decodedJWT.Picture,
    Name: decodedJWT.Name,
    DecodedJWT: decodedJWT,
    EncodedJWT: token,
  }, nil
}

