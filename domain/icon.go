package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Icon struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	Timestamp int64              `json:"timestamp" bson:"timestamp"`
	Path      string             `json:"path" bson:"path"`
}

func NewIcon(path string) Icon {
	return Icon{
		Id:        primitive.NewObjectID(),
		Timestamp: time.Now().Unix(),
		Path:      path,
	}
}
