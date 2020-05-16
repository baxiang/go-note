package orm

import (
	"database/sql"
	"github.com/baxiang/mysql-go/orm/session"
	"log"
)

type Engine struct {
	db *sql.DB
}

func NewEngine(driverName , sourceName string)(e *Engine,err error){
	db, err := sql.Open(driverName, sourceName)
	if err!=nil{
		log.Fatal(err)
		return
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
		return
	}
	e = &Engine{db: db}
	log.Println("Connect database success")
	return
}

func (engine *Engine)Close(){
	if err := engine.db.Close(); err != nil {
		log.Println("Failed to close database")
	}
	log.Println("Close database success")
}

func (engine *Engine)NewSession()*session.Session{
	return session.New(engine.db)
}