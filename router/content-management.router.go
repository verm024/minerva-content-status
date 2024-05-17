package router

import "fmt"

func (r *router) initializeContentManagementRoutes() {
	fmt.Println("Initializing content management routes")
	contentManagementRouter := r.e.Group("/content-management")
	contentManagementRouter.GET("", r.cont.GetContentManagementDashboard)
}
