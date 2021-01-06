package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

func LikeFieldEscape(value string) string {
	value = strings.Replace(value, ";", "\\;", -1)
	value = strings.Replace(value, "\"", "\\\"", -1)
	value = strings.Replace(value, "'", "\\'", -1)
	value = strings.Replace(value, "--", "\\--", -1)
	value = strings.Replace(value, "/", "\\/", -1)
	value = strings.Replace(value, "%", "\\%", -1)
	value = strings.Replace(value, "_", "\\_", -1)
	return value
}

type Department struct {
	ID   int
	Name string
}

type Major struct {
	ID       int
	DepartID int
	Name     string
}
type Student struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"size:20"`
	Gender    int    `gorm:"size:4"`
	Number    string `gorm:"type:char(8);uniqueIndex:uk_num"`
	DepartID  uint   `gorm:"size:32;not null;index:idx_dep"`
	MajorID   uint   `gorm:"size:32;not null;index:idx_major"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {
	type Student struct {
		ID       uint   `gorm:"primarykey"`
		Name     string `gorm:"size:20"`
		DepartID uint   `gorm:"size:32;not null;index:idx_dep"`
	}

	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	db.AutoMigrate(&Student{})

	//stu1 := &Student{
	//	Name:     "std1",
	//	DepartID: 3,
	//}
	//db.Create(&stu1)
	//stu2 := &Student{
	//	Name:     "std2",
	//	DepartID: 3,
	//}
	//db.Create(&stu2)
	//
	//stu2.Name = "std3"
	//db.Debug().
	//	Select("name").
	//	Where("depart_id = ?", stu2.DepartID).
	//	Updates(stu2)

	var students []Student
	db.Debug().Select("name").Where("name LIKE ?","%" ).Find(&students)
	for _, v := range students {
		fmt.Println(v.Name)
	}
}
