package main

import (
	"github.com/baxiang/mysql-go/orm"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	engine, err := orm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	s := engine.NewSession()

	rows, err := s.Raw("SELECT name FROM student LIMIT 10").QueryRows()
    if err!= nil{
   	   log.Fatal(err)
		return
    }
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err == nil {
			log.Println(name)
		} else {
			log.Fatal(err)
		}
	}
}
