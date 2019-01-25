package middlewares

import (
	"time"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"

	"github.com/morgances/hydra/server/services"
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
			var req struct {
				User string `json:"username"`
				Pass string `json:"password"`
			}

			if err := c.BindJSON(&req); err != nil {
				glog.Errorf("[status] parameter error:%s", err)

				return nil, err
			}

			uid, err := services.AdminService.Login(&req.User, &req.Pass)
			if err != nil {
				glog.Errorf("[status] parameter error:%s", err)
				return nil, err
			}

			return uid, nil
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
