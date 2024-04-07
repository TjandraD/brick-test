package constant

import "errors"

const (
	ErrMsgUnknownError = "Unknown Error"
)

var (
	ErrInvalidRequest = errors.New("invalid request")
	ErrBindRequest    = errors.New("error binding request")
)
