package main

import (
	"kubemanager/config"
	"kubemanager/controller"
	"kubemanager/middle"
	"kubemanager/service"

	"github.com/gin-gonic/gin"
)

func main() {
	service.K8s.Init()

	r := gin.Default()

	r.Use(middle.Cors())

	controller.Router.InitApiRouter(r)

	r.Run(config.ListenAddr)
}
