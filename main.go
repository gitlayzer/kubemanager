package main

import (
	"kubemanager/LoadFiles"
	"kubemanager/controller"
	"kubemanager/middle"
	"kubemanager/service"

	"github.com/gin-gonic/gin"
)

func init() {
	LoadFiles.SetFlags()
}

func main() {
	service.K8s.Init()

	r := gin.Default()

	r.Use(middle.Cors())

	controller.Router.InitApiRouter(r)

	r.Run(LoadFiles.ReadAddress())
}
