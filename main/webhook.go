package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Printf("usage addr\n")
		return
	}

	addr := os.Args[1]
	fmt.Printf("webhook start on %s\n", addr)

	http.HandleFunc("/githook", githook)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func githook(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	fmt.Println("githook ", r.Method, r.Form)

	//跨域
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

}
