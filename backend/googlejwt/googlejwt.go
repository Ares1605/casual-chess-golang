package googlejwt

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
func decodeJWT() {
  
}

func New(token string) {
}
