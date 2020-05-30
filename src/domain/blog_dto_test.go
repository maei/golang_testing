package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsTitleUnique(t *testing.T) {
	testService := NewBlogDomainDTO()

	// Fail Testcase
	failString := ""
	res := testService.IsTitleUnique(failString)
	assert.Equal(t, false, res)

	// Pass Testcase
	validString := "vaild"
	res1 := testService.IsTitleUnique(validString)
	assert.Equal(t, true, res1)
}
