package main

import (
	"fmt"
	"sort"
)

type User struct {
	ID int
	Score int
	Name string
}
type UserArr []*User
func (u *UserArr) Len() int {
	return len(*u)
}
func (u *UserArr) Swap(i,j int) {
	(*u)[i],(*u)[j] = (*u)[j],(*u)[i]
}
func (u *UserArr) Less(i,j int) bool {
	return (*u)[i].ID < (*u)[j].ID
}

//单字段排序
//本例采用结构体指针切片 且以指针作为接受者

func test2(){
	var s []*User
	s = []*User{
		&User{11,34,"user11"},
		&User{3,23,"user3"},
		&User{34,11,"user34"},
		&User{43,56,"user43"},
		&User{67,44,"user67"},
	}
	tmp := UserArr(s)
	fmt.Println("sort before:")
	for _, d := range s {
		fmt.Printf("%d %d %s,", d.ID,d.Score,d.Name)
	}
	fmt.Println()
	//若结构体成员方法 内部显式的使用指针 则外部必须传指针进去
	//且不能直接使用 &UserArr(s) 会提示无法获取地址
	sort.Sort(&tmp) //升序
	fmt.Println("sort after:")
	for _, d := range s {
		fmt.Printf("%d %d %s,", d.ID,d.Score,d.Name)
	}
}

func main(){
	test2()
}
