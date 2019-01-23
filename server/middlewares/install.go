package middlewares

import (
	"github.com/gin-gonic/gin"
)

// Install all available middlewares.
func Install(router *gin.Engine) {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	installCorsMiddleware(router)
	installJWTMiddleware(router)
}
