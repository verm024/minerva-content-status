package router

import (
	"fmt"
	"minerva-content-status/controllers"
)

func (r *router) initializeUserRoute() {
	fmt.Println("Initializing user routes")
	userRouter := r.e.Group("/user")
	userRouter.GET("/", controllers.GetAllUsers)
}
