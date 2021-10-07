package usecase

import (
	"github.com/phamtrung99/gowebbasic/repository"
	"github.com/phamtrung99/gowebbasic/usecase/auth"
	"github.com/phamtrung99/gowebbasic/usecase/comment"
	"github.com/phamtrung99/gowebbasic/usecase/healthcheck"
	"github.com/phamtrung99/gowebbasic/usecase/movie"
	"github.com/phamtrung99/gowebbasic/usecase/user"
	"github.com/phamtrung99/gowebbasic/usecase/userfavorite"
)

type UseCase struct {
	HealthCheck healthcheck.IUsecase
	User        user.IUsecase
	Auth        auth.IUsecase
	UserFavor   userfavorite.IUsecase
	Movie       movie.IUsecase
	Comment     comment.IUsecase
}

func New(repo *repository.Repository) *UseCase {
	return &UseCase{
		HealthCheck: healthcheck.New(repo),
		User:        user.New(repo),
		Auth:        auth.New(repo),
		UserFavor:   userfavorite.New(repo),
		Movie:       movie.New(repo),
		Comment:     comment.New(repo),
	}
}
