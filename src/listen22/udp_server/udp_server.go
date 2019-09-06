package main

import (
	"fmt"
	"net"
)

func main(){
	conn,err := net.ListenUDP("udp",
		&net.UDPAddr{
			IP:net.IP{0,0,0,0},
			Port:8082,
		})
	if err != nil {
		fmt.Println("listen udp err:",err)
		return
	}
	defer conn.Close()

	for{
		//读取数据
		var buf [4096]byte
		n,remoteAddr,err := conn.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println("udp read err:",err)
			break
		}
		fmt.Println(n,remoteAddr.IP,remoteAddr.Port)
		fmt.Printf("%s \n\n",string(buf[:n]))

		//发送数据
		sendData := []byte("hello client")
		_,err =conn.WriteToUDP(sendData,remoteAddr)
		if err != nil {
			fmt.Println("udp send data err:",err)
			break
		}
	}
}
