package services

import (
	"github.com/yosa12978/northrend/domain"
	"github.com/yosa12978/northrend/repos"
)

type CommentService interface {
	GetComment(id string) (domain.Comment, error)
	GetComments(postId string) ([]domain.Comment, error)
	Create(comment domain.Comment) (string, error)
	Delete(id string) (string, error)
}

type commentService struct {
	commentRepo repos.CommentRepo
}

func NewCommentService() CommentService {
	return &commentService{
		commentRepo: repos.NewCommentRepo(),
	}
}

func (s *commentService) GetComment(id string) (domain.Comment, error) {
	return s.commentRepo.GetComment(id)
}

func (s *commentService) GetComments(postId string) ([]domain.Comment, error) {
	return s.commentRepo.GetComments(postId)
}

func (s *commentService) Create(comment domain.Comment) (string, error) {
	return s.commentRepo.Create(comment)
}

func (s *commentService) Delete(id string) (string, error) {
	return s.commentRepo.Delete(id)
}
