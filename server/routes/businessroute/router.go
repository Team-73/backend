package businessroute

import (
	"github.com/gin-gonic/gin"
)

// Route holds the product handlers
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

//RegisterRoutes is a routers map of user requests
func (r *Route) RegisterRoutes() {

	r.router.GET("/businesses", r.ctrl.handleGetBusinesses)
	r.router.GET("/business/:id", r.ctrl.handleGetBusinessByID)
	r.router.POST("/business", r.ctrl.handleCreateBusiness)
	r.router.PUT("/business/:id", r.ctrl.handleUpdateBusiness)
	r.router.DELETE("/business/:id", r.ctrl.handleDeleteBusiness)

}
