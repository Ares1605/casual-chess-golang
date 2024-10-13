package apiresps

import (
  "github.com/Ares1605/casual-chess-frontend/user"
)

type AwaitSignIn struct {
  Success bool   `json:"success"`
  Token   string `json:"token"`
  User    user.User   `json:"user"`
  FirstTimeUser bool `json:"first_time_user"`
}
