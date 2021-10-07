package comment

import (
	"github.com/phamtrung99/gowebbasic/repository"
	"github.com/phamtrung99/gowebbasic/repository/comment"
	"github.com/phamtrung99/gowebbasic/repository/user"
)

type Usecase struct {
	cmtRepo  comment.Repository
	userRepo user.Repository
}

// New .
func New(repo *repository.Repository) IUsecase {
	return &Usecase{
		cmtRepo:  repo.Comment,
		userRepo: repo.User,
	}
}
