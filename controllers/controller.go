package controllers

import (
	usecase "minerva-content-status/use-case"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	e  *echo.Echo
	uc *usecase.UseCase
}

func Initialize(e *echo.Echo, uc *usecase.UseCase) *Controller {
	cont := Controller{e, uc}
	return &cont
}
