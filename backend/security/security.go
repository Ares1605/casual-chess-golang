package security

import (
  "net/http"
  "strconv"
  "github.com/golang-jwt/jwt"
  "encoding/json"
	"github.com/gin-gonic/gin"
  "errors"
  "time"
  "github.com/Ares1605/casual-chess-backend/env"
  "fmt"
)

type Security struct {}

func New() *Security {
  return &Security{}
}

func getTokenExpiryExtension() uint64 {
  extension := env.Get("OAUTH_TOKEN_EXPIRY_EXTENSION")
  parsed, err := strconv.ParseUint(extension, 10, 64)
  if err != nil {
    fmt.Println("OAUTH_TOKEN_EXPIRY_EXTENSION invalid format (expects int), default to 0")
    parsed = 0
  }
  return parsed * 60 * 60
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
	extendedExpiry := decoded.Exp + int64(getTokenExpiryExtension())
	if now > extendedExpiry {
		return errors.New("Token has expired")
	}
}
func (*Security) Reject(c *gin.Context) {
  c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
    "error": "Authentication failed",
  }) 
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
