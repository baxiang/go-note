package db

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"

)

var DB *gorm.DB

func InitGorm() error{
	var err error
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}
