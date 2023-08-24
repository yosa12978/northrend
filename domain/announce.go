package domain

import (
	"time"
)

type Announce struct {
	Timestamp int64  `json:"timestamp" bson:"timestamp"`
	Content   string `json:"content"  bson:"content"`
	Expires   int64  `json:"expires" bson:"expires"`
}

func NewAnnounce(content string, expires int64) Announce {
	return Announce{
		Content:   content,
		Expires:   expires,
		Timestamp: time.Now().Unix(),
	}
}
