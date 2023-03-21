package entities

import "time"

type Transaction struct {
	ID             int64
	Type           TypeTransaction
	DateOccurrence time.Time
	Value          float64
	CPF            string
	Card           string
	HourOccurrence time.Time
	Store          Store
}

type TypeTransaction struct {
	IDType      int32
	Description string
	Nature      string
	Signal      string
}
