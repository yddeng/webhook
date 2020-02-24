package main

import (
	"fmt"
	"github.com/yddeng/webhook/conf"
	"github.com/yddeng/webhook/core/git/gitlab"
	"github.com/yddeng/webhook/core/robot"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Printf("usage consif\n")
		return
	}

	conf.LoadConfig(os.Args[1])
	config := conf.GetConfig()

	robot.InitRobots()

	fmt.Printf("webhook start on %s\n", config.NetAddr)

	http.HandleFunc("/githook", gitlab.Hook)
	err := http.ListenAndServe(config.NetAddr, nil)
	if err != nil {
		fmt.Println(err)
	}
}
