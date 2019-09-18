package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main(){
	conn,err := net.Dial("tcp","127.0.0.1:8080")
	if err != nil {
		fmt.Println("client connect server err:",err)
		return
	}
	defer conn.Close()

	inputRead := bufio.NewReader(os.Stdin)
	for{
		input,_ := inputRead.ReadString('\n')
		input = strings.Trim(input,"\r\n")
		if input == "Q" {
			return
		}
		_,err := conn.Write([]byte(input))
		if err != nil {
			fmt.Println("client write err:",err)
			return
		}
	}
}
