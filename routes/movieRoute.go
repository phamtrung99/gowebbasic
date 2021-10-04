package routes

import (
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"trungpham/gowebbasic/controllers"
	"trungpham/gowebbasic/middlewares"
)

type MovieRoute struct {
	MovieEcho *echo.Echo
	MovieCon  *controllers.MovieControl
}

func NewMovieRoute(echoRoot echo.Echo) *MovieRoute {
	return &MovieRoute{
		MovieEcho: &echoRoot,
		MovieCon:  controllers.NewMovieControl(),
	}
}

func (route *MovieRoute) InitMovieRoute() {
	movieRouGr := route.MovieEcho.Group("/movies")
	{
		movieRouGr.GET("", route.MovieCon.SearchMovie)

	}

	adminMovieRouGr := route.MovieEcho.Group("admin/movies")
	{
		adminMovieRouGr.Use(middleware.JWTWithConfig(middlewares.Config))
		adminMovieRouGr.Use(middlewares.CheckAdmin)
		adminMovieRouGr.POST("", route.MovieCon.InsertMovie)
		adminMovieRouGr.DELETE("", route.MovieCon.DeleteMovie)
		adminMovieRouGr.PUT("", route.MovieCon.UpdateMovie)

	}

}
