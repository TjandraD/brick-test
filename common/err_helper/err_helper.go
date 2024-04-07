package errhelper

import (
	"fmt"
	"money_transfer/constant"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

var errorStatusMap = map[error]int{
	constant.ErrInvalidRequest:       http.StatusBadRequest,
	constant.ErrBindRequest:          http.StatusBadRequest,
	constant.ErrInvalidTransactionId: http.StatusBadRequest,
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// returnError is a helper function to return error with default code and message.
// this will search for error code and message from error object in error map.
func ReturnError(ctx echo.Context, err error) error {
	statusCode, message := MapErrorToHttpStatus(err)
	return returnErrorWithCode(ctx, err, statusCode, message)
}

// returnErrorWithCode is a helper function to return error with custom code and message.
func returnErrorWithCode(ctx echo.Context, err error, code int, messages ...string) error {
	fmt.Printf("Error occurs in server: %s", messages)

	message := err.Error()
	if len(messages) > 0 {
		message = strings.Join(messages, " ")
	}

	return ctx.JSON(code, ErrorResponse{
		Error: message,
	})
}

func MapErrorToHttpStatus(err error) (status int, message string) {
	if err == nil {
		status = http.StatusInternalServerError
		message = constant.ErrMsgUnknownError
		return
	}

	err = unwrapAndFoundError(err)
	status, ok := errorStatusMap[err]
	if !ok {
		status = http.StatusInternalServerError
		message = constant.ErrMsgUnknownError
	} else {
		message = err.Error()
	}
	return
}

func unwrapAndFoundError(err error) error {
	switch x := err.(type) {
	case interface{ Unwrap() error }:
		err = x.Unwrap()
		if err != nil {
			// check recursively if is wrapped error
			return unwrapAndFoundError(err)
		}
	case interface{ Unwrap() []error }:
		// find first error that is in errorStatusMap
		wrappedErrors := x.Unwrap()
		for _, wrappedError := range wrappedErrors {
			_, ok := errorStatusMap[wrappedError]
			if ok {
				return wrappedError
			}
		}
	default:
		return err
	}
	return err
}
