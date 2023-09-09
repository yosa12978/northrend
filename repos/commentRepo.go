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

type CommentRepo interface {
	GetComment(id string) (domain.Comment, error)
	GetComments(postId string) ([]domain.Comment, error)
	Create(comment domain.Comment) (string, error)
	Delete(id string) (string, error)
	GetCommentsPaginated(postId string, page, limit int) (domain.Page[domain.Comment], error)
}

type commentRepoMongo struct {
	db *mongo.Database
}

func NewCommentRepo() CommentRepo {
	repo := new(commentRepoMongo)
	repo.db = db.GetDB()
	return repo
}

func (repo *commentRepoMongo) GetComment(id string) (domain.Comment, error) {
	var comment domain.Comment
	objid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return comment, err
	}
	err = repo.db.Collection("comments").FindOne(context.TODO(), bson.M{"_id": objid}).Decode(&comment)
	return comment, err
}

func (repo *commentRepoMongo) GetComments(postId string) ([]domain.Comment, error) {
	var comments []domain.Comment
	objid, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return comments, err
	}
	fopts := options.Find().SetSort(bson.M{"_id": -1})
	cursor, _ := repo.db.Collection("comments").Find(context.TODO(), bson.M{"postId": objid}, fopts)
	err = cursor.All(context.TODO(), &comments)
	return comments, err
}

func (repo *commentRepoMongo) Create(comment domain.Comment) (string, error) {
	res, err := repo.db.Collection("comments").InsertOne(context.TODO(), comment)
	return res.InsertedID.(primitive.ObjectID).Hex(), err
}

func (repo *commentRepoMongo) Delete(id string) (string, error) {
	objid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}
	_, err = repo.db.Collection("comments").DeleteOne(context.TODO(), bson.M{"_id": objid})
	return id, err
}

func (repo *commentRepoMongo) GetCommentsPaginated(postId string, page, limit int) (domain.Page[domain.Comment], error) {
	var comments []domain.Comment
	objid, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return domain.NewPage(comments, false, false), err
	}

	copts := options.Count().SetHint("_id")
	total, _ := repo.db.Collection("comments").CountDocuments(context.TODO(), bson.M{}, copts)

	fopts := domain.NewMongoPaginate(limit, page).GetPageOpts()
	cursor, _ := repo.db.Collection("comments").Find(context.TODO(), bson.M{"postId": objid}, fopts)
	err = cursor.All(context.TODO(), &comments)

	res := domain.NewPage(comments, int64(page) < ((total/int64(limit))+1), page > 1)

	return res, err
}
