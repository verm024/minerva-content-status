package router

import (
	"fmt"
	"minerva-content-status/controllers"
)

func (r *router) initializeUserRoute() {
	fmt.Println("Initializing user routes")
	userRouter := r.e.Group("/user")
	cont := controllers.Initialize(r.e, r.db)
	userRouter.GET("", cont.GetAllUsers)
}
