package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/unrolled/secure"
)

func (r *router) initializePreMiddleware() {
	r.e.Pre(middleware.RemoveTrailingSlash())

	secureMiddleware := secure.New(secure.Options{FrameDeny: true, ContentTypeNosniff: true, BrowserXssFilter: true, STSIncludeSubdomains: true, STSSeconds: 31536000, STSPreload: true})
	r.e.Use(echo.WrapMiddleware(secureMiddleware.Handler))

	r.e.Use(middleware.Logger())
	r.e.Use(middleware.Recover())
}
