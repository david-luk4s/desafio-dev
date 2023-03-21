package database

import (
	"log"
)

const type_transaction_table = `CREATE TABLE IF NOT EXISTS type_transaction (
	id_type int PRIMARY KEY,
	description varchar(255) NOT NULL,
	nature varchar(255) NOT NULL,
	signal varchar(1) NOT NULL
);`

const store_table = `CREATE TABLE IF NOT EXISTS store (
    id serial PRIMARY KEY,
    balance decimal(20,2),
    store_name varchar(19) NOT NULL UNIQUE,
    store_owner varchar(14) NOT NULL
);`

const transaction_table = `CREATE TABLE IF NOT EXISTS transactions (
    id serial PRIMARY KEY,
    id_type int NOT NULL,
    date_occurrence date NOT NULL,
    value decimal(20,2) NOT NULL,
    cpf varchar(11) NOT NULL,
    card varchar(12),
    hour_occurrence time with time zone NOT NULL,
    store_id int NOT NULL,
    FOREIGN KEY (id_type) REFERENCES type_transaction(id_type),
    FOREIGN KEY (store_id) REFERENCES store(id)
);`

func AutoMigrate() {
	tables := []string{type_transaction_table, store_table, transaction_table}

	for _, table := range tables {
		if err := migrateTable(table); err != nil {
			log.Fatal(err)
		}
	}
}

func migrateTable(table_name string) error {
	query, err := DBTest.Prepare(table_name)
	if err != nil {
		return err
	}

	query.Exec()
	return nil
}
