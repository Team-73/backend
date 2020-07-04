package server

import (
	"github.com/Team-73/backend/server/routes/categoryroute"
	"github.com/Team-73/backend/server/routes/pingroute"
	"github.com/Team-73/backend/server/routes/productroute"
	"github.com/Team-73/backend/server/routes/userroute"
	"github.com/Team-73/backend/service"
	"github.com/gin-gonic/gin"
)

type controller struct {
	pingController     *pingroute.Controller
	userController     *userroute.Controller
	productController  *productroute.Controller
	categoryController *categoryroute.Controller
}

//InitServer to initialize the server
func InitServer(svc *service.Service) *gin.Engine {
	svm := service.NewServiceManager()
	_ = svm //remover essa linha quando utilizar o svm
	srv := gin.Default()

	return setupRoutes(srv, &controller{
		pingController:     pingroute.NewController(),
		userController:     userroute.NewController(svm.UserService(svc)),
		productController:  productroute.NewController(svm.ProductService(svc)),
		categoryController: categoryroute.NewController(svm.CategoryService(svc)),
	})
}

//setupRoutes - Register and instantiate the routes
func setupRoutes(srv *gin.Engine, s *controller) *gin.Engine {

	pingroute.NewRoute(s.pingController, srv).RegisterRoutes()
	userroute.NewRoute(s.userController, srv).RegisterRoutes()
	productroute.NewRoute(s.productController, srv).RegisterRoutes()
	categoryroute.NewRoute(s.categoryController, srv).RegisterRoutes()

	return srv
}
