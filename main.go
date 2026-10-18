package main

import (
	"fmt"
	"net/http"
	"studentmanager/handlers"
	"studentmanager/storage"
)

func main() {
	fs, err := storage.NewFileStore()
	if err != nil {
		fmt.Println(err)
		return
	}
	h := &handlers.StudentHandler{Store: fs}

	http.HandleFunc("/students", h.StudentHandler)
	http.HandleFunc("/students/", h.StudentByIDHandler)
	fmt.Println("Starting server at http://localhost:8080/students (Ctrl+C to stop)")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
