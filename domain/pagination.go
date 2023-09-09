package domain

import "go.mongodb.org/mongo-driver/mongo/options"

type MongoPaginate struct {
	limit int64
	page  int64
}

func NewMongoPaginate(limit, page int) *MongoPaginate {
	return &MongoPaginate{
		limit: int64(limit),
		page:  int64(page),
	}
}

func (mp *MongoPaginate) GetPageOpts() *options.FindOptions {
	skip := mp.page*mp.limit - mp.limit
	return &options.FindOptions{Limit: &mp.limit, Skip: &skip}
}

type Page[b any] struct {
	items   []b
	hasNext bool
	hasPrev bool
}

func NewPage[b any](items []b, hasNext bool, hasPrev bool) Page[b] {
	return Page[b]{items: items, hasNext: hasNext, hasPrev: hasPrev}
}
