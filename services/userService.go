package services

import (
	"github.com/yosa12978/northrend/domain"
	"github.com/yosa12978/northrend/repos"
)

type UserService interface {
	GetUserByUsername(username string) (domain.User, error)
	GetUser(username, password string) (domain.User, error)
	IsAdmin(username string) bool
	IsUserExist(username, password string) bool
	IsUsernameTaken(username string) bool
	CreateUser(user domain.User) (string, error)
	DeleteUser(id string) (string, error)
	Seed() error
}

type userService struct {
	userRepo repos.UserRepo
}

func NewUserService() UserService {
	s := new(userService)
	s.userRepo = repos.NewUserRepo()
	return s
}

func (s *userService) GetUserByUsername(username string) (domain.User, error) {
	return s.userRepo.GetUserByUsername(username)
}

func (s *userService) GetUser(username, password string) (domain.User, error) {
	return s.userRepo.GetUser(username, password)
}

func (s *userService) IsAdmin(username string) bool {
	return s.userRepo.IsAdmin(username)
}

func (s *userService) IsUserExist(username, password string) bool {
	return s.userRepo.IsUserExist(username, password)
}

func (s *userService) IsUsernameTaken(username string) bool {
	return s.userRepo.IsUsernameTaken(username)
}

func (s *userService) CreateUser(user domain.User) (string, error) {
	return s.userRepo.CreateUser(user)
}

func (s *userService) DeleteUser(id string) (string, error) {
	return s.userRepo.DeleteUser(id)
}

func (s *userService) Seed() error {
	return s.userRepo.Seed()
}
