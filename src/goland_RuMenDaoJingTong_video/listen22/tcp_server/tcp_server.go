package main

import (
	"fmt"
	"net"
)

func main(){
	//监听端口
	listen,err := net.Listen("tcp","0.0.0.0:8080")
	if err != nil {
		fmt.Println("listen port err:",err)
		return
	}
	fmt.Println("tcp server listen...")
	for{
		//接收请求
		conn,err := listen.Accept()
		if err != nil {
			fmt.Println("accept err:",err)
			return
		}
		//开启协程处理请求
		go process(conn)
	}
}

func process(conn net.Conn){
	defer conn.Close()

	for {
		var b [512]byte
		n,err := conn.Read(b[:])
		//若 err == io.EOF 读完客户端数据 此时即可退出该协程
		if err != nil {
			fmt.Println("read err:",err)
			break
		}
		//底层按字节传输
		newb := b[:n]
		fmt.Printf("recv from client data:%s\n",string(newb))
		//str := "server:"+string(newb)
		//_, err = conn.Write([]byte(str))

	}

}
