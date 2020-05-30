package app

import (
	"github.com/labstack/echo/v4"
	"github.com/maei/golang_testing/src/repository/mongodb"
)

var (
	router = echo.New()
)

func StartApplication() {
	mongoClient := mongodb.InitClient("mongodb://localhost:27017")
	mongoDB := mongodb.InitDatabase("mydb", mongoClient)
	mongodb.InitCollectionBlog("blog", mongoDB)

	mapUrls()
	router.Logger.Fatal(router.Start(":8002"))
}
