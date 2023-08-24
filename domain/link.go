package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Link struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	Timestamp int64              `json:"timestamp" bson:"timestamp"`
	Name      string             `json:"name" bson:"name"`
	Uri       string             `json:"uri" bson:"uri"`
}

func NewLink(name string, uri string) Link {
	return Link{
		Id:        primitive.NewObjectID(),
		Timestamp: time.Now().Unix(),
		Name:      name,
		Uri:       uri,
	}
}
