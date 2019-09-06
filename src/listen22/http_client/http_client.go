package main

import (
	"fmt"
	"io"
	"net"
)

//简单的实现http 请求 抓取百度首页

/*
	http请求特点：
		基于tcp协议之上的文本协议
		每行文本以\r\n结尾 连续两个\r\n表示整个数据包结束
 */

func main(){
	conn,err := net.Dial("tcp","www.baidu.com:80")
	if err != nil {
		fmt.Println("dial tcp err:",err)
		return
	}

	data := "GET / HTTP/1.1\r\n"
	data += "HOST: www.baidu.com\r\n"
	data += "connection: close\r\n"
	data += "\r\n\r\n"

	n,err := io.WriteString(conn,data)
	if err != nil {
		fmt.Println("write conn err:",err)
		return
	}

	var buf [1024]byte
	n,err = conn.Read(buf[:])
	if err != nil {
		fmt.Println("conn read err:",err)
		return
	}

	fmt.Println(string(buf[:n]))


}

