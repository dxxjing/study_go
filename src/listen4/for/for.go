package main

import (
	"fmt"
)

func testFor(){
	//复杂
	for i,j := 10,1; i < 19 && j < 6; i,j = i+2,j+1 {
		fmt.Printf("i = %d, j = %d \n",i,j)
	}
}

//九九乘法表
func test(){

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d , ",j,i,i*j)
		}
		fmt.Printf("\n")
	}
}

func main(){
	//testFor()
	/*
	//无线循环
	for {
		...
	}
	*/
	test()
}