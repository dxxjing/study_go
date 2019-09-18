package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main(){

	http.HandleFunc("/login",login)

	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		fmt.Println("http listen err:",err)
		return
	}
}


func login(w http.ResponseWriter,r *http.Request){
	fmt.Println(r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("./login.html")
		if err != nil {
			fmt.Fprintf(w, "err")
		}
		t.Execute(w,nil)//渲染页面到web
	}else{
		r.ParseForm()//解析表单 否则r.Form 为空
		fmt.Println(r.Form,r.Form["username"],r.Form["password"])
		fmt.Println(r.FormValue("username"),r.FormValue("password"))
	}
}

