package post

import (
	"context"
	"time"

	"github.com/ClanceyLu/echo-api/model"
)

func (p postService) insertPost(post *model.Post) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collection := p.mongo.Database("echo-api").Collection("posts")
	res, err := collection.InsertOne(ctx, post)
	if err != nil {
		return nil, err
	}
	return res.InsertedID, nil
}
