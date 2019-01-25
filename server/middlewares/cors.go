package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func installCorsMiddleware(router *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://192.168.0.108:8000"}
	config.AllowCredentials = true

	router.Use(cors.New(config))
}
