package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func main() {
	router := gin.Default()

	installRouter(router)

	s := &http.Server{
		Addr:    configuration.Address,
		Handler: router,
	}

	if err := s.ListenAndServe(); err != nil {
		glog.Fatalf("server error %s", err)
	}
}
