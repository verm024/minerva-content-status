package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type router struct {
	e  *echo.Echo
	db *gorm.DB
}

func Initialize(e *echo.Echo, db *gorm.DB) {
	r := router{e, db}
	r.initializePreMiddleware()
	r.initializeUserRoute()
}
