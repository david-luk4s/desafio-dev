package application

import (
	"context"
	"fmt"
	"io"

	adapter "github.com/david-luk4s/desafio-dev/adapters/infrastructure/postgresql"
	"github.com/david-luk4s/desafio-dev/config/database"
	"github.com/david-luk4s/desafio-dev/domain/entities"
	ports "github.com/david-luk4s/desafio-dev/domain/ports"
)

func ProcessTransaction(ctx context.Context, f io.Reader) error {
	//Get all Types Transactions
	portP := ports.NewPortTypeTransaction(
		adapter.NewTypeTransactionImpl(database.DB),
	)
	mapTs, err := portP.ServiceGetAll(ctx)
	if err != nil {
		fmt.Println("not found types transactions")
		return err
	}

	// Parse File
	portTs := ports.NewPortTransaction(
		adapter.NewTransactionImpl(database.DB),
	)
	items, err := portTs.ServiceParse(mapTs, f)
	if err != nil {
		fmt.Println(err)
		return err
	}

	portStore := ports.NewPortStore(
		adapter.NewStoreImpl(database.DB),
	)

	for _, item := range items {
		if err := portStore.ServiceGetOrCreate(ctx, &item.Store); err != nil {
			fmt.Println(err)
			continue
		}

		if err := portStore.ServiceUpdateBalance(ctx, &item); err != nil {
			fmt.Println(err)
		}

		if err := portTs.ServiceSave(ctx, &item); err != nil {
			fmt.Println(err)
		}
	}
	return nil
}

func ListTransaction(ctx context.Context) ([]entities.Transaction, error) {
	return ports.NewPortTransaction(
		adapter.NewTransactionImpl(database.DB),
	).ServiceGetAll(ctx)
}
