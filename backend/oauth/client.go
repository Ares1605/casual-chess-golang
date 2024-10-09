package oauth

import (
	"github.com/Ares1605/casual-chess-backend/env"
)

func getClientID() string {
	return env.Get("OAUTH_CLIENT_ID")
}
func getClientSecret() string {
	return env.Get("OAUTH_CLIENT_SECRET")
}
