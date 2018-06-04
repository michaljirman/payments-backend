package payment

import (
	"github.com/michaljirman/payments-backend/src/paymentservice/pkg/entity"
)

//Service service interface
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//Store an payment
func (s *Service) Store(b *entity.Payment) (entity.ID, error) {
	b.ID = entity.NewID()
	return s.repo.Store(b)
}

//Find a payment
func (s *Service) Find(id entity.ID) (*entity.Payment, error) {
	return s.repo.Find(id)
}

//FindAll payments
func (s *Service) FindAll() ([]*entity.Payment, error) {
	return s.repo.FindAll()
}

//Delete a payment
func (s *Service) Delete(id entity.ID) error {
	_, err := s.Find(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

// Update a payment
func (s *Service) Update(id entity.ID, p *entity.Payment) error {
	_, err := s.Find(id)
	if err != nil {
		return err
	}
	return s.repo.Update(id, p)
}
