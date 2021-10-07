package userfavorite

import (
	"context"
	"fmt"

	"github.com/phamtrung99/gopkg/middleware"
	"github.com/phamtrung99/gowebbasic/model"
	"github.com/phamtrung99/gowebbasic/util/myerror"
)

type UserFavorRequest struct {
	Paginator *model.Paginator
	OrderBy   string `json:"order_by,omitempty" query:"order_by"`
	OrderType string `json:"order_type,omitempty" query:"order_type"`
}

func (u *Usecase) GetFavoriteMovie(ctx context.Context, req UserFavorRequest) (*model.MovieResult, error) {

	//Get current userId from Token.
	claim := middleware.GetClaim(ctx)
	userID := claim.UserID

	listMovieID, err := u.userFavorRepo.GetListMovieIDByUserID(ctx, userID)

	if err != nil {
		return &model.MovieResult{}, err
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

	//condition list
	conditionValue := make([]interface{}, len(listMovieID))
	for i, m := range listMovieID {
		conditionValue[i] = m
	}

	conditions := []model.Condition{{Pattern: "movie_id", Values: conditionValue}}

	movieResult, err := u.movieRepo.Find(ctx, conditions, paginator, orders)

	if err != nil {
		return nil, myerror.ErrGetMovie(err)
	}

	return movieResult, nil
}
