package controllers

import (
	"github.com/fengyfei/go-box/system"
	"github.com/gin-gonic/gin"
)

// Server controller.
type Server struct {
}

// Status reports the server statsu.
func (s *Server) Status(c *gin.Context) {
	c.JSON(200, system.Info())
}
