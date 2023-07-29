package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/thejunghare/go-bookstore/pkg/routes"
)

func main() {
	r := gin.Default()

	routes.RegisterBookStoreRoutes(r)

	// start the gin server
	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
