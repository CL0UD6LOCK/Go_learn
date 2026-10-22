package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"studentmanager/models"
	"studentmanager/storage"

	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	Store storage.StudentStore
}

type UpdateStudent struct {
	Name *string `json:"name"`
	Age  *int    `json:"age"`
}

func (h *StudentHandler) ListStudents(c *gin.Context) {
	students, err := h.Store.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, students)
}

func (h *StudentHandler) GetStudent(c *gin.Context) {
	idText := c.Param("id")
	id, err := strconv.Atoi(idText)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	student, err := h.Store.Get(id)
	if errors.Is(err, storage.ErrNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) CreateStudent(c *gin.Context) {
	var s models.Student
	err := c.ShouldBindJSON(&s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created, err := h.Store.Create(s)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)

}

func (h *StudentHandler) DeleteStudent(c *gin.Context) {
	idText := c.Param("id")
	id, err := strconv.Atoi(idText)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	err = h.Store.Delete(id)
	if errors.Is(err, storage.ErrNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *StudentHandler) UpdateStudent(c *gin.Context) {
	idText := c.Param("id")
	id, err := strconv.Atoi(idText)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var req UpdateStudent
	err = c.ShouldBindJSON(&req)
	if err != nil { //客户端发的json有问题
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if c.Request.Method == http.MethodPut && (req.Name == nil || req.Age == nil) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name/age required"})
		return
	}
	if c.Request.Method == http.MethodPatch && req.Name == nil && req.Age == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty request body"})
	}
	updated, err := h.Store.Update(id, req.Name, req.Age)
	if errors.Is(err, storage.ErrNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}
