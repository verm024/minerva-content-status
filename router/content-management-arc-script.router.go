package router

import (
	"minerva-content-status/controllers"
	"minerva-content-status/repository"
	usecase "minerva-content-status/use-case"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type contentManagementArcScriptRouter struct {
	e  *echo.Echo
	db *gorm.DB
}

func (r *contentManagementArcScriptRouter) initialize() {
	repo := repository.InitializeContentManagementArcScriptRepository(r.db)
	cmarepo := repository.InitializeContentManagementArcRepository(r.db)
	uc := usecase.InitializeContentManagementArcScriptUseCase(repo, cmarepo, r.db)
	cont := controllers.InitializeContentManagementArcScriptController(uc)

	cm := initializeCustomMiddleware(r.e)

	contentManagementArcRouter := r.e.Group("/content-management-arc")
	contentManagementArcRouter.Use(cm.appMiddleware)
	contentManagementArcRouter.POST("/:content_management_arc_id/script", cont.CreateCMAScript, cm.roleBasedRouteMiddleware([]string{"SA"}))
	contentManagementArcRouter.GET("/:content_management_arc_id/script", cont.CMAScriptListByCMAId, cm.roleBasedRouteMiddleware([]string{"SA"}))

	contentManagementArcScriptRouter := r.e.Group("/script")
	contentManagementArcScriptRouter.Use(cm.appMiddleware)
	contentManagementArcScriptRouter.PUT("/:content_management_arc_script_id", cont.UpdateCMAScript, cm.roleBasedRouteMiddleware([]string{"SA"}))
	contentManagementArcScriptRouter.DELETE("/:content_management_arc_script_id", cont.DeleteCMAScript, cm.roleBasedRouteMiddleware([]string{"SA"}))
}
