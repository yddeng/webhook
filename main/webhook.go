package main

import (
	"encoding/json"
	"fmt"
	"github.com/yddeng/webhook"
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

	http.HandleFunc("/githook", webhook.GitHook)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println(err)
	}
}
