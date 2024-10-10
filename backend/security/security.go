package security

import (
  "net/http"
  "strconv"
  "github.com/golang-jwt/jwt"
  "encoding/json"
	"github.com/gin-gonic/gin"
  "time"
  "github.com/Ares1605/casual-chess-backend/env"
  "github.com/Ares1605/casual-chess-backend/security/securityerror"
  "github.com/Ares1605/casual-chess-backend/oauth/user"
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

func (security *Security) Authenticate(c *gin.Context) {
  authHeaders := c.Request.Header["Authorization"]
  if len(authHeaders) == 0 {
		security.Reject(c, "Authorization header is missing", securityerror.Authentication)
    return
  }
  authHeader := authHeaders[0]

  prefix := "Bearer "
  if len(authHeader) < len(prefix) {
		security.Reject(c, "Defect authorization header in request", securityerror.Authentication)
    return
  }
  token := authHeader[len(prefix):]
  reqstedUser := user.New(token)
  now := time.Now().Unix()
	extendedExpiry := reqstedUser.Expiry + int64(getTokenExpiryExtension())
	if now > extendedExpiry {
		security.Reject(c, "Provided token has expired", securityerror.Authentication)
		return
	}

  c.Set("user", reqstedUser)
  c.Next()
}
func (*Security) Reject(c *gin.Context, errorMessage string, errorType securityerror.ErrorType) {
  c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
    "success": false,
    "error": gin.H{
      "type": errorType.String(),
      "message": errorMessage,
    },
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

func decodeJWT(token string) *googleJWT {
	parsed, _, err := new(jwt.Parser).ParseUnverified(token, jwt.MapClaims{})
	if err != nil {
	  panic(err)
	}

	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
	  panic(ok)
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
