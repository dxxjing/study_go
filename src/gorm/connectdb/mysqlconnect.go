package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct{
	//gorm.Model
	ID int	`gorm:"column:id"` //设置列明
	Name string	`gorm:"column:name"`
	Age int	`gorm:"column:age"`
}

func (User) TableName() string{
	return "user"
}

func main() {
	db,err := gorm.Open("mysql","root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//db.AutoMigrate(&User{}) //创建表

	/*user := User{
		ID:110,
		Name:"ppp",
		Age:999,
	}
	db.Create(&user)*/
	//var user User // 只取一条
	var user []User //切片 取多条
	//db.LogMode(true).Find(&user) //LogMode(true) 查询日志
	db.Find(&user)
	fmt.Println(user)
}
