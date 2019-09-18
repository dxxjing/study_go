package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 安装驱动 go get github.com/go-sql-driver/mysql
//go 已实现db的连接池

type User struct{
	Id int64	`db:"id"`
	Name string	`db:"name"`
	Age int		`db:"age"`
}

func connect()(*sql.DB,error){
	//db,err := sql.Open("mysql","root:root@/test")
	db,err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/test")

	return db,err
}

func main(){
	var db *sql.DB
	db,err := connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	db.SetMaxOpenConns(100)//设置最大连接数
	db.SetMaxIdleConns(20)//设置最大空闲连接数

	//queryData(db)
	//queryMultiData(db)
	testMultiRows(db)
}

//查询多行数据
func queryMultiData(db *sql.DB){
	sql := "select * from user where id > ?"
	rows,err := db.Query(sql,0)
	if err != nil{
		fmt.Println(err)
		return
	}
	//注意点：结果集rows 必须要释放
	//如果查询出的结果集 没有rows.Scan或者没有Scan全部结果集，结果集rows不会释放，同时该db连接也不会释放
	//会导致 设置的最大连接数被消耗完 程序卡住 问题重现见 testMultiRows函数

	defer rows.Close()//todo 重中之重
	var uList []*User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.Id,&u.Name,&u.Age)
		if err != nil {
			fmt.Println(err)
			return
		}
		uList = append(uList,&u)
	}
	fmt.Println(uList)

}

func testMultiRows(db *sql.DB){
	for i := 0; i < 120; i++ {
		fmt.Printf("%d times\n",i+1)
		sql := "select * from user where id > ?"
		rows,err := db.Query(sql,0)
		if err != nil{
			fmt.Println(err)
			return
		}
		//第一种情况 结果集没有被 rows.Scan,导致结果集rows不能被释放进而导致该db连接不能被释放
		//从而导致程序被卡死,
		//continue //模拟结果集没有被rows.Scan

		//无论第一种还是第二种 程序执行完关闭rows 连接都能得到释放
		//todo 结果集必须被rows.Scan完毕 否则rows.Close 没用
		defer rows.Close()
		continue
		for rows.Next() {
			var u User
			err := rows.Scan(&u.Id,&u.Name,&u.Age)
			if err != nil {
				fmt.Println(err)
				return
			}
			//第二种 结果集没有被全部rows.Scan 导致结果集rows不能被释放进而导致该db连接不能被释放 的情况
			continue //模拟结果集没有被全部rows.Scan
			fmt.Println(u)
		}
	}


}
//查询单行数据
func queryData(db *sql.DB){
	sql := "select * from user where id = ?"
	row := db.QueryRow(sql,1)
	//注意点：结果集row 必须要释放
	//如果查询出的结果集 没有scan或者没有scan全部，结果集row不会释放，同时该db连接也不会释放
	//会导致 设置的最大连接数被消耗完 程序卡住 问题重现见 testRow函数
	//row 没有close方法 所以必须要进行row.Scan
	var u User
	err := row.Scan(&u.Id,&u.Name,&u.Age)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u)

}

func testRow(db *sql.DB){
	sql := "select * from user where id = ?"
	//main中设置db连接池最大连接数为100
	for i := 0; i < 101; i++ {
		fmt.Printf("%d times\n",i+1)
		row := db.QueryRow(sql,1)
		continue //跳过row.Scan 会导致结果集row不能被释放进而导致该db连接不能被释放

		var u User
		err := row.Scan(&u.Id,&u.Name,&u.Age)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(u)
	}

}


