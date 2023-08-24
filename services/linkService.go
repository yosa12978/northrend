package services

import (
	"github.com/yosa12978/northrend/domain"
	"github.com/yosa12978/northrend/repos"
)

type LinkService interface {
	GetLinks() []domain.Link
	CreateLink(link domain.Link) (string, error)
	DeleteLink(id string) (string, error)
}

type linkService struct {
	linkRepo repos.LinkRepo
}

func NewLinkService() LinkService {
	return &linkService{
		linkRepo: repos.NewLinkRepo(),
	}
}

func (service *linkService) GetLinks() []domain.Link {
	return service.linkRepo.GetLinks()
}

func (service *linkService) CreateLink(link domain.Link) (string, error) {
	return service.linkRepo.CreateLink(link)
}

func (service *linkService) DeleteLink(id string) (string, error) {
	return service.linkRepo.DeleteLink(id)
}
