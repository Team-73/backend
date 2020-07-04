package categoryroute

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

	r.router.GET("/categories", r.ctrl.handleGetCategories)
	r.router.GET("/category/:id", r.ctrl.handleGetCategoryByID)
	r.router.POST("/category", r.ctrl.handleCreateCategory)
	r.router.PUT("/category/:id", r.ctrl.handleUpdateCategory)
	r.router.DELETE("/category/:id", r.ctrl.handleDeleteCategory)

}
