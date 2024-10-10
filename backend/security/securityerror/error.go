package securityerror

type ErrorType int

const (
    Authentication ErrorType = iota
    Authorization
    Validation
    Internal
)

func (e ErrorType) String() string {
    return [...]string{"AUTHENTICATION_ERROR", "AUTHORIZATION_ERROR", "VALIDATION_ERROR", "INTERNAL_ERROR"}[e]
}
