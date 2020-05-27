package session

import (
	"github.com/baxiang/mysql-go/orm"
	"log"
	"testing"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id int `orm:"PRIMARY KEY"`
	Name string
	Age int
}

func TestSession_CreateTable(t *testing.T){
	engine, err := orm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	s := engine.NewSession().Model(&User{})
    s.DropTable()
	s.CreateTable()
	if s.HasTable(){

	}
}