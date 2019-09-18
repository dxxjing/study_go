package main

import (
	"fmt"
	"net"
	"time"
)

func main(){
	conn,err := net.DialUDP("udp",nil,
		&net.UDPAddr{
			IP:net.IP{127,0,0,1},
			Port:8082,
		})
	if err != nil {
		fmt.Println("connect udp server err:",err)
		return
	}
	defer conn.Close()

	for{
		//写数据
		str := "nihao server!"
		_,err = conn.Write([]byte(str))
		if err != nil {
			fmt.Println("write server err:",err)
			break
		}
		var buf [4096]byte
		n,err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("read from server err:",err)
			break
		}
		fmt.Printf("read from server:%s\n",buf[:n])

		time.Sleep(time.Second *2)
	}
}
