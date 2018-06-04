package payment

import (
	"github.com/michaljirman/payments-backend/src/paymentservice/pkg/entity"
)

//IRepo in memory repo
type IRepo struct {
	m map[string]*entity.Payment
}

//NewInmemRepository create new repository
func NewInmemRepository() *IRepo {
	var m = map[string]*entity.Payment{}
	return &IRepo{
		m: m,
	}
}

//Store a Bookmark
func (r *IRepo) Store(a *entity.Payment) (entity.ID, error) {
	r.m[a.ID.String()] = a
	return a.ID, nil
}

//Find a Bookmark
func (r *IRepo) Find(id entity.ID) (*entity.Payment, error) {
	if r.m[id.String()] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id.String()], nil
}

//FindAll Bookmarks
func (r *IRepo) FindAll() ([]*entity.Payment, error) {
	var d []*entity.Payment
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

//Delete a Bookmark
func (r *IRepo) Delete(id entity.ID) error {
	if r.m[id.String()] == nil {
		return entity.ErrNotFound
	}
	r.m[id.String()] = nil
	return nil
}

func (r *IRepo) Update(id entity.ID, b *entity.Payment) error {
	return nil
}
