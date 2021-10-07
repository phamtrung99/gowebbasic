package myerror

import (
	"net/http"

	"github.com/phamtrung99/gopkg/apperror"
)

// ErrUserFavorite.
func ErrFavorMovieExist(err error) apperror.AppError {
	return apperror.AppError{
		Raw:       err,
		ErrorCode: 200000010,
		HTTPCode:  http.StatusNotAcceptable,
		Info:      "movie is existed in user favorite",
		Message:   "movie is existed in user favorite.",
	}
}
