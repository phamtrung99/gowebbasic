package comment

import (
	"context"
	"fmt"

	"github.com/phamtrung99/gowebbasic/model"
	"github.com/phamtrung99/gowebbasic/util/myerror"
)

type ListCommentRequest struct {
	Filter    *model.CommentFilter
	Paginator *model.Paginator
	OrderBy   string `json:"order_by,omitempty" query:"order_by"`
	OrderType string `json:"order_type,omitempty" query:"order_type"`
}

func (u *Usecase) GetList(ctx context.Context, req *ListCommentRequest) (*model.CommentResult, error) {

	//condition list
	listMovieID := make([]interface{}, 1)
	listMovieID[0] = req.Filter.MovieID

	listParentID := make([]interface{}, 1)
	listParentID[0] = 1 //Default 1 is root comment, not subcomment

	//Check if get list sub comment
	if req.Filter.ParentID != 0 {
		listParentID[0] = req.Filter.ParentID
	}

	conditions := []model.Condition{
		{Pattern: "movie_id",
			Values: listMovieID},
		{Pattern: "parent_id",
			Values: listParentID},
	}

	//Order
	orders := make([]string, 0)
	if req.OrderBy != "" {
		orders = []string{fmt.Sprintf("%s %s", req.OrderBy, req.OrderType)}
	}

	//Paging
	paginator := &model.Paginator{
		Page:  1,
		Limit: 20,
	}

	if req.Paginator != nil {
		paginator = req.Paginator
	}

	result, err := u.cmtRepo.Find(ctx, conditions, paginator, orders)

	if err != nil {
		return nil, myerror.ErrFindComment(err)
	}

	return result, nil
}
