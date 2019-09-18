package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

//gzip 文件读取 读取出的内容是解压出来的

func main(){
	zipfile,err := os.OpenFile("./file.tar.gz",os.O_RDONLY,0666)
	if err != nil {
		fmt.Printf("open file err")
		return
	}
	defer zipfile.Close()
	read ,_ := gzip.NewReader(zipfile)
	bufreader := bufio.NewReader(read)
	var data []byte
	var buf [128]byte

	for{
		_,err = bufreader.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read err")
			return
		}
		data = append(data,buf[:]...)
	}

	fmt.Printf("read zip file:%s\n",string(data))
}