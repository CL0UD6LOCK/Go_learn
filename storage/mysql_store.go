package storage

import "database/sql"

type MySQLStore struct {
	db *sql.DB
}

func NewMySQLStore(db *sql.DB) {

}
