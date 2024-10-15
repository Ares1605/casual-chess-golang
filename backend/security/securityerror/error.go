package securityerror

type ErrorType int

const (
    Authentication ErrorType = iota
    Authorization
    Validation
    Internal
    Custom
)

func (e ErrorType) String() string {
    return [...]string{"AUTHENTICATION_ERROR", "AUTHORIZATION_ERROR", "VALIDATION_ERROR", "INTERNAL_ERROR", ""}[e]
}

// convert Custom -> null
func (e ErrorType) MarshalJSON() ([]byte, error) {
    if e == Custom {
        return []byte("null"), nil
    }
    return []byte(`"` + e.String() + `"`), nil
}
