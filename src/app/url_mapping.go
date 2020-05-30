package app

import (
	"github.com/maei/golang_testing/src/controller"
)

func publicRoutes() {
	router.GET("/post", controller.BlogController.PostBlog)
}

func mapUrls() {
	publicRoutes()
}
