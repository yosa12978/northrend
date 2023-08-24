package repos

import (
	"context"

	"github.com/yosa12978/northrend/db"
	"github.com/yosa12978/northrend/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PostRepo interface {
	GetPosts() []domain.Post
	CreatePost(post domain.Post) (string, error)
	DeletePost(id string) (string, error)
	GetPost(id string) (domain.Post, error)
}

type postRepoMongo struct {
	db *mongo.Database
}

func NewPostRepo() PostRepo {
	repo := new(postRepoMongo)
	repo.db = db.GetDB()
	return repo
}

func (repo *postRepoMongo) GetPosts() []domain.Post {
	var posts []domain.Post
	fopts := options.Find().SetSort(bson.M{"_id": -1})
	cursor, _ := repo.db.Collection("posts").Find(context.TODO(), bson.M{}, fopts)
	cursor.All(context.TODO(), &posts)
	return posts
}

func (repo *postRepoMongo) CreatePost(post domain.Post) (string, error) {
	res, err := repo.db.Collection("posts").InsertOne(context.TODO(), post)
	return res.InsertedID.(primitive.ObjectID).Hex(), err
}

func (repo *postRepoMongo) DeletePost(id string) (string, error) {
	objid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}
	_, err = repo.db.Collection("posts").DeleteOne(context.TODO(), bson.M{"_id": objid})
	return id, err
}

func (repo *postRepoMongo) GetPost(id string) (domain.Post, error) {
	var post domain.Post
	objid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return post, err
	}
	err = repo.db.Collection("posts").FindOne(context.TODO(), bson.M{"_id": objid}).Decode(&post)
	return post, err
}
