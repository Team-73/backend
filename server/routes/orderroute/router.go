package orderroute

import (
	"github.com/gin-gonic/gin"
)

// Route holds the ping handlers
type Route struct {
	ctrl   *Controller
	router *gin.Engine
}

// NewRoute returns a new Route instance
func NewRoute(ctrl *Controller, router *gin.Engine) *Route {
	return &Route{
		ctrl:   ctrl,
		router: router,
	}
}

//RegisterRoutes is a routers map of ping requests
func (r *Route) RegisterRoutes() {
	r.router.POST("/order", r.ctrl.handleCreateOrder)
	r.router.GET("/order/detail/:id", r.ctrl.handleGetOrderDetail)

	r.router.GET("/orders/user/:id", r.ctrl.handleGetOrdersByUserID)
}
