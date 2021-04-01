package main

import (
	"github.com/gin-gonic/gin"
	"github.com/laironacosta/git-deploy-aws-ecs/router"
)

func main() {
	gin := gin.Default()

	r := router.NewRouter(gin)
	r.Init()

	gin.Run() // listen and serve on 0.0.0.0:8080
}
