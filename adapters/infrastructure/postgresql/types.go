package postgresql

import (
	"context"
	"database/sql"

	"github.com/david-luk4s/desafio-dev/domain/entities"
)

const createTypeTransaction = `
INSERT INTO type_transaction(id_type, description, nature, signal) VALUES($1,$2,$3,$4);
`

type TypeTransactionImpl struct {
	sdb *sql.DB
}

func NewTypeTransactionImpl(sdb *sql.DB) *TypeTransactionImpl {
	return &TypeTransactionImpl{sdb: sdb}
}

func (t *TypeTransactionImpl) Save(ctx context.Context, arg *entities.TypeTransaction) error {
	_, err := t.sdb.ExecContext(ctx, createTypeTransaction,
		arg.IDType,
		arg.Description,
		arg.Nature,
		arg.Signal,
	)
	return err
}

const getTypeTransactions = `
SELECT id_type, description, nature, signal FROM type_transaction
`

func (t *TypeTransactionImpl) GetAll(ctx context.Context) (map[int32]*entities.TypeTransaction, error) {
	rows, err := t.sdb.QueryContext(ctx, getTypeTransactions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items = make(map[int32]*entities.TypeTransaction)

	for rows.Next() {
		var i entities.TypeTransaction
		if err := rows.Scan(
			&i.IDType,
			&i.Description,
			&i.Nature,
			&i.Signal,
		); err != nil {
			return nil, err
		}
		items[i.IDType] = &i
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
