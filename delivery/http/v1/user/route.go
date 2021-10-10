package user

import (
	"github.com/labstack/echo/v4"
	"github.com/phamtrung99/gowebbasic/usecase"
	"github.com/phamtrung99/gowebbasic/usecase/user"
)

type Route struct {
	userUseCase user.IUsecase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{
		userUseCase: useCase.User,
	}

	group.PUT("",r.Update)
	group.POST("/changepassword",r.ChangePassword)
}
