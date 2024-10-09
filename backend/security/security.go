package security

import (
  "github.com/golang-jwt/jwt"
  "encoding/json"
	"github.com/gin-gonic/gin"
  "errors"
  "time"
)

type Security struct {}

func New() *Security {
  return &Security{}
}

func (*Security) Authenticate(c *gin.Context) error {
  authHeaders := c.Request.Header["Authorization"]
  if len(authHeaders) == 0 {
    return errors.New("Authorizarion header missing")
  }
  authHeader := authHeaders[0]

  prefix := "Bearer "
  if len(authHeader) < len(prefix) {
    return errors.New("Defective authorization header in request")
  }
  token := authHeader[len(prefix):]
  decoded := decodeJWT(token)
  now := time.Now().Unix()
	if now > decoded.Exp {
		return errors.New("Token has expired")
	}
}

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

func Authenticate(token string) *googleJWT {
}
func decodeJWT(token string) *googleJWT {
	parsed, _, err := new(jwt.Parser).ParseUnverified(token, jwt.MapClaims{})
	if err != nil {
	  panic(err)
	}

	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
	  panic(err)
	}

	jsonClaims, err := json.Marshal(claims)
	if err != nil {
	  panic(err)
	}

	var googleJWT googleJWT
	if err := json.Unmarshal(jsonClaims, &googleJWT); err != nil {
	  panic(err)
	}

	return &googleJWT
}
