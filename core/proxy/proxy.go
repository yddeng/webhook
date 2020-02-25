package proxy

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/yddeng/dnet"
	"github.com/yddeng/dnet/socket"
	"github.com/yddeng/webhook/codec"
	"github.com/yddeng/webhook/conf"
)

func StartTcpProxy() {
	config := conf.GetConfig()
	addr := config.TcpProxyAddr
	if addr == "" {
		return
	}

	l, err := socket.NewTcpListener("tcp", addr)
	if err != nil {
		fmt.Println(1, err)
		return
	}

	gHandler := codec.NewHandler()
	gHandler.Register(&codec.Heartbeat{}, onHeartbeat)
	gHandler.Register(&codec.Login{}, onLogin)

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
					gHandler.Dispatch(session, data.(proto.Message))
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

	fmt.Println("server start on :", addr)
}
