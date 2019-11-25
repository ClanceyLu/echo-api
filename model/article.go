package model

// Article 文章 model
type Article struct {
	Title   string `json:"title" bson:"title"`
	Content string `json:"content" bson:"content"`
}
