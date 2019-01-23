package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func installCorsMiddleware(router *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://*", "https://*"}

	router.Use(cors.New(config))
}
