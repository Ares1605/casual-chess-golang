package user

import (
	"github.com/Ares1605/casual-chess-golang/backend/models"
	"github.com/Ares1605/casual-chess-golang/backend/oauth/googlejwt"
	"github.com/Ares1605/casual-chess-golang/backend/oauth/googleuser"
	"errors"
)

type User struct {
  ID int64 `json:"id"`
  Username string `json:"username"`
  SetupComplete bool `json:"setup_complete"`
  GoogleID string `json:"google_id"`
  ProfileURL string `json:"profile_url"`
  Email string `json:"email"`
  Name string `json:"name"`
  DecodedJWT *googlejwt.GoogleJWT `json:"decoded_jwt"`
  EncodedJWT string `json:"encoded_jwt"`
}

func MergeUsers (googleUser *googleuser.GoogleUser, dbUser *models.User) (*User, error) {
  if googleUser.ID != dbUser.GoogleID {
    return nil, errors.New("google user must have the same id as the db user google id")
  }
  return &User{
    ID: dbUser.ID,
    Username: dbUser.Username,
    SetupComplete: dbUser.SetupComplete,
    GoogleID: googleUser.ID,
    ProfileURL: googleUser.ProfileURL,
    Email: googleUser.Email,
    Name: googleUser.Name,
    DecodedJWT: googleUser.DecodedJWT,
    EncodedJWT: googleUser.EncodedJWT,
  }, nil
}
