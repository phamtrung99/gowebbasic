package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	echo "github.com/labstack/echo/v4"

	"trungpham/gowebbasic/config"
	"trungpham/gowebbasic/routes"
)

func main() {
	echoRoot := echo.New()
	echoRoot.Static("/public/avatar/", "public/avatar")
	userRoute := routes.NewUserRoute(*echoRoot)
	userRoute.InitUserRoute()

	movieRoute := routes.NewMovieRoute(*echoRoot)
	movieRoute.InitMovieRoute()

	cmtRoute := routes.NewCmtRoute(*echoRoot)
	cmtRoute.InitCmtRoute()

	err := echoRoot.Start(":" + config.GetServerConfig().Port)

	if err != nil {
		log.Fatal(err)
	}
}
