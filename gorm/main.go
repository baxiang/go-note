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

type Department struct{
	ID int
	Name string
}

type Major struct {
	ID int
	DepartID int
	Name string
}
type Student struct {
	ID        uint `gorm:"primarykey"`
	Name string `gorm:"size:20"`
	Gender int `gorm:"size:4"`
	Number string `gorm:"type:char(8);uniqueIndex:uk_num"`
	DepartID uint `gorm:"size:32;not null;index:idx_dep"`
	MajorID uint `gorm:"size:32;not null;index:idx_major"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil{
		fmt.Println(err.Error())
	}
	//db.Logger.LogMode(4)
	////db.AutoMigrate(&Student{})
	//stu :=&Student{
	//	ID: 1,
	//}
    //db.Debug().First(&stu)
	//fmt.Println(stu)
	//createStu :=&Student{
	//	Name:      "秦寿生",
	//	Gender:    1,
	//	Number:    "20200003",
	//	DepartID:  3,
	//	MajorID:   1,
	//}
	//db.Debug().Create(&createStu)
	//fmt.Println(createStu.ID)

	//upd :=&Student{
	//	ID: 2,
	//	Number:    "20200005",
	//	MajorID:   1,
	//}
	db.Debug().Model(&Student{}).
		//Select("number").
		//Where("major_id = ?",1).
		Updates(map[string]interface{}{"depart_id": 4,"number": "20200009"})

	//sql :="%"
	////sql = LikeFieldEscape("%")
	////sql := fmt.Sprintf("/%%")
	//var list []user
	//err = db.Debug().Table("student").Where("name LIKE ? ", sql).Find(&list).Error
	//if err!=nil{
	//	fmt.Println(err.Error())
	//}
	//
	//fmt.Println(list)
}




