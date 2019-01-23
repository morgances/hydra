package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"

	"github.com/morgances/hydra/server/middlewares"
	"github.com/morgances/hydra/server/routes"
)

func main() {
	router := gin.Default()

	middlewares.Install(router)
	routes.Install(router)

	s := &http.Server{
		Addr:    configuration.Address,
		Handler: router,
	}

	if err := s.ListenAndServe(); err != nil {
		glog.Fatalf("server error %s", err)
	}
}
