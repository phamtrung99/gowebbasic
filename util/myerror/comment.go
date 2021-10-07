package myerror

import (
	"net/http"

	"github.com/phamtrung99/gopkg/apperror"
)

// Err comment  .
func ErrBadWordContent(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 50000010,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "content was contained bad word",
		Message:   "content was contained bad word",
	}
}

func ErrFindComment(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 50000020,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "fail to get list comment",
		Message:   "fail to get list comment",
	}
}

func ErrNotCmtActor(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 50000030,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "current user is not actor of this comment",
		Message:   "current user is not actor of this comment",
	}
}
