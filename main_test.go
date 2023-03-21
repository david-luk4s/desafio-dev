package main

import (
	"context"
	"os"
	"testing"

	adapter "github.com/david-luk4s/desafio-dev/adapters/infrastructure/postgresql"
	"github.com/david-luk4s/desafio-dev/config/database"
	"github.com/david-luk4s/desafio-dev/domain/entities"
	"github.com/david-luk4s/desafio-dev/domain/ports"
)

var (
	ITEMS_TYPE_TRANSACTION map[int32]*entities.TypeTransaction
	ITEMS_TRANSACTIONS     []entities.Transaction
)

func init() {
	database.ConnectionDBTest()
}

func TestSaveTypeTransaction(t *testing.T) {
	ctx := context.Background()
	port := ports.NewPortTypeTransaction(
		adapter.NewTypeTransactionImpl(database.DBTest),
	)

	typeTransaction := []entities.TypeTransaction{
		entities.TypeTransaction{IDType: 1, Description: "Débito", Nature: "Entrada", Signal: "+"},
		entities.TypeTransaction{IDType: 2, Description: "Boleto", Nature: "Saída", Signal: "-"},
		entities.TypeTransaction{IDType: 3, Description: "Financiamento", Nature: "Saída", Signal: "-"},
		entities.TypeTransaction{IDType: 4, Description: "Crédito", Nature: "Entrada", Signal: "+"},
		entities.TypeTransaction{IDType: 5, Description: "Recebimento Empréstimo", Nature: "Entrada", Signal: "+"},
		entities.TypeTransaction{IDType: 6, Description: "Vendas", Nature: "Entrada", Signal: "+"},
		entities.TypeTransaction{IDType: 7, Description: "Recebimento TED", Nature: "Entrada", Signal: "+"},
		entities.TypeTransaction{IDType: 8, Description: "Recebimento DOC", Nature: "Entrada", Signal: "+"},
		entities.TypeTransaction{IDType: 9, Description: "Aluguel", Nature: "Saída", Signal: "-"},
	}

	for _, item := range typeTransaction {
		if err := port.ServiceSave(ctx, &item); err != nil {
			t.Errorf("Save(ctx, TypeTransaction) FAILED. expcted nil, got %s", err)
		} else {
			t.Logf("Save(ctx, %d) PASSED. expcted nil, got nil", item.IDType)
		}
	}

}

func TestRecoveryTypeTransaction(t *testing.T) {
	var err error
	ctx := context.Background()

	port := ports.NewPortTypeTransaction(
		adapter.NewTypeTransactionImpl(database.DBTest),
	)

	ITEMS_TYPE_TRANSACTION, err = port.ServiceGetAll(ctx)
	if err != nil {
		t.Errorf("ServiceGetAll(ctx) FAILED. expcted %d items, got %d items", 9, len(ITEMS_TYPE_TRANSACTION))
		t.Errorf("ServiceGetAll(ctx) FAILED. expcted nil error, got %s error", err)
	} else {
		if len(ITEMS_TYPE_TRANSACTION) == 9 {
			t.Logf("ServiceGetAll(ctx) PASSED. expcted %d items, got %d items", 9, len(ITEMS_TYPE_TRANSACTION))
			t.Logf("ServiceGetAll(ctx) PASSED. expcted nil error, got %s error", err)
		} else {
			t.Errorf("ServiceGetAll(ctx) FAILED. expcted %d items, got %d items", 9, len(ITEMS_TYPE_TRANSACTION))
		}

	}
}

func TestParseFile(t *testing.T) {
	port := ports.NewPortTransaction(
		adapter.NewTransactionImpl(database.DBTest),
	)

	f, err := os.Open("./CNAB.txt")
	if err != nil {
		t.Error(err)
	}

	ITEMS_TRANSACTIONS, err = port.ServiceParse(ITEMS_TYPE_TRANSACTION, f)
	if err != nil {
		t.Errorf("ServiceParse(ctx) FAILED. expcted nil error, got %s error", err)
	} else {
		t.Logf("ServiceParse(ctx) PASSED. expcted nil error, got %s error", err)
		t.Logf("ServiceParse(ctx) PASSED. expcted %d items, got %d items", 21, len(ITEMS_TRANSACTIONS))
	}

}

func TestProcessTransaction(t *testing.T) {
	ctx := context.Background()
	portStore := ports.NewPortStore(
		adapter.NewStoreImpl(database.DBTest),
	)

	portTransaction := ports.NewPortTransaction(
		adapter.NewTransactionImpl(database.DBTest),
	)

	for _, item := range ITEMS_TRANSACTIONS {
		//Test Adapters Service GetOrCreate of Store
		if err := portStore.ServiceGetOrCreate(ctx, &item.Store); err != nil {
			t.Errorf("Store:ServiceGetOrCreate(ctx, &item.Store) FAILED. expcted nil error, got %s error", err)
		} else {
			t.Logf("ServiceGetOrCreate(ctx, &item.Store) PASSED. expcted nil error, got %s error", err)
		}

		//Test Adapters Service UpdateBalance of Store
		if err := portStore.ServiceUpdateBalance(ctx, &item); err != nil {
			t.Errorf("Store:ServiceUpdateBalance(ctx, &item) FAILED. expcted nil error, got %s error", err)
		} else {
			t.Logf("Store:ServiceUpdateBalance(ctx, &item) PASSED. expcted nil error, got %s error", err)
		}

		//Test Adapters Service Save of Transaction
		if err := portTransaction.ServiceSave(ctx, &item); err != nil {
			t.Errorf("Transaction:ServiceSave(ctx, &item) FAILED. expcted nil error, got %s error", err)
		} else {
			t.Logf("Transaction:ServiceSave(ctx, &item) PASSED. expcted nil error, got %s error", err)
		}
	}

}
