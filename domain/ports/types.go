package ports

import (
	"context"

	"github.com/david-luk4s/desafio-dev/domain/entities"
)

type TypeTransaction interface {
	Save(context.Context, *entities.TypeTransaction) error
	GetAll(context.Context) (map[int32]*entities.TypeTransaction, error)
}

type PortTypeTransaction struct {
	Port TypeTransaction
}

func NewPortTypeTransaction(impl TypeTransaction) *PortTypeTransaction {
	return &PortTypeTransaction{Port: impl}
}

func (p *PortTypeTransaction) ServiceSave(ctx context.Context, arg *entities.TypeTransaction) error {
	return p.Port.Save(ctx, arg)
}

func (p *PortTypeTransaction) ServiceGetAll(ctx context.Context) (map[int32]*entities.TypeTransaction, error) {
	return p.Port.GetAll(ctx)
}
