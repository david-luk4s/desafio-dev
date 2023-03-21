package postgresql

import (
	"context"
	"database/sql"

	"github.com/david-luk4s/desafio-dev/domain/entities"
)

const createStore = `
INSERT INTO store(store_name, store_owner) VALUES($1,$2) RETURNING id;
`

const getStore = `
SELECT id, balance, store_name, store_owner FROM store WHERE store_name=$1;
`

type StoreImpl struct {
	sdb *sql.DB
}

func NewStoreImpl(sdb *sql.DB) *StoreImpl {
	return &StoreImpl{sdb: sdb}
}

func (s *StoreImpl) GetOrCreate(ctx context.Context, store *entities.Store) error {
	row := s.sdb.QueryRowContext(ctx, getStore, store.StoreName)

	if err := row.Scan(
		&store.ID,
		&store.Balance,
		&store.StoreName,
		&store.StoreOwner,
	); err != nil {
		err := s.sdb.QueryRow(createStore,
			store.StoreName,
			store.StoreOwner,
		).Scan(&store.ID)
		return err
	}

	return nil
}

const updateStore = `
UPDATE store SET balance=$1 WHERE id=$2;
`

func (s *StoreImpl) UpdateBalance(ctx context.Context, ts *entities.Transaction) error {
	var balance float64

	if ts.Type.Signal == "+" {
		//entrada +
		balance = ts.Store.Balance.Float64 + ts.Value
	} else {
		//saida -
		balance = ts.Store.Balance.Float64 - ts.Value
	}

	ts.Store.Balance.Scan(balance)
	_, err := s.sdb.ExecContext(ctx, updateStore, ts.Store.Balance.Float64, ts.Store.ID)

	return err
}
