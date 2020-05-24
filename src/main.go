package main

import (
	"fmt"
	"github.com/maei/golang_testing/src/service"
)

func main() {
	test := service.BlogService.Validate("dsfdf")
	fmt.Println(test)
	wow := service.BlogService.SomeTest("Matthias")
	fmt.Println(wow)

}
