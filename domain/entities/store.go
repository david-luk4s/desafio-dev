package entities

import "database/sql"

type Store struct {
	ID         int64
	Balance    sql.NullFloat64
	StoreName  string
	StoreOwner string
}
