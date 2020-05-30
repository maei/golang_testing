package mongodb

import (
	"context"
	"github.com/maei/shared_utils_go/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	MongoClient   mongoClientInterface   = &mongoClient{}
	MongoDatabase mongoDatabaseInterface = &mongoDatabase{}
)

type mongoClientInterface interface {
	setClient(*mongo.Client) *mongo.Client
	Disconnect(context.Context) error
}

type mongoDatabaseInterface interface {
	setDB(string, *mongo.Client) *mongo.Database
}

type mongoClient struct {
	client *mongo.Client
}

type mongoDatabase struct {
	db *mongo.Database
}

func InitClient(mongoURI string) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		logger.Error("Error while connecting to Database", err)
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		logger.Error("Error while connect with context", err)
		panic(err)
	}
	// Check the connection
	err = client.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	c := MongoClient.setClient(client)
	return c

}

func InitDatabase(dbName string, client *mongo.Client) *mongo.Database {
	d := MongoDatabase.setDB(dbName, client)
	return d
}

func (d *mongoDatabase) setDB(dbName string, client *mongo.Client) *mongo.Database {
	d.db = client.Database(dbName)
	return d.db

}

func (c *mongoClient) setClient(client *mongo.Client) *mongo.Client {
	c.client = client
	return c.client
}

func (c *mongoClient) Disconnect(ctx context.Context) error {
	err := c.client.Disconnect(ctx)
	if err != nil {
		logger.Error("Error disconnecting MongoDB Client", err)
		panic(err)
	}
	return nil
}
