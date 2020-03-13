package node

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/yddeng/dnet"
	"github.com/yddeng/dnet/socket"
	"github.com/yddeng/dutil/queue"
	"github.com/yddeng/webhook/codec"
	conf "github.com/yddeng/webhook/configs/node"
	"github.com/yddeng/webhook/protocol"
	"time"
)

type Proxy struct {
	RemoteAddr    string
	Session       dnet.Session
	HeartbeatTime int64

	CloseSignal chan<- string
	config      *conf.Config
	dialing     bool
}

func (this *Proxy) Send(msg proto.Message) error {
	if this.Session == nil {
		return fmt.Errorf("Session is nil")
	}
	return this.Session.Send(codec.NewMessage(msg))
}

var (
	myProxy    *Proxy
	dispatcher = codec.NewDispatcher()
	eventQueue = queue.NewEventQueue(128, pcall)
)

func Push(msg interface{}) {
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

func Tick(now time.Time) {
	if myProxy.Session == nil {
		myProxy.dial()
	} else {
		_ = myProxy.Send(&protocol.Heartbeat{})
	}
}

func Start(closeSign chan<- string, config *conf.Config) error {

	myProxy = &Proxy{
		RemoteAddr:  config.ProxyAddr,
		CloseSignal: closeSign,
		config:      config,
	}

	eventQueue.Run(1)

	go func() {
		timer := time.NewTimer(time.Second)
		for {
			now := <-timer.C
			Push(func() { Tick(now) })
			timer.Reset(time.Second)
		}
	}()

	return nil
}

func (this *Proxy) dial() {
	if this.dialing {
		return
	}

	this.dialing = true

	go func() {
		for {
			session, err := socket.TCPDial("tcp", myProxy.RemoteAddr, time.Second*5)
			if nil == err {
				this.onConnected(session)
				return
			} else {
				fmt.Println("dial error", this.RemoteAddr, err)
				time.Sleep(1 * time.Second)
			}
		}
	}()
}

func (this *Proxy) onConnected(session dnet.Session) {
	Push(func() {
		this.dialing = false
		this.Session = session

		session.SetCodec(codec.NewCodec())
		session.SetCloseCallBack(func(reason string) {
			Push(func() {
				myProxy.Session = nil
				fmt.Printf("session closed, reason: %s\n", reason)
			})
		})

		_ = session.Start(func(data interface{}, err error) {
			if err != nil {
				session.Close(err.Error())
			} else {
				Push(func() {
					dispatcher.Dispatch(session, data.(*codec.Message))
				})
			}
		})

		this.login()

	})
}

func (this *Proxy) login() {
	login := &protocol.Login{
		Name: proto.String(this.config.Name),
	}

	list := map[string]struct{}{}
	for _, h := range this.config.Hooks {
		if _, ok := list[h.Homepage]; !ok {
			list[h.Homepage] = struct{}{}
			login.Homepages = append(login.Homepages, h.Homepage)
		}
	}

	_ = this.Send(login)
}
