package service

import (
	"context"
	"github.com/maei/golang_testing/src/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockBlogDomain struct {
	mock.Mock
}

type MockHelpService struct {
	mock.Mock
}

func (m *MockBlogDomain) IsTitleUnique(testString string) bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockBlogDomain) InsertOne(ctx context.Context, document interface{}) (string, error) {
	args := m.Called()
	return args.String(0), nil
}

func (m *MockHelpService) DoSomething(data string) bool {
	args := m.Called()
	return args.Bool(0)
}

func TestBlogService_SaveBlog(t *testing.T) {
	mockRepo := new(MockBlogDomain)
	mockHelper := new(MockHelpService)

	blog := domain.BlogItemRequest{
		AuthorID: "test_authorid",
		Content:  "test_content",
		Title:    "test_title",
	}
	mockRepo.On("IsTitleUnique").Return(true)
	mockRepo.On("InsertOne", mock.Anything, mock.AnythingOfType("string")).Return("5ed280ed525237966e2be780", nil)

	testService := NewBlogService(mockRepo, mockHelper)
	res, _ := testService.SaveBlog(blog)
	mockRepo.AssertExpectations(t)
	assert.Equal(t, res.ID, "5ed280ed525237966e2be780")
	assert.Equal(t, res.Content, "test_content")
	assert.Equal(t, res.Title, "test_title")
}
