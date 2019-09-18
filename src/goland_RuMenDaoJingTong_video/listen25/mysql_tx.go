package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//mysql 事务

func main(){
	db,err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println(err)
		return
	}

	testTx(db)
}

func testTx(db *sql.DB){
	tx,err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}

	sql1 := "insert into user(name,age) values(?,?)"
	sql2 := "insert into user(name,age) values(?,?)"

	result1,err := tx.Exec(sql1,"hello",23)
	result2,err := tx.Exec(sql2,"nihao",21)

	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
		err = tx.Rollback()
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	id1,_ := result1.LastInsertId()
	id2,_ := result2.LastInsertId()
	fmt.Println(id1,id2)
}

