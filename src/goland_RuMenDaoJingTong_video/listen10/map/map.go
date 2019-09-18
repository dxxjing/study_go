package main

import (
	"fmt"
)

func createMap(){
	var m map[string]int
	//使用前 必须用make 初始化
	m = make(map[string]int,10) //len 可不指定 最好预估下 否则会动态分配
	m["a"] = 1
	fmt.Printf("map : %v \n",m)
}

func createMap2(){
	//map第二种创建方法
	var m map[string]int
	m = map[string]int{"a":1,"b":2,"c":3}
	fmt.Println(m)

	//取值 key存在 ok 为true  否则 为false
	v,ok := m["d"]
	if ok {
		fmt.Printf("m[\"a\"] = %d \n",v)
	}else{
		fmt.Println("the key not exist")
	}

}



func main(){
	//createMap()
	createMap2()
}
