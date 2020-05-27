package orm

import (
	"database/sql"
	"github.com/baxiang/mysql-go/orm/dialect"
	"github.com/baxiang/mysql-go/orm/session"
	"log"
)

type Engine struct {
	db *sql.DB
	dialect dialect.Dialect
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

	dial, ok := dialect.GetDialect(driverName)
	if !ok{
		log.Printf("dialect %s Not Found\n", driverName)
		return
	}
	e = &Engine{db: db,dialect: dial}
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
	return session.New(engine.db,engine.dialect)
}