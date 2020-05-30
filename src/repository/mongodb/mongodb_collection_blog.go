package mongodb

import (
	"context"
	"github.com/maei/shared_utils_go/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	MongoCollectionBlog mongoCollectionBlogInterface = &mongoCollectionBlog{}
)

type mongoCollectionBlogInterface interface {
	setCollectionBlog(string, *mongo.Database)
	InsertOne(context.Context, interface{}) (*mongo.InsertOneResult, error)
}

type mongoCollectionBlog struct {
	coll *mongo.Collection
}

func InitCollectionBlog(colName string, db *mongo.Database) {
	MongoCollectionBlog.setCollectionBlog(colName, db)
}

func (collection *mongoCollectionBlog) setCollectionBlog(colName string, db *mongo.Database) {
	collection.coll = db.Collection(colName)
}

func (collection *mongoCollectionBlog) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	res, err := collection.coll.InsertOne(ctx, document)
	if err != nil {
		logger.Error("MongoDB: Cannot InsertOne operation", err)
		return nil, err
	}
	return res, err
}
