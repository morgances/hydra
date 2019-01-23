package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/morgances/hydra/server/controllers"
)

// Install all available routes.
func Install(router *gin.Engine) {
	server := &controllers.Server{}

	router.POST("/status", server.Status)
}
