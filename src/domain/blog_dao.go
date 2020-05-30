package domain

import (
	"context"
	"github.com/maei/golang_testing/src/repository/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func (*blogItem) InsertOne(ctx context.Context, document interface{}) (string, error) {
	log.Println("hello")
	insert, errMongo := mongodb.MongoCollectionBlog.InsertOne(ctx, document)
	if errMongo != nil {
		return "", errMongo
	}

	oid, ok := insert.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", nil
	}

	return oid.Hex(), nil
}
