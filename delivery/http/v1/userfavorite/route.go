package userfavorite

import (
	"github.com/labstack/echo/v4"
	"github.com/phamtrung99/gowebbasic/usecase"
	"github.com/phamtrung99/gowebbasic/usecase/userfavorite"
)

type Route struct {
	userFavorUseCase userfavorite.IUsecase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{
		userFavorUseCase: useCase.UserFavor,
	}

	group.GET("", r.GetFavorite)
	group.POST("", r.AddFavorite)
}
