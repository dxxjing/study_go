package main

import (
	"fmt"
	"os"
)

//os.Args 切片 命令行参数 第一个是程序名

func main(){
	fmt.Println(os.Args)
	for _,v := range os.Args {
		fmt.Println(v)
	}
}
