package entities

import "database/sql"

type Store struct {
	ID         sql.NullInt64
	Balance    sql.NullFloat64
	StoreName  string
	StoreOwner string
}
