package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"studentmanager/models"
	"studentmanager/storage"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql" //只注册驱动
)

//type NullInt64 struct {
//	Int64 int64 //真正的数字
//	Valid bool  //true有值，false代表NULL
//}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, using system environment")
	}
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		log.Fatal("MYSQL_DSN environment variable is not set")
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	store := storage.NewMySQLStore(db)
	created, err := store.Create(models.Student{Name: "createtest", Age: 20})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("created:", created)

	rows, err := db.Query("SELECT id,name,age FROM students")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age sql.NullInt64
		if err := rows.Scan(&id, &name, &age); err != nil {
			log.Fatal(err)
		}
		if age.Valid {
			fmt.Println(id, name, age.Int64)
		} else {
			fmt.Println(id, name, "NULL")
		}
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
