package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readfile(){
	file,err := os.OpenFile("./file.go",os.O_RDWR,0666)
	if err != nil {
		fmt.Printf("open file err")
		return
	}
	defer file.Close()
	var data []byte
	s := make([]byte,128)
	for{
		_,err := file.Read(s)
		if(err == io.EOF){ //读到文件结尾
			break
		}
		if(err != nil){
			fmt.Printf("read file err")
			return
		}
		data = append(data,s...)
		s = []byte{0} //必须将切片置空,或者用数组代替s切片该步骤可不用
	}

	fmt.Printf("read %d byte:%s\n",len(data),string(data))
}

func readFileWithBufio(){
	file,err := os.OpenFile("./file.go",os.O_RDWR,0666)
	if err != nil {
		fmt.Println("open file err")
		return
	}
	defer file.Close()
	bufioReader := bufio.NewReader(file)
	var data []byte //读取文件 切片必须分配空间
	data = make([]byte,100,100)
	_,err = bufioReader.Read(data)
	if err != nil {
		fmt.Println("bufio read err")
		return
	}
	fmt.Printf("bufio read:%s\n",string(data))
}

func main(){
	//readfile()
	readFileWithBufio()
}

