package constant

import "errors"

var (
	ErrCommon = errors.New("something went wrong")
	ErrLoginNotFound = errors.New("email or password is incorrect")
	ErrInvalidJWT = errors.New("invalid access")
	ErrInvalidJWTSubject = errors.New("invalid subject")
	ErrUserAlreadyExist = errors.New("user already exist")
	ErrTagAlreadyExist = errors.New("tag already exist")
)

type LoginErrorType struct {
	Msg string
}

type JwtErrorType struct {
	Msg string
}

func (c LoginErrorType) Error() string {
	return c.Msg
}

func (c JwtErrorType) Error() string {
	return c.Msg
}