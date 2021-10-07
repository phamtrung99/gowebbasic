package userfavorite

import (
	"github.com/phamtrung99/gowebbasic/repository"
	"github.com/phamtrung99/gowebbasic/repository/movie"
	"github.com/phamtrung99/gowebbasic/repository/userfavorite"
)

type Usecase struct {
	userFavorRepo userfavorite.Repository
	movieRepo     movie.Repository
}

// New .
func New(repo *repository.Repository) IUsecase {
	return &Usecase{
		userFavorRepo: repo.UserFavor,
		movieRepo: repo.Movie,
	}
}
