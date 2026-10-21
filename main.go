package main

import (
	"fmt"
	"studentmanager/handlers"
	"studentmanager/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	fs, err := storage.NewFileStore()
	if err != nil {
		fmt.Println(err)
		return
	}
	h := &handlers.StudentHandler{Store: fs}
	router := gin.Default()
	router.GET("/students", h.ListStudents)
	router.GET("/students/:id", h.GetStudent)
	router.POST("/students", h.CreateStudent)
	router.DELETE("/students/:id", h.DeleteStudent)

	fmt.Println("Starting server at http://localhost:8080/students (Ctrl+C to stop)")
	err = router.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
