package services

import (
	"github.com/yosa12978/northrend/domain"
	"github.com/yosa12978/northrend/repos"
)

type PostService interface {
	GetPosts() []domain.Post
	GetPost(id string) (domain.Post, error)
	DeletePost(id string) (string, error)
	CreatePost(post domain.Post) (string, error)
}

type postService struct {
	postRepo repos.PostRepo
}

func NewPostService() PostService {
	return &postService{
		postRepo: repos.NewPostRepo(),
	}
}

func (service *postService) GetPosts() []domain.Post {
	return service.postRepo.GetPosts()
}

func (service *postService) GetPost(id string) (domain.Post, error) {
	return service.postRepo.GetPost(id)
}

func (service *postService) DeletePost(id string) (string, error) {
	return service.postRepo.DeletePost(id)
}

func (service *postService) CreatePost(post domain.Post) (string, error) {
	return service.postRepo.CreatePost(post)
}
