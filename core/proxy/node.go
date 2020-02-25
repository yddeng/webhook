package proxy

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/yddeng/dnet"
	"github.com/yddeng/dutil/queue"
	"github.com/yddeng/webhook/codec"
)

var (
	homeNodes  map[string]map[string]*Node
	eventQueue *queue.EventQueue
)

type Node struct {
	Name     string
	Homepage string
	Session  dnet.Session
}

func onLogin(session dnet.Session, msg proto.Message) {
	req := msg.(*codec.Login)
	homepage := req.GetHomepage()
	name := req.GetName()

	nodes, ok := homeNodes[homepage]
	if !ok {
		homeNodes[homepage] = map[string]*Node{}
		nodes = homeNodes[homepage]
	}

	node, ok := nodes[name]
	if ok {
		node.Session.SetUserData(nil)
		node.Session.Close("new session login,close")
		delete(nodes, name)
	}

	node = &Node{
		Name:     name,
		Homepage: homepage,
		Session:  session,
	}

	session.SetUserData(node)
	nodes[node.Name] = node
	fmt.Println("onlogin", name, homepage)
}

func onHeartbeat(session dnet.Session, msg proto.Message) {
	req := msg.(*codec.Heartbeat)
	_ = session.Send(&codec.Heartbeat{
		Timestamp: proto.Int64(req.GetTimestamp()),
	})
}

func onClose(session dnet.Session, reason string) {

	node := session.GetUserData()
	if node != nil {
		n := node.(*Node)
		nodes := homeNodes[n.Homepage]
		if _, ok := nodes[n.Name]; ok {
			fmt.Printf("client: %s closed, reason: %s", n.Name, reason)
			delete(nodes, n.Name)
		}
	}
}

func Push(msg interface{}) {
	switch msg.(type) {
	case func():
	case *codec.Event:
	default:
		return
	}
	_ = eventQueue.Push(msg)
}

func pcall(i interface{}) {
	switch i.(type) {
	case func():
		i.(func())()
	case *codec.Event:
		msg := i.(*codec.Event)
		fmt.Println("proxy pcall *codec.Event", msg)
		nodes, ok := homeNodes[msg.GetHomepage()]
		if ok {
			for _, n := range nodes {
				_ = n.Session.Send(msg)
			}
		}
	}
}

func init() {
	eventQueue = queue.NewEventQueue(128, pcall)
	eventQueue.Run(1)

	homeNodes = map[string]map[string]*Node{}
}
