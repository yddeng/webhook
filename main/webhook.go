package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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

	//跨域
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	if r.Method == "GET" {
		fmt.Println("method:", r.Method) //获取请求的方法

		fmt.Println("username", r.Form["username"])
		fmt.Println("password", r.Form["password"])

		for k, v := range r.Form {
			fmt.Print("key:", k, "; ")
			fmt.Println("val:", strings.Join(v, ""))
		}
	} else if r.Method == "POST" {
		result, _ := ioutil.ReadAll(r.Body)
		_ = r.Body.Close()
		fmt.Printf("%s\n", result)

		//未知类型的推荐处理方法

		var f interface{}
		_ = json.Unmarshal(result, &f)
		m := f.(map[string]interface{})
		fmt.Println(m)
		for k, v := range m {
			switch vv := v.(type) {
			case string:
				fmt.Println(k, "is string", vv)
			case int:
				fmt.Println(k, "is int", vv)
			case float64:
				fmt.Println(k, "is float64", vv)
			case []interface{}:
				fmt.Println(k, "is an array:")
				for i, u := range vv {
					fmt.Println(i, u)
				}
			default:
				fmt.Println(k, "is of a type I don't know how to handle")
			}
		}

	}
}
