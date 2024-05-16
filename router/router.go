package router

import (
	"minerva-content-status/controllers"
	"minerva-content-status/repository"
	usecase "minerva-content-status/use-case"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type router struct {
	e    *echo.Echo
	db   *gorm.DB
	cont *controllers.Controller
	uc   *usecase.UseCase
	repo *repository.Repository
}

func Initialize(e *echo.Echo, db *gorm.DB) {
	repo := repository.Initialize(db)
	uc := usecase.Initialize(repo)
	cont := controllers.Initialize(e, uc)

	r := router{e, db, cont, uc, repo}
	r.initializePreMiddleware()
	r.initializeUserRoute()

}
