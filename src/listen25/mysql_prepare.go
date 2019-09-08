package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Users struct{
	Id int64	`db:"id"`
	Name string	`db:"name"`
	Age int		`db:"age"`
}

func main(){
	var db *sql.DB
	db,err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println(err)
		return
	}

	db.SetMaxOpenConns(100)

	//prepareExec(db)
	//prepareQuery(db)
	testStmt(db)
}

func prepareQuery(db *sql.DB){
	sql := "select * from user where id = ?"
	stmt,err := db.Prepare(sql)
	if err != nil {
		fmt.Println(err)
		return
	}
	//todo 重点 stmt 和rows row 一样 必须要关闭
	defer stmt.Close()

	rows,_ := stmt.Query(1)
	defer rows.Close()

	for rows.Next() {
		var u Users
		_ = rows.Scan(&u.Id,&u.Name,&u.Age)
		fmt.Println(u)
	}

	rows,_ = stmt.Query(2)
	defer rows.Close()

	for rows.Next() {
		var u Users
		_ = rows.Scan(&u.Id,&u.Name,&u.Age)
		fmt.Println(u)
	}

}

func testStmt(db *sql.DB){
	for i := 0; i< 102; i++ {
		fmt.Printf("%d times\n",i+1)
		sql := "select * from user where id = ?"
		stmt,err := db.Prepare(sql)
		if err != nil {
			fmt.Println(err)
			return
		}
		//todo 重点 stmt 和rows row 一样 必须要关闭
		defer stmt.Close()

		rows,_ := stmt.Query(1)
		defer rows.Close()

		for rows.Next() {
			var u Users
			_ = rows.Scan(&u.Id,&u.Name,&u.Age)
			fmt.Println(u)
		}
	}
}

func prepareExec(db *sql.DB){
	sql := "insert into user(name,age) values(?,?)"
	stmt,err := db.Prepare(sql)
	if err != nil {
		fmt.Println(err)
		return
	}

	result,_ := stmt.Exec("herry",40)
	fmt.Println(result.RowsAffected())

	result,_ = stmt.Exec("lily",50)
	fmt.Println(result.RowsAffected())

}
