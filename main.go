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
	//http.HandleFunc("/students", h.StudentHandler)
	//http.HandleFunc("/students/", h.StudentByIDHandler)
	fmt.Println("Starting server at http://localhost:8080/students (Ctrl+C to stop)")
	err = router.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
