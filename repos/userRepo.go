package repos

import (
	"context"

	"github.com/yosa12978/northrend/db"
	"github.com/yosa12978/northrend/domain"
	"github.com/yosa12978/northrend/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo interface {
	GetUserByUsername(username string) (domain.User, error)
	GetUser(username, password string) (domain.User, error)
	IsAdmin(username string) bool
	IsUserExist(username, password string) bool
	IsUsernameTaken(username string) bool
	CreateUser(user domain.User) (string, error)
	DeleteUser(id string) (string, error)
	Seed() error
}

type userRepo struct {
	db *mongo.Database
}

func NewUserRepo() UserRepo {
	return &userRepo{
		db: db.GetDB(),
	}
}

func (repo *userRepo) GetUserByUsername(username string) (domain.User, error) {
	var user domain.User
	err := repo.db.Collection("users").FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	return user, err
}

func (repo *userRepo) GetUser(username, password string) (domain.User, error) {
	var user domain.User
	filter := bson.M{"username": username, "password": helpers.MD5Hash(password)}
	err := repo.db.Collection("users").FindOne(context.TODO(), filter).Decode(&user)
	return user, err
}

func (repo *userRepo) IsAdmin(username string) bool {
	var user domain.User
	filter := bson.M{"username": username}
	repo.db.Collection("users").FindOne(context.TODO(), filter).Decode(&user)
	return user.Role == domain.ROLE_ADMIN
}

func (repo *userRepo) IsUserExist(username, password string) bool {
	var user domain.User
	filter := bson.M{"username": username, "password": helpers.MD5Hash(password)}
	err := repo.db.Collection("users").FindOne(context.TODO(), filter).Decode(&user)
	return err == nil
}

func (repo *userRepo) IsUsernameTaken(username string) bool {
	var user domain.User
	filter := bson.M{"username": username}
	err := repo.db.Collection("users").FindOne(context.TODO(), filter).Decode(&user)
	return err == nil
}

func (repo *userRepo) CreateUser(user domain.User) (string, error) {
	res, err := repo.db.Collection("users").InsertOne(context.TODO(), user)
	return res.InsertedID.(primitive.ObjectID).Hex(), err
}

func (repo *userRepo) DeleteUser(id string) (string, error) {
	objid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}
	_, err = repo.db.Collection("users").DeleteOne(context.TODO(), bson.M{"_id": objid})
	return id, err
}

func (repo *userRepo) Seed() error {
	user := domain.NewUser("admin", "admin", domain.ROLE_ADMIN)
	if repo.IsUsernameTaken(user.Username) {
		return nil
	}
	_, err := repo.db.Collection("users").InsertOne(context.TODO(), user)
	return err
}
