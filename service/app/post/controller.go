package post

import (
	"log"
	"net/http"

	"github.com/ClanceyLu/echo-api/model"
	"github.com/labstack/echo/v4"
)

func (p post) PostPost(c echo.Context) error {
	post := &model.Post{}
	if err := c.Bind(post); err != nil {
		return err
	}
	log.Print(post)
	id, err := p.insertPost(post)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"id": id,
	})
}

func (p post) GetPosts(c echo.Context) error {
	posts, err := p.findPosts()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"list": posts,
	})
}
