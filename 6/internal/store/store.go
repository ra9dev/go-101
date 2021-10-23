package store

import (
	"context"
	"lectures-6/internal/models"
)

type Store interface {
	Laptops() LaptopsRepository
	Phones() PhonesRepository
}

// TODO дома почитать, вернемся в будущих лекциях
type LaptopsRepository interface {
	Create(ctx context.Context, laptop *models.Laptop) error
	All(ctx context.Context) ([]*models.Laptop, error)
	ByID(ctx context.Context, id int) (*models.Laptop, error)
	Update(ctx context.Context, laptop *models.Laptop) error
	Delete(ctx context.Context, id int) error
}

type PhonesRepository interface {
	Create(ctx context.Context, phone *models.Phone) error
	All(ctx context.Context) ([]*models.Phone, error)
	ByID(ctx context.Context, id int) (*models.Phone, error)
	Update(ctx context.Context, laptop *models.Phone) error
	Delete(ctx context.Context, id int) error
}
