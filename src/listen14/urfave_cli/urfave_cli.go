package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

//命令行框架 urfave cli
//安装 go get github.com/urfave/cli  详细用法 参见文档
//todo 待完成

func main(){
	err := cli.NewApp().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
