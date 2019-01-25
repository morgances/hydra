package middlewares

import (
	"time"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

var (
	// JWTMiddleware should be exported for user authentication.
	JWTMiddleware *jwt.GinJWTMiddleware

	skipMap = map[string]bool{
		"/api/v1/user/login": true,
	}
)

func init() {
	JWTMiddleware = &jwt.GinJWTMiddleware{
		Realm:   "Template",
		Key:     []byte("hydra"),
		Timeout: 24 * time.Hour,
		Authenticator: func(c *gin.Context) (interface{}, error) {
			glog.Info("[JWT] Authenticator triggered")
			return "testing", nil
		},
	}
}

func installJWTMiddleware(router *gin.Engine) {
	router.Use(func(c *gin.Context) {
		glog.Infof("[JWT] Middleware, path %s", c.Request.URL.Path)
		if _, skip := skipMap[c.Request.URL.Path]; skip {
			c.Next()
			return
		}

		JWTMiddleware.MiddlewareFunc()(c)
	})
}
