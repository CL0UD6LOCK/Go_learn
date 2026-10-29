package storage

import (
	"database/sql"
	"fmt"
	"log"
)

type MySQLStore struct {
	db *sql.DB
}

func NewMySQLStore(db *sql.DB) *MySQLStore {
	return &MySQLStore{db: db}
}

func List() {
	db, err := sql.Open("mysql", "List")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT id,name,age FROM students")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age sql.NullInt64
		if err := rows.Scan(&id, &name, &age); err != nil {
			log.Fatal(err)
		}
		if age.Valid {
			fmt.Println(id, name, age.Int64)
		} else {
			fmt.Println(id, name, "NULL")
		}
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
