package main

import (
	"fmt"
)

//证明 结构体是值传递  map、slice 是引用传递
//test3 test4 说明若切片、map 含有结构体 尽量都是用指针 []*User map[string]*User
//否则会产生不必要的麻烦

type User struct {
	ID int
	Name string
}

func (u *User) Modify_1(){
	u.Name = "tom"
}

func (u User) Modify_2(){
	u.Name = "amy"
}

func test1(){
	u := User{
		ID:1,
		Name:"jdx",
	}
	u.Modify_2()
	fmt.Println(u)//jdx

	u.Modify_1()
	fmt.Println(u)//tom
}

func update(m map[string]int){
	m["jdx"] = 2
}

func test2(){
	m := make(map[string]int)
	m["jdx"] = 1
	//map 引用传递
	update(m)
	fmt.Println(m)//map[jdx:2]
}

func update2(m map[string]*User){
	m["jdx"].Name = "tom"
}

func test3(){
	//如果想要修改map 中的User 成员 map的值必须为*User
	//原因是结构体是值传递
	//
	m := make(map[string]*User)
	m["jdx"] = &User{ID:1,Name:"jdx"}
	update2(m)
	fmt.Println(m["jdx"])//&{1 tom}

}

func update4(s []User){
	s[0].Name = "tom"
}
func test4(){
	//test3 test4 表明 切片可以为[]User 而map必须为map[string]*User
	s := make([]User,10)
	s[0] = User{ID:1,Name:"jdx"}
	update4(s)
	fmt.Println(s[0])//{1 tom}
}
func main(){
	test1()
	test2()
	test3()
	test4()
}
