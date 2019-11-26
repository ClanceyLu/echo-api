package post

import (
	"github.com/ClanceyLu/echo-api/service"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type postService struct {
	mongo *mongo.Client
}

func New(mongo *mongo.Client) service.Service {
	return &postService{mongo}
}

func (p postService) Router(r *echo.Group) {
	r.POST("/post", p.postPost)
}
