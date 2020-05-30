package service

import (
	"context"
	"errors"
	"github.com/maei/golang_testing/src/domain"
)

var (
	BlogService BlogServiceInterface = NewBlogService(domain.NewBlogDomain(), NewHelperService())
	blogDomain  domain.BlogDomainInterface
	helper      HelperServiceInterface
)

type BlogServiceInterface interface {
	SaveBlog(blog domain.BlogItemRequest) (*domain.BlogItemResponse, error)
}

type blogService struct{}

func NewBlogService(inc domain.BlogDomainInterface, help HelperServiceInterface) BlogServiceInterface {
	helper = help
	blogDomain = inc
	return &blogService{}
}

func (b *blogService) SaveBlog(blog domain.BlogItemRequest) (*domain.BlogItemResponse, error) {
	if !blogDomain.IsTitleUnique(blog.Content) {
		return nil, errors.New("missing content")
	}

	res, saveErr := blogDomain.InsertOne(context.Background(), blog)
	if saveErr != nil {
		return nil, saveErr
	}

	blogRes := &domain.BlogItemResponse{
		ID:      res,
		Content: blog.Content,
		Title:   blog.Title,
	}

	return blogRes, nil
}
