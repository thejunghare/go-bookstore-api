package routes

// This file contains all the routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thejunghare/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(r *gin.Engine) {
	r.GET("/book/", controllers.GetBooks)
	r.GET("/book/:id", controllers.GetBookByID)
	r.PUT("/book/", controllers.CreateBook)
	// r.PUT("/book/:id", controllers.UpdatedBook)
	r.DELETE("/book/:id", controllers.DeleteBook)
}
