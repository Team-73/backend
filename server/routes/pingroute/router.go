package pingroute

import (
	"github.com/gin-gonic/gin"
)

// PingRouter holds the ping handlers
type PingRouter struct {
	ctrl   *Controller
	router *gin.Engine
}

// NewRouter returns a new PingRouter instance
func NewRouter(ctrl *Controller, router *gin.Engine) *PingRouter {
	return &PingRouter{
		ctrl:   ctrl,
		router: router,
	}
}

//RegisterRoutes is a routers map of ping requests
func (r *PingRouter) RegisterRoutes() {
	r.router.GET("/ping", r.ctrl.handlePing)
}
