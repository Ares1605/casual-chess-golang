package googleuser

import (
	"github.com/Ares1605/casual-chess-backend/oauth/googlejwt"
)

type GoogleUser struct {
  UUID string
  Email string
  JWT *googlejwt.GoogleJWT
}

func New(token string) (*GoogleUser, error) {
  parsedJWT, err := googlejwt.New(token)
  if err != nil {
    return nil, err
  }
  return &GoogleUser {
    UUID: parsedJWT.Sub,
    Email: parsedJWT.Email,
    JWT: parsedJWT,
  }, nil
}

