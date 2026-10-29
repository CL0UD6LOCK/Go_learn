package storage

import "database/sql"

type MySQLStore struct {
	db *sql.DB
}
