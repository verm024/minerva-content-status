package router

import (
	"minerva-content-status/controllers"
	"minerva-content-status/repository"
	usecase "minerva-content-status/use-case"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type contentManagementArcRouter struct {
	e  *echo.Echo
	db *gorm.DB
}

func (r *contentManagementArcRouter) initialize() {
	repo := repository.InitializeContentManagementArcRepository(r.db)
	cmrepo := repository.InitializeContentManagementRepository(r.db)
	uc := usecase.InitializeContentManagementArcUseCase(repo, cmrepo, r.db)
	cont := controllers.InitializeContentManagementArcController(uc)

	cm := initializeCustomMiddleware(r.e)

	contentManagementRouter := r.e.Group("/content-management")
	contentManagementRouter.Use(cm.appMiddleware)
	contentManagementRouter.POST("/:content_management_id/content-management-arc", cont.CreateCMArc, cm.roleBasedRouteMiddleware([]string{"SA"}))
	contentManagementRouter.GET("/:content_management_id/content-management-arc", cont.CMArcListByCMId, cm.roleBasedRouteMiddleware([]string{"SA"}))

	contentManagementArcRouter := r.e.Group("/content-management-arc")
	contentManagementArcRouter.Use(cm.appMiddleware)
	contentManagementArcRouter.PUT("/:content_management_arc_id", cont.UpdateCMArc, cm.roleBasedRouteMiddleware([]string{"SA"}))
	contentManagementArcRouter.DELETE("/:content_management_arc_id", cont.DeleteCMArc, cm.roleBasedRouteMiddleware([]string{"SA"}))
}
