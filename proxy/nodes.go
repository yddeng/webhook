package proxy

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/yddeng/dnet"
	"github.com/yddeng/webhook/codec"
	"github.com/yddeng/webhook/common"
	"github.com/yddeng/webhook/protocol"
	"time"
)

var (
	homeNodes  map[string]map[string]*Node
	nameNodes  map[string]*Node
	dispatcher *codec.Dispatcher
	tcpStarted bool
)

func tcpStart() {
	tcpStarted = true
	homeNodes = map[string]map[string]*Node{}
	nameNodes = map[string]*Node{}

	dispatcher = codec.NewDispatcher()
	dispatcher.Register(&protocol.Login{}, onLogin)
	dispatcher.Register(&protocol.Heartbeat{}, onHeartbeat)
	dispatcher.Register(&protocol.NotifyResp{}, onNotifyResp)
}

type Node struct {
	Name       string
	Homepages  []string
	Session    dnet.Session
	heartstamp time.Time
}

func (this *Node) Send(msg proto.Message) error {
	return this.Session.Send(codec.NewMessage(msg))
}

func onLogin(session dnet.Session, msg *codec.Message) {
	req := msg.GetData().(*protocol.Login)
	homepages := req.GetHomepages()
	name := req.GetName()

	if _, ok := nameNodes[name]; ok {
		_ = session.Send(codec.NewMessage(&protocol.LoginResp{
			Ok:  proto.Bool(false),
			Msg: proto.String("already Login"),
		}))
		return
	}

	client := &Node{
		Name:      name,
		Homepages: homepages,
		Session:   session,
	}
	nameNodes[name] = client

	for _, homepage := range homepages {
		clients, ok := homeNodes[homepage]
		if !ok {
			homeNodes[homepage] = map[string]*Node{}
			clients = homeNodes[homepage]
		}

		clients[name] = client
	}

	session.SetUserData(client)
	fmt.Println("onlogin", client)
}

func onHeartbeat(session dnet.Session, msg *codec.Message) {
	c := session.GetUserData()
	if c != nil {
		client := c.(*Node)
		client.heartstamp = time.Now()
		_ = client.Send(&protocol.Heartbeat{})
	}
}

func onNotifyResp(session dnet.Session, msg *codec.Message) {
	req := msg.GetData().(*protocol.NotifyResp)

	Push(&Event{
		Homepage: req.GetHomepage(),
		Cmd:      common.Message,
		Args:     []string{req.GetName(), req.GetMessage()},
	})
}

func onClose(session dnet.Session, reason string) {
	node := session.GetUserData()
	if node != nil {
		client := node.(*Node)

		delete(nameNodes, client.Name)
		for _, home := range client.Homepages {
			clients := homeNodes[home]
			delete(clients, client.Name)
		}

		fmt.Printf("client: %s closed, reason: %s", client.Name, reason)
		client.Session.SetUserData(nil)
		client.Session = nil
	}
}
