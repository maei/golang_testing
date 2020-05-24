package service

import "github.com/maei/golang_testing/src/domain"

var (
	BlogService BlogServiceInterface = &blogService{}
	blog        domain.BlogItemInterface
)

type BlogServiceInterface interface {
	Validate(string) string
}

type blogService struct{}

func NewBlogService(inc domain.BlogItemInterface) BlogServiceInterface {
	blog = inc
	return &blogService{}
}

func (b *blogService) Validate(title string) string {
	if !blog.IsTitleUnique(title) {
		//if !domain.BlogItemService.IsTitleUnique(title) {
		return "No Good"
	}
	return "good to go"
}
