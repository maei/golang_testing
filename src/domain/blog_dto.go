package domain

import (
	"fmt"
)

var BlogItemService BlogItemInterface = &BlogItem{}

type BlogItemInterface interface {
	IsTitleUnique(string) bool
}

func NewBlogItemRepo() BlogItemInterface {
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
