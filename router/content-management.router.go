package router

func (r *router) initializeContentManagementRoutes() {
	contentManagementRouter := r.e.Group("/content-management")
	contentManagementRouter.Use(r.appMiddleware)
	contentManagementRouter.GET("", r.cont.GetContentManagementDashboard, r.roleBasedRouteMiddleware([]string{"SA"}))
	contentManagementRouter.POST("", r.cont.CreateContent)
	contentManagementRouter.PUT("/:content_management_id", r.cont.UpdateContent)
}
