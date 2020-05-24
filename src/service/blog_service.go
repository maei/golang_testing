package service

import (
	"github.com/maei/golang_testing/src/domain"
)

var (
	BlogService BlogServiceInterface = NewBlogService(domain.NewBlogItemDomain())
	blogDomain  domain.BlogItemDomainInterface
)

type BlogServiceInterface interface {
	Validate(string) string
	ValidateBlog(blog domain.BlogItem) string
}

type blogService struct{}

func NewBlogService(inc domain.BlogItemDomainInterface) BlogServiceInterface {
	blogDomain = inc
	return &blogService{}
}

func (b *blogService) Validate(title string) string {
	if !blogDomain.IsTitleUnique(title) {
		return "No Good"
	}
	return "good to go"
}

func (b *blogService) ValidateBlog(blog domain.BlogItem) string {
	if !blogDomain.IsTitleUnique(blog.Content) {
		return "No Good"
	}
	return "good to go"
}
