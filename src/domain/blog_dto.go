package domain

import (
	"fmt"
)

var (
	BlogItemDomain BlogItemDomainInterface = NewBlogItemDomain()
)

type BlogItemDomainInterface interface {
	IsTitleUnique(string) bool
}

func NewBlogItemDomain() BlogItemDomainInterface {
	return &BlogItem{}
}

type BlogItem struct {
	ID      string
	Content string
	Title   string
}

func (b *BlogItem) IsTitleUnique(testString string) bool {
	fmt.Println("Call ValidateContent")
	if testString == "" {
		return false
	}
	return true
}

func (b *BlogItem) testIng() {

}
