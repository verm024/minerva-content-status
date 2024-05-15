package router

import (
	"github.com/labstack/echo/v4"
)

type router struct {
	e *echo.Echo
}

func Initialize(e *echo.Echo)  {
	r := router{e}
	r.initializeUserRoute()
}
