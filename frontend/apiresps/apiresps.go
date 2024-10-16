package apiresps

import (
  "github.com/Ares1605/casual-chess-golang/backend/models"
)

type AwaitSignIn struct {
  Success bool   `json:"success"`
  Token   string `json:"token"`
  User    any `json:"user"`
  FirstTimeUser bool `json:"first_time_user"`
}
type Friends struct {
  Success bool `json:"success"`
  Data []models.User `json:"data"`
}
