package main

import (
	"context"
	"github.com/maei/golang_testing/src/app"
	"github.com/maei/golang_testing/src/repository/mongodb"
)

func main() {
	app.StartApplication()
	mongodb.MongoClient.Disconnect(context.Background())
}
