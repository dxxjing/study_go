package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	var db *sql.DB
	db,err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println(err)
		return
	}

	insertData(db)
}
//更新、删除、插入 用法相同， 可使用result.RowsAffected()获取影响的行数
func insertData(db *sql.DB){
	sql := "insert into user(name,age) values(?,?)"
	result,err := db.Exec(sql,"ali",29)
	if err != nil {
		fmt.Println(err)
		return
	}
	lastId,err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("insert data success,id : ",lastId)
}



