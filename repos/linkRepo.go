package repos

import (
	"context"

	"github.com/yosa12978/northrend/db"
	"github.com/yosa12978/northrend/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LinkRepo interface {
	CreateLink(link domain.Link) (string, error)
	DeleteLink(id string) (string, error)
	GetLinks() []domain.Link
}

type linkRepoMongo struct {
	db *mongo.Database
}

func NewLinkRepo() LinkRepo {
	repo := new(linkRepoMongo)
	repo.db = db.GetDB()
	return repo
}

func (repo *linkRepoMongo) CreateLink(link domain.Link) (string, error) {
	res, err := repo.db.Collection("links").InsertOne(context.TODO(), link)
	return res.InsertedID.(primitive.ObjectID).Hex(), err
}

func (repo *linkRepoMongo) DeleteLink(id string) (string, error) {
	objid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}
	_, err = repo.db.Collection("links").DeleteOne(context.TODO(), bson.M{"_id": objid})
	return id, err
}

func (repo *linkRepoMongo) GetLinks() []domain.Link {
	var links []domain.Link
	cursor, _ := repo.db.Collection("links").Find(context.TODO(), bson.M{})
	cursor.All(context.TODO(), &links)
	return links
}
