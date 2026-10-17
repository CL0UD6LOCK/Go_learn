package storage

import (
	"errors"
	"studentmanager/models"
	"sync"
)

var ErrNotFound = errors.New("student not found")

type StudentStore interface {
	List() ([]models.Student, error)
}

type FileStore struct {
	students []models.Student
	mu       sync.Mutex
}

func NewFileStore() (*FileStore, error) {
	fs := &FileStore{}
	err := LoadStudents(&fs.students)
	if err != nil {
		return nil, err
	}
	return fs, nil
}

func (fs *FileStore) List() ([]models.Student, error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	students := make([]models.Student, len(fs.students))
	copy(students, fs.students)
	return students, nil
}
