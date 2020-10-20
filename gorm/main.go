package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
    "github.com/jmoiron/sqlx"
	"strings"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)
type Rule struct {


	TenantID int `gorm:"type:bigint NOT NULL"`

	RuleClassID int `gorm:"type:bigint NOT NULL"`

	Enabled   bool   `gorm:"type:tinyint NOT NULL"`
	RuleName  string `gorm:"type:varchar(20) NOT NULL"`
	RuleDesc  string `gorm:"type:varchar(200) NOT NULL"`
	RuleScore int    `gorm:"type:bigint NOT NULL"`
}
type user struct {
	Id int64
	Username string
}
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
func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gateway?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil{
		fmt.Println(err.Error())
	}

	//sql :="%"
	sql := LikeFieldEscape("%")
	//sql := fmt.Sprintf("/%%")
	var list []user
	err = db.Debug().Table("userinfo").Where("username LIKE ? ", sql).Find(&list).Error
	if err!=nil{
		fmt.Println(err.Error())
	}

	fmt.Println(len(list))
    //for _,v :=range list{
	//	fmt.Printf("%d-%s\n",v.Id,v.Username)
	//}
  // initBeego()



   //initDB()
}




func initDB()  {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gateway?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)

	sqlStr := "select id, username from userinfo where username LIKE concat('%',?,'%');"
	var list []user
	err = db.Select(&list, sqlStr, "%")
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	for _, v := range list {
		fmt.Printf("%d-%s\n",v.Id,v.Username)
	}
}

func init() {
	//orm.RegisterDriver("mysql", orm.DRMySQL)

	err := orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/gateway?charset=utf8mb4&parseTime=True",
	)
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
}
func initBeego(){

	orm.Debug = true
	o := orm.NewOrm()

	var users []user
	num, err := o.Raw("SELECT id, username FROM userinfo WHERE username LIKE ? escape '/'", "/%").QueryRows(&users)
	if err == nil {
		fmt.Println("user nums: ", num)
	}
	//// 获取 QuerySeter 对象，user 为表名
	//qs := o.QueryTable("userinfo")
	//var list []user
	//all, err := qs.Filter("username__contains", "%").All(&list)
	//if err!=nil{
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Println(all)

}