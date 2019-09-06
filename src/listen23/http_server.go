package main

import (
	"fmt"
	"net/http"
)

func showIndex(w http.ResponseWriter, r *http.Request){
	r.ParseForm() //必须调用 才能取出Form中的get、post 参数
	fmt.Println(r.URL.Scheme)
	fmt.Println(r.Form)

	fmt.Fprintf(w,"nihao")
}

func main(){
	http.HandleFunc("/",showIndex)
	err := http.ListenAndServe("0.0.0.0:8081",nil)
	if err != nil {
		fmt.Println("http listen err:",err)
	}
}


