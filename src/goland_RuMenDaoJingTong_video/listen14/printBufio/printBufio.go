package main

import (
	"bufio"
	"fmt"

	//"fmt"
	"os"
)

//实现从终端读取一行数据 带缓冲  类似fmt.Scanln()
func MyScanln(){
	r := bufio.NewReader(os.Stdin)
	data,_ := r.ReadString('\n')

	//输出到终端
	os.Stdout.WriteString(string(data))
}

func Myinput(){
	var s []byte = make([]byte,10)
	_,err := os.Stdin.Read(s) //从终端读取一行 不带缓冲
	if err != nil {
		fmt.Println("os read err")
		return
	}

	os.Stdout.WriteString(string(s))
}

func main(){
	//MyScanln()
	Myinput()
}
