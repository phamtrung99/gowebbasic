package myerror

import (
	"net/http"

	"github.com/phamtrung99/gopkg/apperror"
)

// ErrUserList .
func ErrUserList(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 20000010,
		HTTPCode:  http.StatusInternalServerError,
		Info:      "fail to get user list",
		Message:   "fail to get user list.",
	}
}

// ErrUserGet .
func ErrUserGet(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 20000020,
		HTTPCode:  http.StatusInternalServerError,
		Info:      "fail to get user",
		Message:   "fail to get user.",
	}
}

// ErrUserUpdate .
func ErrUserUpdate(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 200000030,
		HTTPCode:  http.StatusInternalServerError,
		Info:      "fail to update user",
		Message:   "fail to update user.",
	}
}

// ErrUserInvalidUserID .
func ErrUserInvalidUserID(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 200000040,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "invalid user id",
		Message:   "invalid user id.",
	}
}

// ErrUserCreate .
func ErrUserCreate(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 200000050,
		HTTPCode:  http.StatusInternalServerError,
		Info:      "fail to create user",
		Message:   "fail to create user.",
	}
}

// ErrUserCreate .
func ErrPwdMatching(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 200000060,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "retype password not matched",
		Message:   "retype password not matched.",
	}
}
