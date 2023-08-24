package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	Timestamp int64              `json:"timestamp" bson:"timestamp"`
	Content   string             `json:"content" bson:"content"`
}

func NewPost(content string) Post {
	return Post{
		Id:        primitive.NewObjectID(),
		Timestamp: time.Now().Unix(),
		Content:   content,
	}
}
