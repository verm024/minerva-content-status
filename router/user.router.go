package router

import (
	"fmt"
)

func (r *router) initializeUserRoute() {
	fmt.Println("Initializing user routes")
	userRouter := r.e.Group("/user")
	userRouter.GET("", r.cont.GetAllUsers)
}
