package main

import (
	"fmt"
)

func switchV1(){
	num := 10
	//表达式
	switch {
	case num >= 1 && num < 5 :
		fmt.Println("small")
	case num >= 5 && num < 100:
		fmt.Println("middle")
	case num >= 100:
		fmt.Println("big")
	default:
		fmt.Println("little 1")
	}
}

func main(){
	switchV1()
}