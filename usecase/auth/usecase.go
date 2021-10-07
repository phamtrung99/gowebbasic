package auth

import (
	"github.com/phamtrung99/gowebbasic/repository"
	"github.com/phamtrung99/gowebbasic/repository/user"
)

type Usecase struct {
	userRepo user.Repository
}

// New .
func New(repo *repository.Repository) IUsecase {
	return &Usecase{
		userRepo: repo.User,
	}
}
