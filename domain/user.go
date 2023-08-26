package domain

import (
	"github.com/yosa12978/northrend/helpers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	ROLE_USER  = "USER"
	ROLE_ADMIN = "ADMIN"
)

type User struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"-" bson:"password"`
	Role     string             `json:"role" bson:"role"`
}

func NewUser(username, password, role string) User {
	return User{
		Id:       primitive.NewObjectID(),
		Username: username,
		Password: helpers.MD5Hash(password),
		Role:     role,
	}
}
