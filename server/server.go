package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	installRouter(router)

	s := &http.Server{
		Addr:    configuration.Address,
		Handler: router,
	}
	s.ListenAndServe()
}
