package inmemory

import (
	"context"
	"fmt"
	"lectures-6/internal/models"
	"lectures-6/internal/store"
	"sync"
)

type DB struct {
	data map[int]*models.Laptop

	mu *sync.RWMutex
}

func NewDB() store.Store {
	return &DB{
		data: make(map[int]*models.Laptop),
		mu:   new(sync.RWMutex),
	}
}

func (db *DB) Create(ctx context.Context, laptop *models.Laptop) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[laptop.ID] = laptop
	return nil
}

func (db *DB) All(ctx context.Context) ([]*models.Laptop, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	laptops := make([]*models.Laptop, 0, len(db.data))
	for _, laptop := range db.data {
		laptops = append(laptops, laptop)
	}

	return laptops, nil
}

func (db *DB) ByID(ctx context.Context, id int) (*models.Laptop, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	laptop, ok := db.data[id]
	if !ok {
		return nil, fmt.Errorf("No laptop with id %d", id)
	}

	return laptop, nil
}

func (db *DB) Update(ctx context.Context, laptop *models.Laptop) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[laptop.ID] = laptop
	return nil
}

func (db *DB) Delete(ctx context.Context, id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.data, id)
	return nil
}
