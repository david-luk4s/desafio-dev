package ports

import (
	"context"
	"io"

	"github.com/david-luk4s/desafio-dev/domain/entities"
)

type Transaction interface {
	Parse(map[int32]*entities.TypeTransaction, io.Reader) ([]entities.Transaction, error)
	Save(context.Context, *entities.Transaction) error
	GetAll(context.Context) ([]entities.Transaction, error)
}

type PortTransaction struct {
	Port Transaction
}

func NewPortTransaction(impl Transaction) *PortTransaction {
	return &PortTransaction{Port: impl}
}

func (p *PortTransaction) ServiceParse(mapTy map[int32]*entities.TypeTransaction, f io.Reader) ([]entities.Transaction, error) {
	return p.Port.Parse(mapTy, f)
}

func (p *PortTransaction) ServiceSave(ctx context.Context, ts *entities.Transaction) error {
	return p.Port.Save(ctx, ts)
}

func (p *PortTransaction) ServiceGetAll(ctx context.Context) ([]entities.Transaction, error) {
	return p.Port.GetAll(ctx)
}
