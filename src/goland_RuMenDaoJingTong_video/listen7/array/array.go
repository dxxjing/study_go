package main

import (
	"fmt"
)

//数组传指针
func modify_v2(a *[5]int){
	//*a[1] = 100  error
	//解地址必须要括起来 否则编译出错 
	(*a)[1] = 100
}

//数组传值
func modify_v1(a [5]int){
	a[3] = 100 
}

//结论：默认数组是值传递
func main(){
	arr := [5]int{1,2,3,4,5}
	modify_v1(arr)
	fmt.Println(arr) //[1,2,3,4,5]

	//传入数组地址
	modify_v2(&arr)
	fmt.Println(arr)//[1,100,3,4,5]
}
