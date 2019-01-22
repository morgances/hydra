package controllers

import (
	"net/http"

	"github.com/fengyfei/go-box/system"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

// Server controller.
type Server struct {
}

// Status reports the server statsu.
func (s *Server) Status(c *gin.Context) {
	var req struct {
		Token string `json:"token"`
	}

	if err := c.BindJSON(&req); err != nil {
		glog.Errorf("[status] parameter error:%s", err)

		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(200, system.Info())
}
