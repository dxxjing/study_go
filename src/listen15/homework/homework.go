package main

import (
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
	"path/filepath"
)

//实现tree命令 
/*
|---src
|   |---listen15
|   |   |---file
|   |   |   |---file.go
|   |   |   |---gzip.go
|   |   |---homework
|   |---|listen14
|   |   |---sort
 */

func ListDir(dirPath string,deep int){
	dir,_ := ioutil.ReadDir(dirPath)
	if deep == 1 {
		fmt.Printf("!---%s\n", filepath.Base(dirPath))
	}
	// window的目录分隔符是 \
	// linux 的目录分隔符是 /
	sep := string(os.PathSeparator)
	for _,fi := range dir {
		if fi.IsDir() {
			fmt.Printf("|")
			for i := 0; i < deep; i++ {
				fmt.Printf("    |")
			}
			fmt.Printf("----%s\n", fi.Name())
			ListDir(dirPath+sep+fi.Name(), deep+1)
			continue
		}
		fmt.Printf("|")
		for i := 0; i < deep; i++ {
			fmt.Printf("    |")
		}
		fmt.Printf("----%s\n", fi.Name())
	}
}

func main(){
	app := cli.NewApp()
	app.Name = "tree"
	app.Usage = "list dir for tree"
	app.Action = func(c *cli.Context) error{
		var dirPath = "."
		if c.NArg() > 0 {
			dirPath = c.Args()[0]
		}
		ListDir(dirPath,1)
		return nil
	}

	app.Run(os.Args)
}
