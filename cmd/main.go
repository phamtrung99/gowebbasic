package main

import (
	"log"
	"github.com/phamtrung99/gowebbasic/config"
)

func main() {
	cfg := config.GetConfig()
	log.Fatal(cfg.MySQL)
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
