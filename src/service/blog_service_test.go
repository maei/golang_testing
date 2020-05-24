package service

import (
	"github.com/maei/golang_testing/src/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) IsTitleUnique(testString string) bool {
	args := m.Called()
	return args.Bool(0)
}

func TestValidate(t *testing.T) {
	mockRepo := new(MockRepository)
	testString := ""

	mockRepo.On("IsTitleUnique").Return(false)

	testService := NewBlogService(mockRepo)

	res := testService.Validate(testString)
	mockRepo.AssertExpectations(t)
	assert.Equal(t, "No Good", res)
}

func TestValidateBlog(t *testing.T) {
	mockRepo := new(MockRepository)
	blog := domain.BlogItem{
		ID:      "asdasd",
		Content: "asdsd",
		Title:   "asdasdasd",
	}
	mockRepo.On("IsTitleUnique").Return(true)
	testService := NewBlogService(mockRepo)
	res := testService.ValidateBlog(blog)
	mockRepo.AssertExpectations(t)
	assert.Equal(t, "good to go", res)

}
