package main

import (
	"github.com/gin-gonic/gin"
	"github.com/morgances/hydra/server/controllers"
)

func installRouter(router *gin.Engine) {
	server := &controllers.Server{}

	router.POST("/status", server.Status)
}
