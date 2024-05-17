package router

import "fmt"

func (r *router) initializeContentManagementRoutes() {
	fmt.Println("Initializing content management routes")
	contentManagementRouter := r.e.Group("/content-management")
	contentManagementRouter.Use(r.appMiddleware)
	contentManagementRouter.GET("", r.cont.GetContentManagementDashboard)
	contentManagementRouter.POST("", r.cont.CreateContent)
}
