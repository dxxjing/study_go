package main

import (
	"encoding/json"
	"fmt"
)
//结构体 tag
//json encode 时 将key 转化为 tag 对应的key
type User struct {
	ID int			`json:"id"` //该字段在json中的key 为小写id
	Name string		`json:"name"`
	Addr string		`json:"-"` //该字段在json中将被忽略
}

func test(){
	u := User{
		ID:3,
		Name:"jdx",
		Addr:"上海",
	}
	fmt.Println(u) //{3 jdx 上海}

	data,_ := json.Marshal(u)
	fmt.Printf("json :%s \n",string(data))
	//json :{"id":3,"name":"jdx"}
}

func main(){
	test()
}
