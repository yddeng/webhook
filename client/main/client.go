package main

import (
	"fmt"
	"github.com/yddeng/webhook/client"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Printf("usage consif\n")
		return
	}

	client.LoadConfig(os.Args[1])
	config := client.GetConfig()

	fmt.Printf("client start on %s\n", config.NetAddr)

	http.HandleFunc("/hook/client", client.Client)
	err := http.ListenAndServe(config.NetAddr, nil)
	if err != nil {
		fmt.Println(err)
	}
}
