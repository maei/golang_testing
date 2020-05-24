package service

import (
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
	testString := "gdfgdfg"

	mockRepo.On("IsTitleUnique").Return(true)

	testService := NewBlogService(mockRepo)

	res := testService.Validate(testString)
	mockRepo.AssertExpectations(t)
	assert.Equal(t, "No Good", res)
}
