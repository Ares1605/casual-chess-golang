package apiresps

import (
  "github.com/Ares1605/casual-chess-golang/backend/models"
  "github.com/Ares1605/casual-chess-golang/backend/user"
)

type ValidateReasonEnum string
const (
  ReasonAlreadyExists ValidateReasonEnum = "This username already exists!"
  ReasonTooLong       ValidateReasonEnum = "Username must be less than 16 characters!"
  ReasonTooShort      ValidateReasonEnum = "Username must be more than 3 characters!"
)

type ValidateUsernameData struct {
  Valid bool `json:"valid"`
  Reason ValidateReasonEnum `json:"reason,omitempty"`
}
type ValidateUsername struct {
  Resp[ValidateUsernameData]
}
type SetupUserData struct {
  Username string `json:"username"`
}
type SetupUser struct {
  Resp[SetupUserData]
}
type User struct {
  Resp[user.User]
}
type AwaitSignIn struct {
  Resp[struct {
    Token   string `json:"token"`
    User    user.User `json:"user"`
  }]
}
type Friends struct {
  Resp[[]models.BasicUser]
}
type Ping struct {
  Message string `json:"message"`
}
