package userroute

import (
	"github.com/gin-gonic/gin"
)

// UserRouter holds the user handlers
type UserRouter struct {
	ctrl   *Controller
	router *gin.Engine
}

// NewRouter returns a new UserRouter instance
func NewRouter(ctrl *Controller, router *gin.Engine) *UserRouter {
	return &UserRouter{
		ctrl:   ctrl,
		router: router,
	}
}

//RegisterRoutes is a routers map of user requests
func (r *UserRouter) RegisterRoutes() {

	r.router.GET("/users/:id", r.ctrl.handleGetUser)

	r.router.POST("/users", r.ctrl.handleCreateUser)
	r.router.POST("/users/login", r.ctrl.handleLogin)

	r.router.PUT("/users/:id", r.ctrl.handleUpdateUser)   //PUT 		we update every fields
	r.router.PATCH("/users/:id", r.ctrl.handleUpdateUser) //PATCH 	we update a partial fields, just few fileds
	r.router.DELETE("/users/:id", r.ctrl.handleDeleteUser)

}
