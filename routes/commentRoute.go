package routes

import (
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"trungpham/gowebbasic/controllers"
	"trungpham/gowebbasic/middlewares"
)

type CmtRoute struct {
	CmtEcho *echo.Echo
	CmtCon  *controllers.CmtControl
}

func NewCmtRoute(echoRoot echo.Echo) *CmtRoute {
	return &CmtRoute{
		CmtEcho: &echoRoot,
		CmtCon:  controllers.NewCmtControl(),
	}
}

func (route *CmtRoute) InitCmtRoute() {
	cmtRouGr := route.CmtEcho.Group("/comment")
	{
		cmtRouGr.Use(middleware.JWTWithConfig(middlewares.Config))
		cmtRouGr.POST("", route.CmtCon.AddComment)
		cmtRouGr.DELETE("", route.CmtCon.DeleteComment)
	}
	route.CmtEcho.GET("/comment", route.CmtCon.GetComment)

}
