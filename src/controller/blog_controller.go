package controller

import "github.com/maei/golang_testing/src/service"

var (
	BlogController BlogControllerInterface = NewBlogController(service.BlogService)
	blogService    service.BlogServiceInterface
)

type BlogControllerInterface interface {
}

type blogController struct {
}

func NewBlogController(inc service.BlogServiceInterface) BlogControllerInterface {
	blogService = inc
	return &blogController{}
}
