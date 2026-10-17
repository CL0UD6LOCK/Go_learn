package main

import (
	"fmt"
	"net/http"
	"studentmanager/handlers"
	"studentmanager/models"
	"studentmanager/storage"
)

var students = []models.Student{
	{ID: 1, Name: "Jack", Age: 20},
	{ID: 2, Name: "Tom", Age: 19},
}

func main() {
	err := storage.LoadStudents(&students)
	if err != nil {
		fmt.Println(err)
		return
	}
	fs, err := storage.NewFileStore()
	if err != nil {
		fmt.Println(err)
		return
	}
	h := &handlers.StudentHandler{
		Students: students,
		Store:    fs,
	}

	http.HandleFunc("/students", h.StudentHandler)
	http.HandleFunc("/students/", h.StudentByIDHandler)
	fmt.Println("Starting server at http://localhost:8080/students (Ctrl+C to stop)")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
