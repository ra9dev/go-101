package store

import (
	"context"
	"lectures-6/internal/models"
)

type Store interface {
	Create(ctx context.Context, laptop *models.Laptop) error
	All(ctx context.Context) ([]*models.Laptop, error)
	ByID(ctx context.Context, id int) (*models.Laptop, error)
	Update(ctx context.Context, laptop *models.Laptop) error
	Delete(ctx context.Context, id int) error

	// Laptops() LaptopsRepository
	// Phones() PhonesRepository
}

// electronics
//   laptops
//   phones

// TODO дома почитать, вернемся в будущих лекциях
// type LaptopsRepository interface {
// 	Create(ctx context.Context, laptop *models.Laptop) error
// 	All(ctx context.Context) ([]*models.Laptop, error)
// 	ByID(ctx context.Context, id int) (*models.Laptop, error)
// 	Update(ctx context.Context, laptop *models.Laptop) error
// 	Delete(ctx context.Context, id int) error
// }
