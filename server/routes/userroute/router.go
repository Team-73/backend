package userroute

import (
	"github.com/gin-gonic/gin"
)

// Router holds the user handlers
type Router struct {
	ctrl   *Controller
	router *gin.Engine
}

// NewRoute returns a new Route instance
func NewRoute(ctrl *Controller, router *gin.Engine) *Router {
	return &Router{
		ctrl:   ctrl,
		router: router,
	}
}

//RegisterRoutes is a routers map of user requests
func (r *Router) RegisterRoutes() {

	r.router.GET("/users", r.ctrl.handleGetUsers)
	r.router.GET("/user/:id", r.ctrl.handleGetUserByID)
	r.router.POST("/user", r.ctrl.handleCreateUser)
	r.router.POST("/user/login", r.ctrl.handleLogin)
	r.router.PUT("/user/:id", r.ctrl.handleUpdateUser)
	r.router.DELETE("/user/:id", r.ctrl.handleDeleteUser)

}
