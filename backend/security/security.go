package security

import (
  "net/http"
  "strconv"
	"github.com/gin-gonic/gin"
  "time"
  "github.com/Ares1605/casual-chess-backend/env"
  "github.com/Ares1605/casual-chess-backend/security/securityerror"
  "github.com/Ares1605/casual-chess-backend/oauth/googleuser"
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
  reqstedUser, err := googleuser.New(token)
  if err != nil {
    security.Reject(c, "Provided token could not be parsed", securityerror.Authentication)
    return
  }
  now := time.Now().Unix()
	extendedExpiry := reqstedUser.DecodedJWT.Exp + int64(getTokenExpiryExtension())
	if now > extendedExpiry {
		security.Reject(c, "Provided token has expired", securityerror.Authentication)
		return
	}

  c.Set("googleuser", reqstedUser)
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
func (*Security) Accept(c *gin.Context, data any, message string) {
  response := gin.H{
    "success": true,
    "data": data,
  }
  if message == "" {
    response["message"] = nil
  } else {
    response["message"] = message
  }
  c.AbortWithStatusJSON(http.StatusOK, response) 
}
