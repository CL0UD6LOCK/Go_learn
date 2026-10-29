package storage

import (
	"database/sql"
	"studentmanager/models"
)

type MySQLStore struct {
	db *sql.DB
}

func NewMySQLStore(db *sql.DB) *MySQLStore {
	return &MySQLStore{db: db}
}

func (s *MySQLStore) List() ([]models.Student, error) {
	rows, err := s.db.Query("SELECT id,name,age FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var st models.Student
		var age sql.NullInt64
		if err := rows.Scan(&st.ID, &st.Name, &age); err != nil {
			return nil, err
		}
		if age.Valid {
			st.Age = int(age.Int64)
		} else {
			st.Age = 0
		}
		students = append(students, st)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return students, nil
}
