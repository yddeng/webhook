package proxy

import (
	"fmt"
	"github.com/yddeng/dnet"
	"github.com/yddeng/dnet/socket"
	"github.com/yddeng/dutil/queue"
	"github.com/yddeng/webhook/codec"
	conf "github.com/yddeng/webhook/configs/proxy"
	"github.com/yddeng/webhook/proxy/robot"
)

var (
	eventQueue = queue.NewEventQueue(128, pcall)
)

func Push(msg interface{}) {
	switch msg.(type) {
	case func():
	case *Event:
	default:
		return
	}
	_ = eventQueue.Push(msg)
}

func pcall(i interface{}) {
	switch i.(type) {
	case func():
		i.(func())()
	case *Event:
		doEvent(i.(*Event))
	default:
	}
}

func InitRobot(config *conf.Config) {
	for _, r := range config.Robots {
		err := robot.MakeRobot(r.RobotType, r.Homepage, r.RobotUrl, r.NotifyCmd)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func ListenTcp(addr string) {
	if addr == "" {
		return
	}

	l, err := socket.NewTcpListener("tcp", addr)
	if err != nil {
		fmt.Println(1, err)
		return
	}

	tcpStart()

	err = l.StartService(func(session dnet.Session) {
		fmt.Println("new client", session.RemoteAddr().String())

		session.SetCodec(codec.NewCodec())
		session.SetCloseCallBack(func(reason string) {
			Push(func() {
				onClose(session, reason)
			})
		})

		errr := session.Start(func(data interface{}, err error) {
			if err != nil {
				session.Close(err.Error())
			} else {
				Push(func() {
					dispatcher.Dispatch(session, data.(*codec.Message))
				})
			}
		})
		if errr != nil {
			fmt.Println(2, err)
		}
	})
	if err != nil {
		fmt.Println(3, err)
	}

	fmt.Println("proxy start on :", addr)
}

func init() {
	eventQueue.Run(1)
}
