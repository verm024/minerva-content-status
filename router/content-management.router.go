package router

import (
	"minerva-content-status/controllers"
	"minerva-content-status/repository"
	usecase "minerva-content-status/use-case"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type contentManagementRouter struct {
	e  *echo.Echo
	db *gorm.DB
}

func (r *contentManagementRouter) initialize() {
	repo := repository.InitializeContentManagementRepository(r.db)
	uc := usecase.InitializeContentManagementUseCase(repo)
	cont := controllers.InitializeContentManagementController(uc)

	cm := initializeCustomMiddleware(r.e)

	contentManagementRouter := r.e.Group("/content-management")
	contentManagementRouter.Use(cm.appMiddleware)
	contentManagementRouter.GET("", cont.GetContentManagementDashboard, cm.roleBasedRouteMiddleware([]string{"SA"}))
	contentManagementRouter.POST("", cont.CreateContent)
	contentManagementRouter.PUT("/:content_management_id", cont.UpdateContent)
}
