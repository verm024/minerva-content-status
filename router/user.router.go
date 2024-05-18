package router

func (r *router) initializeUserRoute() {
	userRouter := r.e.Group("/user")
	userRouter.GET("", r.cont.GetAllUsers)
	userRouter.POST("", r.cont.RegisterNewUser)
	userRouter.POST("/login", r.cont.Login)
}
