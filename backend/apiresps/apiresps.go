package apiresps

import (
  "github.com/Ares1605/casual-chess-golang/backend/models"
  "github.com/Ares1605/casual-chess-golang/backend/user"
)

type ReasonEnum string
const (
  ReasonAlreadyExists ReasonEnum = "This username already exists!"
  ReasonTooLong       ReasonEnum = "Username must be less than 16 characters!"
  ReasonTooShort      ReasonEnum = "Username must be more than 3 characters!"
)
type ValidateUsernameData struct {
  Valid bool `json:"valid"` 
  Reason ReasonEnum `json:"reason,omitempty"`
}
type ValidateUsername struct {
  Success bool `json:"success"`
  Data ValidateUsernameData
  Message string `json:"message"`
}
type CreateUsernameData struct {
  Username string `json:"username"`
}
type CreateUsername struct {
  Success bool `json:"success"`
  Data CreateUsernameData
  Message string `json:"message"`
}
type User struct {
  Success bool   `json:"success"`
  Data user.User `json:"data"`
  Message string `json:"message"`
}
type AwaitSignInData struct {
  Token   string `json:"token"`
  User    user.User `json:"user"`
}
type AwaitSignIn struct {
  Success bool   `json:"success"`
  Data AwaitSignInData `json:"data"`
  Message string `json:"message"`
}
type Friends struct {
  Success bool `json:"success"`
  Data []models.User `json:"data"`
}

