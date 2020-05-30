package controller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/maei/golang_testing/src/domain"
	"github.com/maei/golang_testing/src/service"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	BlogController BlogControllerInterface = NewBlogController(service.BlogService, echo.Validator())
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
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	blog := domain.BlogItemRequest{}

	err = json.Unmarshal(bs, &blog)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	res, error := blogService.SaveBlog(blog)
	if error != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": error.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]*domain.BlogItemResponse{
		"message": res,
	})
}
