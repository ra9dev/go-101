package store

import (
	"context"
	"lecture-8/internal/models"
)

type Store interface {
	Connect(url string) error
	Close() error

	Categories() CategoriesRepository
	Goods() GoodsRepository
}

type CategoriesRepository interface {
	Create(ctx context.Context, category *models.Category) error
	All(ctx context.Context) ([]*models.Category, error)
	ByID(ctx context.Context, id int) (*models.Category, error)
	Update(ctx context.Context, category *models.Category) error
	Delete(ctx context.Context, id int) error
}

type GoodsRepository interface {
	Create(ctx context.Context, good *models.Good) error
	All(ctx context.Context) ([]*models.Good, error)
	ByID(ctx context.Context, id int) (*models.Good, error)
	Update(ctx context.Context, good *models.Good) error
	Delete(ctx context.Context, id int) error
}
