package router

import (
	"github.com/gin-gonic/gin"
	"github.com/laironacosta/ms-gin-go/controllers"
)

type Router struct {
	server *gin.Engine
}

func NewRouter(server *gin.Engine) *Router {
	return &Router{
		server,
	}
}

func (r *Router) Init() {
	//create a default router with default middleware
	basePath := r.server.Group("/git-deploy-aws-ecs")

	basePath.GET("/health", controllers.Health)
}
