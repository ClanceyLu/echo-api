package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Post 文章结构
type Post struct {
	ID      primitive.ObjectID `bson:"_id" json:"id"`
	Title   string             `bson:"title" json:"title"`
	Content string             `bson:"content" json:"content"`
}
