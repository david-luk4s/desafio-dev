package ports

import (
	"context"

	"github.com/david-luk4s/desafio-dev/domain/entities"
)

type Store interface {
	GetOrCreate(context.Context, *entities.Store) error
	UpdateBalance(context.Context, *entities.Transaction) error
}

type PortStore struct {
	Port Store
}

func NewPortStore(impl Store) *PortStore {
	return &PortStore{Port: impl}
}

func (s *PortStore) ServiceGetOrCreate(ctx context.Context, store *entities.Store) error {
	return s.Port.GetOrCreate(ctx, store)
}

func (s *PortStore) ServiceUpdateBalance(ctx context.Context, ts *entities.Transaction) error {
	return s.Port.UpdateBalance(ctx, ts)
}
