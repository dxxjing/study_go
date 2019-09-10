package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func index(w http.ResponseWriter,r *http.Request){
	_ = r
	////模拟超时 进来的请求 将有一半概率超时
	num := rand.Intn(2)
	if num == 0 {
		time.Sleep(time.Second * 5)
	}

	fmt.Fprintf(w,"srv send:%d",num)
}

func main(){
	http.HandleFunc("/",index)
	err := http.ListenAndServe("0.0.0.0:9090",nil)
	if err != nil{
		fmt.Println(err)
		return
	}
}
