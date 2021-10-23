package inmemory

import (
	"context"
	"fmt"
	"lectures-6/internal/models"
	"sync"
)

type PhonesRepo struct {
	data map[int]*models.Phone

	mu *sync.RWMutex
}

func (db *PhonesRepo) Create(ctx context.Context, phone *models.Phone) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[phone.ID] = phone
	return nil
}

func (db *PhonesRepo) All(ctx context.Context) ([]*models.Phone, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	phones := make([]*models.Phone, 0, len(db.data))
	for _, phone := range db.data {
		phones = append(phones, phone)
	}

	return phones, nil
}

func (db *PhonesRepo) ByID(ctx context.Context, id int) (*models.Phone, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	phone, ok := db.data[id]
	if !ok {
		return nil, fmt.Errorf("No phone with id %d", id)
	}

	return phone, nil
}

func (db *PhonesRepo) Update(ctx context.Context, phone *models.Phone) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[phone.ID] = phone
	return nil
}

func (db *PhonesRepo) Delete(ctx context.Context, id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.data, id)
	return nil
}
