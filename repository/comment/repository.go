package comment

import (
	"context"

	"github.com/phamtrung99/gowebbasic/model"
)

// Repository .
type Repository interface {
	Find(
		ctx context.Context,
		conditions []model.Condition,
		paginator *model.Paginator,
		orders []string,
	) (*model.CommentResult, error)
	Insert(ctx context.Context, comment *model.Comment) (*model.Comment, error)
	Delete(ctx context.Context, id int64) error
	DeleteSubCmt(ctx context.Context, parentID int64) error
	GetById(ctx context.Context, id int64) (*model.Comment, error)
}
