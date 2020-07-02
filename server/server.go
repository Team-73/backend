package server

import (
	"github.com/Team-73/backend/server/routes/pingroute"
	"github.com/Team-73/backend/service"
	"github.com/gin-gonic/gin"
)

type controller struct {
	pingController *pingroute.Controller
}

//InitServer to initialize the server
func InitServer(svc *service.Service) *gin.Engine {
	svm := service.NewServiceManager()
	_ = svm //remover essa linha quando utilizar o svm
	srv := gin.Default()

	return setupRoutes(srv, &controller{
		pingController: pingroute.NewController(),
	})
}

//setupRoutes - Register and instantiate the routes
func setupRoutes(srv *gin.Engine, s *controller) *gin.Engine {

	pingroute.NewRouter(s.pingController, srv).RegisterRoutes()

	return srv
}
