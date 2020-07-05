package companyroute

import (
	"github.com/gin-gonic/gin"
)

// Router holds the company handlers
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

//RegisterRoutes is a routers map of company requests
func (r *Router) RegisterRoutes() {

	r.router.GET("/companies", r.ctrl.handleGetCompanies)
	r.router.GET("/company/:id", r.ctrl.handleGetCompanyByID)
	r.router.GET("/company/:id/products", r.ctrl.handleGetProductsByCompanyID)
	r.router.POST("/company", r.ctrl.handleCreateCompany)
	r.router.PUT("/company/:id", r.ctrl.handleUpdateCompany)
	r.router.DELETE("/company/:id", r.ctrl.handleDeleteCompany)
}
