package main

import (
	"fmt"
)

func empty(){
	//定义空接口 空接口可以存储任意类型
	var i interface{}

	var a int = 3
	i = a
	fmt.Printf("%T %v\n",i,i)

	var b float32 = 3.2
	i = b
	fmt.Printf("%T %v\n",i,i)
	var s string = "abc"
	i = s
	fmt.Printf("%T %v\n",i,i)
	var m map[string]int = map[string]int{
		"a":1,
		"b":2,
		"c":3,
	}
	i = m
	fmt.Printf("%T %v\n",i,i)
	type ST struct{
		name string
		id int
	}
	st := ST{
		name:"jdx",
		id:3,
	}
	i = st
	fmt.Printf("%T %v\n",i,i)
	var sl []int = []int{1,2,3}
	i = sl
	fmt.Printf("%T %v\n",i,i)
}

//类型断言 即判断该空接口存储的是何种类型的数据
func test(i interface{}){
	if s,ok := i.(int); ok {
		fmt.Println(s)
	}else if s,ok := i.(string); ok {
		fmt.Println(s)
	}else if s,ok := i.(float32); ok {
		fmt.Println(s)
	}else if s,ok := i.(map[string]int);ok {
		fmt.Println(s)
	}else if s,ok := i.([]int); ok{
		fmt.Println(s)
	}else if s,ok := i.(Dog); ok {//Dog结构体
		fmt.Println(s)
	}else{
		fmt.Println("unknow type")
	}
}

//只能用于switch 中的i.(type)的类型断言
func test2(i interface{}){
	switch i.(type) {
	case int:
		fmt.Println("int",i.(int))
	case string:
		fmt.Println("string",i.(string))
	case float32:
		fmt.Println("float32",i.(float32))
	case map[string]int :
		fmt.Println("map[string]int",i.(map[string]int))
	case []int:
		fmt.Println("[]int",i.([]int))
	case Dog:
		fmt.Println("Dog",i.(Dog))//结构体
	default:
		fmt.Println("unknow type")
	}
}

//test2 的变种
func test3(i interface{}){
	switch v := i.(type) {
	case int:
		fmt.Println("int",v)
	case string:
		fmt.Println("string",v)
	case float32:
		fmt.Println("float32",v)
	case map[string]int :
		fmt.Println("map[string]int",v)
	case []int:
		fmt.Println("[]int",v)
	case Dog:
		fmt.Println("Dog",v)//结构体
	default:
		fmt.Println("unknow type")
	}
}

func main(){
	//empty()
	var i float32 = 3
	test(i)
	test2(i)
	test3(i)
}
