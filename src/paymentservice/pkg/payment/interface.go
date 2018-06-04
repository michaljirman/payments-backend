package payment

import (
	"github.com/michaljirman/payments-backend/src/paymentservice/pkg/entity"
)

//Reader payment reader
type Reader interface {
	Find(id entity.ID) (*entity.Payment, error)
	FindAll() ([]*entity.Payment, error)
}

//Writer payment writer
type Writer interface {
	Store(b *entity.Payment) (entity.ID, error)
	Delete(id entity.ID) error
	Update(id entity.ID, b *entity.Payment) error
}

//Repository repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase use case interface
type UseCase interface {
	Reader
	Writer
}
