package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Initialize(e *echo.Echo, db *gorm.DB) {
	cm := initializeCustomMiddleware(e)
	cm.initializePreMiddleware()

	contentManagementRouter := contentManagementRouter{e, db}
	contentManagementRouter.initialize()

	contentManagementArcRouter := contentManagementArcRouter{e, db}
	contentManagementArcRouter.initialize()

	userRouter := userRouter{e, db}
	userRouter.initialize()

}
