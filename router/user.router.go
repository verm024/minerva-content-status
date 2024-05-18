package router

import (
	"minerva-content-status/controllers"
	"minerva-content-status/repository"
	usecase "minerva-content-status/use-case"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type userRouter struct {
	e  *echo.Echo
	db *gorm.DB
}

func (r *userRouter) initialize() {
	repo := repository.InitializeUserRepository(r.db)
	uc := usecase.InitializeUserUseCase(repo)
	cont := controllers.InitializeUserController(uc)

	userRouter := r.e.Group("/user")
	userRouter.GET("", cont.GetAllUsers)
	userRouter.POST("", cont.RegisterNewUser)
	userRouter.POST("/login", cont.Login)
}
