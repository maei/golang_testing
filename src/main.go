package main

import (
	"fmt"
	"github.com/maei/golang_testing/src/service"
)

/*var (
	blogItemServicePub domain.BlogItemInterface = domain.NewBlogItemRepo()
	blogService     service.BlogServiceInterface = service.NewBlogService(blogItemServicePub)
)*/

func main() {
	test := service.BlogService.Validate("dsfdf")
	fmt.Println(test)
	wow := service.BlogService.SomeTest("Matthias")
	fmt.Println(wow)

}
