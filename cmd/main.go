package main

import (
	"time"

	"github.com/phamtrung99/gowebbasic/client/mysql"
	"github.com/phamtrung99/gowebbasic/config"
	"github.com/phamtrung99/gowebbasic/migration"
)

func main() {
	cfg := config.GetConfig()

	// setup locale
	{
		loc, _ := time.LoadLocation(cfg.TimeZone)
		time.Local = loc
	}

	mysql.Init()
	migration.Up()

	repo := repository.New(mysql.GetClient)

}

// func main() {
// 	echoRoot := echo.New()
// 	echoRoot.Static("/public/avatar/", "public/avatar")
// 	userRoute := routes.NewUserRoute(*echoRoot)
// 	userRoute.InitUserRoute()

// 	movieRoute := routes.NewMovieRoute(*echoRoot)
// 	movieRoute.InitMovieRoute()

// 	cmtRoute := routes.NewCmtRoute(*echoRoot)
// 	cmtRoute.InitCmtRoute()

// 	err := echoRoot.Start(":" + config.GetServerConfig().Port)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
