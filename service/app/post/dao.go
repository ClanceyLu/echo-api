package post

import (
	"context"

	"github.com/ClanceyLu/echo-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (p post) insertPost(post *model.Post) (interface{}, error) {
	// ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	mongo := p.Mongo
	collection := mongo.Database("echo-api").Collection("posts")
	res, err := collection.InsertOne(context.TODO(), post)
	if err != nil {
		return nil, err
	}
	return res.InsertedID, nil
}

func (p post) findPosts() (*[]*model.Post, error) {
	mongo := p.Mongo
	collection := mongo.Database("echo-api").Collection("posts")
	var list []*model.Post
	findOptions := options.Find()
	findOptions.SetLimit(2)
	res, err := collection.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		return &list, err
	}
	for res.Next(context.TODO()) {
		var po model.Post
		err := res.Decode(&po)
		if err != nil {
			return &list, err
		}
		list = append(list, &po)
	}
	return &list, nil
}
