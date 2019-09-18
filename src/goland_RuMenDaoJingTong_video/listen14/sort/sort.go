package main

import (
	"fmt"
	"sort"
)
//sort 包
//该包定义了三种结构体 均已实现 Interface 接口 故这三个结构体可直接使用

//切片排序
func test1(){
	//排序
	var arr [10]int
	arr = [10]int{11,3,34,43,22,15,32,98,54,67}
	s := arr[:]
	sort.Ints(s) //只能升序
	fmt.Println("sort after:",s)//sort after: [3 11 15 22 32 34 43 54 67 98]

	//降序 sort.IntSlice 已经实现该包中的Interface 接口 所以将s 强转为sort.IntSlice(s)
	var s2 sort.IntSlice
	s2 = sort.IntSlice(s)
	sort.Sort(sort.Reverse(s2))//降序
	fmt.Println("sort reverse after:",s2)//[98 67 54 43 34 32 22 15 11 3]

}

type User struct {
	ID int
	Score int
	Name string
}
type ByIDSort []User
func (sortid ByIDSort) Len() int {
	return len(sortid)
}
func (sortid ByIDSort) Swap(i,j int) {
	sortid[i],sortid[j] = sortid[j],sortid[i]
}
func (sortid ByIDSort) Less(i,j int) bool {
	return sortid[i].ID < sortid[j].ID
}

//单字段排序
func test2(){
	var s []User
	s = []User{
		{11,34,"user11"},
		{3,23,"user3"},
		{34,11,"user34"},
		{43,56,"user43"},
		{67,44,"user67"},
	}
	fmt.Println("sort before:",s)
	sort.Sort(ByIDSort(s)) //升序
	fmt.Println("sort after:",s)
	sort.Sort(sort.Reverse(ByIDSort(s))) //降序
	fmt.Println("sort reverse after:",s)
}
//多字段排序
type UserTP []User
func (u UserTP) Len () int {
	return len(u)
}
func (u UserTP) Swap (i,j int) {
	u[i],u[j] = u[j],u[i]
}

type SortID struct {
	UserTP
}
func (sID SortID) Less(i,j int) bool {
	return sID.UserTP[i].ID < sID.UserTP[j].ID
}
type SortScore struct {
	UserTP
}
func (sScore SortScore) Less (i,j int) bool {
	return sScore.UserTP[i].Score < sScore.UserTP[j].Score
}
func test3(){
	var s []User
	s = []User{
		{3,34,"user11"},
		{3,23,"user3"},
		{34,11,"user34"},
		{43,56,"user43"},
		{67,44,"user67"},
	}
	fmt.Println("sort before:",s)
	//多字段排序 以下方式 是以id为主排序 score为副排序
	sort.Sort(SortScore{UserTP(s)})
	sort.Sort(SortID{UserTP(s)})
	fmt.Println("sort after:",s)
}

func main(){
	//test1()
	//test2()
	test3()
}
