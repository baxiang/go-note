package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	// 验证链接
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT name,number FROM student LIMIT 10")

	var name, num string
	for rows.Next() {
		if err := rows.Scan(&name, &num); err == nil {
			log.Println(name, num)
		} else {
			log.Fatal(err)
		}
	}
}
