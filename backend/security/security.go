package security

import (
  "net/http"
  "strconv"
	"github.com/gin-gonic/gin"
  "time"
	pkgerrors "github.com/pkg/errors"
  "github.com/Ares1605/casual-chess-golang/backend/env"
  "github.com/Ares1605/casual-chess-golang/backend/security/securityerror"
  "github.com/Ares1605/casual-chess-golang/backend/oauth/googleuser"
  "fmt"
)

func getTokenExpiryExtension() uint64 {
  extension := env.Get("OAUTH_TOKEN_EXPIRY_EXTENSION")
  parsed, err := strconv.ParseUint(extension, 10, 64)
  if err != nil {
    fmt.Println("OAUTH_TOKEN_EXPIRY_EXTENSION invalid format (expects int), default to 0")
    parsed = 0
  }
  return parsed * 60 * 60
}

func Authenticate(c *gin.Context) {
  authHeaders := c.Request.Header["Authorization"]
  if len(authHeaders) == 0 {
		Reject(c, "Authorization header is missing", securityerror.Authentication)
    return
  }
  authHeader := authHeaders[0]

  prefix := "Bearer "
  if len(authHeader) < len(prefix) {
		Reject(c, "Defect authorization header in request", securityerror.Authentication)
    return
  }
  token := authHeader[len(prefix):]
  reqstedUser, err := googleuser.New(token)
  if err != nil {
    Reject(c, "Provided token could not be parsed", securityerror.Authentication)
    return
  }
  now := time.Now().Unix()
	extendedExpiry := reqstedUser.DecodedJWT.Exp + int64(getTokenExpiryExtension())
	if now > extendedExpiry {
		Reject(c, "Provided token has expired", securityerror.Authentication)
		return
	}

  c.Set("googleuser", reqstedUser)
  c.Next()
}
func RejectHTML(c *gin.Context, err error) {
  stackTrace := fmt.Sprintf("%+v", pkgerrors.WithStack(err))
  c.HTML(200, "error.html", gin.H{
    "stack_trace": stackTrace,
    "error": err.Error(),
  })
}
func RejectError(c *gin.Context, err error) {
  Reject(c, err.Error(), securityerror.Internal)
}
func Reject(c *gin.Context, message string, errorType securityerror.ErrorType) {
  c.AbortWithStatusJSON(http.StatusOK, gin.H{
    "success": false,
    "error": gin.H{
      "type": errorType.String(),
      "message": message,
    },
  }) 
}
func Accept(c *gin.Context, data any, message string) {
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
