package inmemory

import (
	"lectures-6/internal/models"
	"lectures-6/internal/store"
	"sync"
)

type DB struct {
	laptopsRepo store.LaptopsRepository
	phonesRepo  store.PhonesRepository

	mu *sync.RWMutex
}

func NewDB() store.Store {
	return &DB{
		mu: new(sync.RWMutex),
	}
}

func (db *DB) Laptops() store.LaptopsRepository {
	if db.laptopsRepo == nil {
		db.laptopsRepo = &LaptopsRepo{
			data: make(map[int]*models.Laptop),
			mu:   new(sync.RWMutex),
		}
	}

	return db.laptopsRepo
}

func (db *DB) Phones() store.PhonesRepository {
	if db.phonesRepo == nil {
		db.phonesRepo = &PhonesRepo{
			data: make(map[int]*models.Phone),
			mu:   new(sync.RWMutex),
		}
	}

	return db.phonesRepo
}
