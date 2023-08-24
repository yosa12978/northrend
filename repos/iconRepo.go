package repos

import (
	"github.com/yosa12978/northrend/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type IconRepo interface {
}

type iconRepoMongo struct {
	db *mongo.Database
}

func NewIconRepo() IconRepo {
	repo := new(iconRepoMongo)
	repo.db = db.GetDB()
	return repo
}
