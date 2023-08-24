package services

import (
	"github.com/yosa12978/northrend/domain"
	"github.com/yosa12978/northrend/repos"
)

type AnnounceService interface {
	Get() *domain.Announce
	Create(announce domain.Announce)
	Remove()
}

type announceService struct {
	announceRepo repos.AnnounceRepo
}

func NewAnnounceService() AnnounceService {
	s := new(announceService)
	s.announceRepo = repos.NewAnnounceRepo()
	return s
}

func (s *announceService) Get() *domain.Announce {
	return s.announceRepo.Get()
}

func (s *announceService) Create(announce domain.Announce) {
	s.announceRepo.Create(announce)
}

func (s *announceService) Remove() {
	s.announceRepo.Remove()
}
