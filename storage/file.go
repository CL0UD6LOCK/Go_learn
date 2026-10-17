package storage

import (
	"encoding/json"
	"os"
	"studentmanager/models"
)

func SaveStudents(students []models.Student) error {
	data, err := json.MarshalIndent(students, "", "\t")
	if err != nil {
		return err
	}
	f, err := os.Create("students.json.temp")
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			f.Close()
			os.Remove("students.json.temp")
		}
	}()
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	err = f.Sync()
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}
	err = os.Rename("students.json.temp", "students.json")
	return err
}

func LoadStudents(students *[]models.Student) error {
	data, err := os.ReadFile("students.json")
	if err != nil {
		if os.IsNotExist(err) {
			return SaveStudents(*students)
		}
		return err
	}
	return json.Unmarshal(data, students)
}
