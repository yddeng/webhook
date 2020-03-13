package main

import (
	"fmt"
	conf "github.com/yddeng/webhook/configs/proxy"
	"github.com/yddeng/webhook/proxy"
	"github.com/yddeng/webhook/proxy/git/gitlab"
	_ "github.com/yddeng/webhook/proxy/robot/workweixin"
	"net/http"
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

	proxy.InitRobot(config)
	proxy.ListenTcp(config.TcpAddr)

	fmt.Println("hook start on :", config.HookAddr)

	http.HandleFunc("/githook", gitlab.Hook)
	http.HandleFunc("/hook/gitlab", gitlab.Hook)
	err = http.ListenAndServe(config.HookAddr, nil)
	if err != nil {
		fmt.Println(err)
	}

	select {}
}
