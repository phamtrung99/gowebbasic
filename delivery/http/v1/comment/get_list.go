package comment

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/phamtrung99/gopkg/apperror"
	"github.com/phamtrung99/gopkg/utils"
	"github.com/phamtrung99/gowebbasic/model"
	"github.com/phamtrung99/gowebbasic/usecase/comment"
)

func (r *Route) GetList(c echo.Context) error {
	var (
		ctx       = &utils.CustomEchoContext{Context: c}
		appError  = apperror.AppError{}
		paginator = model.Paginator{}
	)

	if err := c.Bind(&paginator); err != nil {
		_ = errors.As(err, &appError)

		return utils.Response.Error(ctx, apperror.ErrInvalidInput(err))
	}

	req := comment.ListCommentRequest{
		Paginator: &paginator,
	}
	// Bind order by
	if err := c.Bind(&req); err != nil {
		_ = errors.As(err, &appError)

		return utils.Response.Error(ctx, apperror.ErrInvalidInput(err))
	}

	res, err := r.commentUseCase.GetList(ctx, &req)

	if err != nil {
		_ = errors.As(err, &appError)

		return utils.Response.Error(ctx, appError)
	}

	return utils.Response.Success(ctx, res)
}
