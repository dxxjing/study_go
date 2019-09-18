package main

import (
	"fmt"
)

func createSlice(){
	var s []int
	s = make([]int,5,10) //长度 5 容量 10 超过此容量会自动扩容
	s[0] = 1
	fmt.Println(s) //[1,0,0,0,0]
}

//切片是引用传递
func modifySlice(s []int){
	s[0] = 100
}

func main(){
	//createSlice()

	s := []int{1,2,3,4}
	modifySlice(s)
	fmt.Println(s) 
}
