package main

import (
	"fmt"
	conf "github.com/yddeng/webhook/configs/node"
	"github.com/yddeng/webhook/node"
	"os"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Printf("usage consif\n")
		return
	}

	err := conf.LoadConfig(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	config := conf.GetConfig()

	closeSignal := make(chan string)
	err = node.Start(closeSignal, config)
	if err != nil {
		fmt.Println(err)
		return
	}

	str := <-closeSignal
	fmt.Println("stop message :", str)

}
