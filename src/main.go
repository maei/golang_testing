package main

import (
	"fmt"
	"github.com/maei/golang_testing/src/domain"
	"github.com/maei/golang_testing/src/service"
)

var (
	blogItemService domain.BlogItemInterface     = domain.NewBlogItemRepo()
	blogService     service.BlogServiceInterface = service.NewBlogService(blogItemService)
)

func main() {
	test := blogService.Validate("")
	fmt.Println(test)
}
