package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/phamtrung99/gowebbasic/delivery/http/v1/auth"
	"github.com/phamtrung99/gowebbasic/delivery/http/v1/comment"
	"github.com/phamtrung99/gowebbasic/delivery/http/v1/movie"
	"github.com/phamtrung99/gowebbasic/delivery/http/v1/user"
	"github.com/phamtrung99/gowebbasic/delivery/http/v1/userfavorite"
	"github.com/phamtrung99/gowebbasic/repository"
	"github.com/phamtrung99/gowebbasic/usecase"
)

// NewHTTPHandler .
func NewHTTPHandler(repo *repository.Repository, ucase *usecase.UseCase) *echo.Echo {
	e := echo.New()
	// cfg := config.GetConfig()

	// skipper := func(c echo.Context) bool {
	// 	p := c.Request().URL.Path

	// 	return p == "/health_check" || strings.HasPrefix(p, "/docs")
	// }

	// loggerCfg := middleware.DefaultLoggerConfig
	// loggerCfg.Skipper = skipper

	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	Skipper:      middleware.DefaultSkipper,
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{
	// 		http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions,
	// 	},
	// }))
	// e.Use(middleware.LoggerWithConfig(loggerCfg))
	// e.Use(middleware.Recover())
	// e.Pre(middleware.RemoveTrailingSlash())
	// e.Use(sentryecho.New(sentryecho.Options{
	// 	Repanic: true,
	// }))
	// e.Use(myMiddleware.Auth(cfg.Jwt.Key, skipper, false))

	// if cfg.Endpoints.DatadogAgentEndpoint != "" {
	// 	e.Use(myMiddleware.DataDogTrace("hus-echo"))
	// }

	e.GET("/health_check", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	//e.GET("/docs/*", echoSwagger.WrapHandler)

	apiV1 := e.Group("/v1")

	user.Init(apiV1.Group("/users"), ucase)
	auth.Init(apiV1.Group("/auth"), ucase)
	userfavorite.Init(apiV1.Group("/favorites"), ucase)
	comment.Init(apiV1.Group("/comments"), ucase)
	movie.Init(apiV1.Group("/movies"), ucase)

	return e
}
