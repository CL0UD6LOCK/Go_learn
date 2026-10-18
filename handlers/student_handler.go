package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"studentmanager/models"
	"studentmanager/storage"
	"sync"
)

type StudentHandler struct {
	Students []models.Student
	mu       sync.Mutex
	Store    storage.StudentStore
}

func (h *StudentHandler) Count() int {
	h.mu.Lock()
	temp := len(h.Students)
	h.mu.Unlock()
	return temp
}

type UpdateStudent struct {
	Name *string `json:"name"`
	Age  *int    `json:"age"`
}

func (h *StudentHandler) StudentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		students, err := h.Store.List()
		if err != nil {
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(students)
		return
	}
	if r.Method == "POST" { //POST现在只管 JSON 和 状态码
		var s models.Student
		err := json.NewDecoder(r.Body).Decode(&s) // 请求体 JSON -> 结构体
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		created, err := h.Store.Create(s)
		if err != nil {
			http.Error(w, "failed to save students", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(created) // 结构体 -> 响应 JSON
		return
	}

	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}

func (h *StudentHandler) StudentByIDHandler(w http.ResponseWriter, r *http.Request) {
	idText := strings.TrimPrefix(r.URL.Path, "/students")
	idText = strings.Trim(idText, "/")
	if idText == "" {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idText)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	if r.Method == "GET" {
		s, err := h.Store.Get(id)
		if errors.Is(err, storage.ErrNotFound) {
			http.Error(w, "student not found", http.StatusNotFound)
			return
		}
		if err != nil {
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(s)
		return
	}

	if r.Method == "PATCH" || r.Method == "PUT" {
		var updated UpdateStudent
		err := json.NewDecoder(r.Body).Decode(&updated)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		if r.Method == http.MethodPatch && updated.Name == nil && updated.Age == nil {
			http.Error(w, "no fields to update", http.StatusBadRequest)
			return
		}
		if r.Method == http.MethodPut {
			if updated.Name == nil || updated.Age == nil {
				http.Error(w, "missing fields", http.StatusBadRequest)
				return
			}
		}
		updated2, err := h.Store.Update(id, updated.Name, updated.Age)
		if errors.Is(err, storage.ErrNotFound) {
			http.Error(w, "student not found", http.StatusNotFound)
			return
		}
		if err != nil {
			http.Error(w, "failed to save students", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(updated2)
		return
	}

	if r.Method == "DELETE" {
		found := false
		var saveErr error
		h.mu.Lock()
		students := make([]models.Student, len(h.Students))
		copy(students, h.Students)
		for i := range students {
			if students[i].ID == id {
				found = true
				students = append(students[:i], students[i+1:]...)
				saveErr = storage.SaveStudents(students)
				break
			}
		}
		if found && saveErr == nil {
			h.Students = students
		}
		h.mu.Unlock()
		if !found {
			http.Error(w, "student not found", http.StatusNotFound)
			return
		}
		if saveErr != nil {
			http.Error(w, "failed to save students", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
