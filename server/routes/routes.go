package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/morgances/hydra/server/controllers"
	"github.com/morgances/hydra/server/middlewares"
)

// Install all available routes.
func Install(router *gin.Engine) {
	server := &controllers.Server{}

	router.POST("/user/login", middlewares.JWTMiddleware.LoginHandler)

	router.POST("/insights/status", server.Status)
}
