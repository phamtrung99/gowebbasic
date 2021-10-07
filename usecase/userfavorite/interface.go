package userfavorite

import (
	"context"

	"github.com/phamtrung99/gowebbasic/model"
)

// IUsecase .
type IUsecase interface {
	AddMovieToFavorite(ctx context.Context, req AddFavoriteRequest) (*model.UserFavorite, error)
	GetFavoriteMovie(ctx context.Context, req UserFavorRequest) (*model.MovieResult, error)
}
