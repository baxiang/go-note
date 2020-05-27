package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Student struct {
	gorm.Model
	Name string
	Score uint
}

func main() {
	// 注意 需要提前创建数据库create DATABASE student  mysql的docker是 mysql
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/student?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Student{})

	// create
	db.Create(&Student{Name: "tony", Score: 90})

	// read
	var s Student
	db.First(&s, 1)

	// update
	db.Model(&s).Update("score", 100)
	select {}
}
