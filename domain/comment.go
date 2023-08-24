package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	Timestamp int64              `json:"timestamp" bson:"timestamp"`
	Email     string             `json:"email" bson:"email"`
	Name      string             `json:"name" bson:"name"`
	Content   string             `json:"content" bson:"content"`
	Post      primitive.ObjectID `json:"postId" bson:"postId,omitempty"`
}

func NewComment(email, name, content string, postId primitive.ObjectID) Comment {
	return Comment{
		Id:        primitive.NewObjectID(),
		Timestamp: time.Now().Unix(),
		Email:     email,
		Name:      name,
		Content:   content,
		Post:      postId,
	}
}
