package service

import (
	"fmt"
	"github.com/maei/golang_testing/src/domain"
)

var (
	BlogService blogServiceInterface = NewBlogService(domain.NewBlogItemRepo())
	blog        domain.BlogItemInterface
)

type blogServiceInterface interface {
	Validate(string) string
	SomeTest(data string) string
}

type blogService struct{}

func NewBlogService(inc domain.BlogItemInterface) blogServiceInterface {
	blog = inc
	return &blogService{}
}

func (b *blogService) Validate(title string) string {
	if !blog.IsTitleUnique(title) {
		return "No Good"
	}
	return "good to go"
}

func (b *blogService) SomeTest(data string) string {
	return fmt.Sprintf("Hello %v", data)
}
