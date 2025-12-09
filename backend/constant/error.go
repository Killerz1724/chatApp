package constant

import "errors"

var (
	ErrCommon = errors.New("something went wrong")
)

type CustomErrorType struct {
	Msg string
}

func (c CustomErrorType) Error() string {
	return c.Msg
}