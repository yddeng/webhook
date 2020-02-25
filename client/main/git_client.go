package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/golang/protobuf/proto"
	"github.com/yddeng/dnet"
	"github.com/yddeng/dnet/socket"
	"github.com/yddeng/dutil/queue"
	"github.com/yddeng/webhook/codec"
	"github.com/yddeng/webhook/core/event"
	"os"
	"os/exec"
	"time"
)

type Config struct {
	ProxyAddr string `toml:"ProxyAddr"`
	Name      string `toml:"Name"`
	Homepage  string `toml:"Homepage"`
	Branch    string `toml:"Branch"`
	ShellPath string `toml:"ShellPath"`
}

func LoadConfig(path string) {
	config = &Config{}
	_, err := toml.DecodeFile(path, config)
	if err != nil {
		panic(err)
	}
	fmt.Println(config)
}

var (
	config     *Config
	gHandler   *codec.Handler
	eventQueue = queue.NewEventQueue(128, pcall)
)

func push(msg interface{}) {
	switch msg.(type) {
	case func():
	default:
		return
	}
	_ = eventQueue.Push(msg)
}

func pcall(i interface{}) {
	switch i.(type) {
	case func():
		i.(func())()
	default:
	}
}

func Command(path string) {
	cmd := exec.Command("/bin/bash", "-c", path)

	//err := cmd.Run()
	//if err != nil {
	//	fmt.Println("Execute Command failed:" + err.Error())
	//}

	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("execute shell:%s failed with error:%s \n\n", path, err.Error())
		return
	}
	fmt.Printf("execute shell: %s finished with output:\n %s \n\n ", path, string(output))
}

func onEvent(session dnet.Session, msg proto.Message) {
	req := msg.(*codec.Event)
	fmt.Println("onEvent", req)
	if req.GetHomepage() == config.Homepage && req.GetBranch() == config.Branch {
		switch req.GetMsgType() {
		case event.PushEvent:
		case event.MergeRequest:
		default:
			return
		}
		Command(config.ShellPath)
	}
}

func onHeartbeat(session dnet.Session, msg proto.Message) {

}

func login(session dnet.Session) {
	_ = session.Send(&codec.Login{
		Homepage: proto.String(config.Homepage),
		Name:     proto.String(config.Name),
	})
}

func main() {
	if len(os.Args) < 1 {
		fmt.Printf("usage consif\n")
		return
	}

	LoadConfig(os.Args[1])
	gHandler = codec.NewHandler()
	gHandler.Register(&codec.Event{}, onEvent)
	gHandler.Register(&codec.Heartbeat{}, onHeartbeat)
	eventQueue.Run(1)

	session, err := socket.TCPDial("tcp", config.ProxyAddr, 0)
	if err != nil {
		fmt.Println(1, err)
		return
	}
	fmt.Printf("conn ok,remote:%s\n", session.RemoteAddr())

	session.SetCodec(codec.NewCodec())
	session.SetCloseCallBack(func(reason string) {
		fmt.Println("onClose", reason)
	})
	err = session.Start(func(data interface{}, err2 error) {
		if err2 != nil {
			session.Close(err2.Error())
		} else {
			push(func() {
				gHandler.Dispatch(session, data.(proto.Message))
			})
		}
	})
	if err != nil {
		fmt.Println(2, err)
		return
	}

	push(func() { login(session) })

	go func() {
		timer := time.NewTimer(time.Minute)
		for {
			now := <-timer.C
			push(func() {
				_ = session.Send(&codec.Heartbeat{
					Timestamp: proto.Int64(now.UnixNano()),
				})
			})
			timer.Reset(time.Minute)

		}
	}()

	select {}

}
