package controller

import (
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/maei/golang_testing/src/domain"
	"github.com/maei/golang_testing/src/service"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	BlogController BlogControllerInterface = NewBlogController(service.BlogService)
	blogService    service.BlogServiceInterface
)

type BlogControllerInterface interface {
	PostBlog(echo.Context) error
}

type blogController struct {
}

func NewBlogController(inc service.BlogServiceInterface) BlogControllerInterface {
	blogService = inc
	return &blogController{}
}

func (*blogController) PostBlog(c echo.Context) error {
	defer c.Request().Body.Close()
	req := c.Request().Body

	bs, err := ioutil.ReadAll(req)
	if err != nil {
		return errors.New("decode json failed")
	}

	blog := domain.BlogItem{}

	err = json.Unmarshal(bs, &blog)
	if err != nil {
		log.Println(err)
		return errors.New("unmarshal to blog-item failed")
	}

	res := blogService.ValidateBlog(blog)

	return c.JSON(http.StatusOK, map[string]string{
		"message": res,
	})
}
