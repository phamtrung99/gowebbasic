package myerror

import (
	"net/http"

	"github.com/phamtrung99/gopkg/apperror"
)

// ErrAuthLogin .
func ErrGetMovie(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 30000010,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "fail to get movie",
		Message:   "fail to get movie",
	}
}


