package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"github.com/david-luk4s/desafio-dev/domain/entities"
)

type TransactionImpl struct {
	sdb *sql.DB
}

func NewTransactionImpl(sdb *sql.DB) *TransactionImpl {
	return &TransactionImpl{sdb: sdb}
}

func (t *TransactionImpl) Parse(mapTy map[int32]*entities.TypeTransaction, file io.Reader) ([]entities.Transaction, error) {
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var itens []entities.Transaction

	rows := strings.Split(string(fileBytes), "\n")

	for _, row := range rows {
		// loc, _ := time.LoadLocation("America/Sao_Paulo")
		if row == "" {
			continue
		}

		//Type Transaction
		typeT, _ := strconv.Atoi(string([]rune(row)[0:1]))

		//Date Ocurrence
		dateOccurrence, _ := time.Parse("20060102", string([]rune(row)[2-1:(2-1)+8]))

		//Value normalize (valor / 100.00)
		value, _ := strconv.Atoi(string([]rune(row)[10-1 : (10-1)+10]))
		valueF64 := float64(int64(value))
		valueF64 = valueF64 / 100.00

		// CPF of beneficiaries
		cpf := string([]rune(row)[20-1 : (20-1)+11])

		//Card used in the transaction
		card := string([]rune(row)[31-1 : (31-1)+12])

		//Time by UTC-3
		timeOnly, _ := time.Parse("150405", string([]rune(row)[43-1:(43-1)+6]))

		// Data about store
		storeOwner := string([]rune(row)[49-1 : (49-1)+14])

		// Last position subtract 1
		last := (63 - 1) + 19
		if last > len([]rune(row)) {
			last = last - 1
		}
		// Name of Store
		storeName := string([]rune(row)[63-1 : last])

		// Append all data
		itens = append(itens, entities.Transaction{
			Type:           *mapTy[int32(typeT)],
			DateOccurrence: dateOccurrence,
			Value:          valueF64,
			CPF:            cpf,
			Card:           card,
			HourOccurrence: timeOnly,
			Store:          entities.Store{StoreName: storeName, StoreOwner: storeOwner},
		})

	}
	return itens, nil
}

const createTransaction = `
INSERT INTO transactions(id_type, date_occurrence, value, cpf, card, hour_occurrence, store_id)
VALUES($1,$2,$3,$4,$5,$6,$7);
`

func (t *TransactionImpl) Save(ctx context.Context, ts *entities.Transaction) error {
	_, err := t.sdb.ExecContext(ctx, createTransaction,
		ts.Type.IDType,
		ts.DateOccurrence,
		ts.Value,
		ts.CPF,
		ts.Card,
		ts.HourOccurrence,
		ts.Store.ID,
	)
	return err
}

const getTransactions = `
SELECT 
		t.id,tt.id_type,tt.description, tt.nature, t.date_occurrence, 
		t.value, t.cpf, t.card, t.hour_occurrence, 
		s.id,s.balance,s.store_name, s.store_owner
	FROM public.transactions as t 
inner join store as s on s.id = t.store_id
inner join type_transaction as tt on tt.id_type=t.id_type;
`

func (t *TransactionImpl) GetAll(ctx context.Context) ([]entities.Transaction, error) {
	rows, err := t.sdb.QueryContext(ctx, getTransactions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []entities.Transaction

	for rows.Next() {
		var i entities.Transaction
		if err := rows.Scan(
			&i.ID,
			&i.Type.IDType,
			&i.Type.Description,
			&i.Type.Nature,
			&i.DateOccurrence,
			&i.Value,
			&i.CPF,
			&i.Card,
			&i.HourOccurrence,
			&i.Store.ID,
			&i.Store.Balance,
			&i.Store.StoreName,
			&i.Store.StoreOwner,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
