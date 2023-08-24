package repos

import "github.com/yosa12978/northrend/domain"

type AnnounceRepo interface {
	Get() *domain.Announce
	Create(announce domain.Announce)
	Remove()
}

type announceRepo struct {
	announce *domain.Announce
}

func NewAnnounceRepo() AnnounceRepo {
	repo := new(announceRepo)
	repo.announce = nil
	return repo
}

func (repo *announceRepo) Get() *domain.Announce {
	return repo.announce
}
func (repo *announceRepo) Create(announce domain.Announce) {
	repo.announce = &announce
}
func (repo *announceRepo) Remove() {
	repo.announce = nil
}
