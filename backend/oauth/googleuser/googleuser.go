package googleuser

import (
	"github.com/Ares1605/casual-chess-golang/backend/oauth/googlejwt"
)

type GoogleUser struct {
  ID string `json:"id"`
  Email string `json:"email"`
  Name string `json:"name"`
  ProfileURL string `json:"profile"`
  DecodedJWT *googlejwt.GoogleJWT `json:"decoded_jwt"`
  EncodedJWT string `json:"encoded_jwt"`
}

func New(token string) (*GoogleUser, error) {
  decodedJWT, err := googlejwt.New(token)
  if err != nil {
    return nil, err
  }
  return &GoogleUser {
    ID: decodedJWT.Sub,
    Email: decodedJWT.Email,
    ProfileURL: decodedJWT.Picture,
    Name: decodedJWT.Name,
    DecodedJWT: decodedJWT,
    EncodedJWT: token,
  }, nil
}
