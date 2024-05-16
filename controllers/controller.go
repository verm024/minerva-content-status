package controllers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type controller struct {
	e *echo.Echo
	db *gorm.DB
}

func Initialize(e *echo.Echo, db *gorm.DB) *controller {
	cont := controller{e, db}
	return &cont
}