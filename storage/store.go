package storage

import (
	"errors"
	"studentmanager/models"
	"sync"
)

var ErrNotFound = errors.New("student not found")

type StudentStore interface {
	List() ([]models.Student, error)
	Get(id int) (models.Student, error)
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

func (fs *FileStore) Get(id int) (models.Student, error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	// 遍历 fs.students，找到 ID 相等的学生
	for i := 0; i < len(fs.students); i++ {
		if fs.students[i].ID == id {
			return fs.students[i], nil
		}
	}
	// 循环结束还没找到，return models.Student{} 和 ErrNotFound
	return models.Student{}, ErrNotFound
}
