package routes

import (
	"github.com/labstack/echo/v4/middleware"
	echo "github.com/labstack/echo/v4"

	"trungpham/gowebbasic/controllers"
	"trungpham/gowebbasic/middlewares"
)

type UserRoute struct {
	UserEcho *echo.Echo
	UserCon  *controllers.UserControl
}

func NewUserRoute(echoRoot echo.Echo) *UserRoute {
	return &UserRoute{
		UserEcho: &echoRoot,
		UserCon:  controllers.NewUserControl(),
	}
}

func (route *UserRoute) InitUserRoute() {
	authRouGr := route.UserEcho.Group("/auth")
	{
		authRouGr.POST("/login", route.UserCon.Login)
		authRouGr.POST("/register", route.UserCon.Register)
	}

	UserRouGr := route.UserEcho.Group("/users")
	{
		UserRouGr.Use(middleware.JWTWithConfig(middlewares.Config))
		UserRouGr.PUT("", route.UserCon.UpdateUser)
		UserRouGr.POST("/changepassword", route.UserCon.ChangePassword)
	}

	favorRouGr := route.UserEcho.Group("/favorites")
	{
		favorRouGr.Use(middleware.JWTWithConfig(middlewares.Config))
		favorRouGr.POST("", route.UserCon.AddMovieToFavorite)
		favorRouGr.GET("", route.UserCon.GetFavoriteMovie)
	}

}
