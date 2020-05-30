package domain

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	BlogDomain BlogDomainInterface = NewBlogDomain()
)

type BlogItem struct{}

type BlogItemRequest struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id" json:"author_id"`
	Content  string             `bson:"content" json:"content"`
	Title    string             `bson:"title" json:"title"`
}

type BlogItemResponse struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Title   string `json:"title"`
}

type BlogDomainInterface interface {
	IsTitleUnique(string) bool
	InsertOne(context.Context, interface{}) (string, error)
}

func NewBlogDomain() BlogDomainInterface {
	return &BlogItem{}
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
