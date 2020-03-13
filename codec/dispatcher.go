package codec

import (
	"github.com/golang/protobuf/proto"
	"github.com/yddeng/dnet"
	"log"
)

type handler func(dnet.Session, *Message)

type Dispatcher struct {
	handlers map[string]handler
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		handlers: map[string]handler{},
	}
}

func (this *Dispatcher) Register(descriptor proto.Message, callback handler) {
	msgName := proto.MessageName(descriptor)
	if nil == callback {
		return
	}
	_, ok := this.handlers[msgName]
	if ok {
		return
	}

	this.handlers[msgName] = callback
}

func (this *Dispatcher) Dispatch(session dnet.Session, msg *Message) {
	if nil != msg {
		name := msg.GetName()
		handler, ok := this.handlers[name]
		if ok {
			handler(session, msg)
		} else {
			log.Printf("invaild name:%s", name)
		}
	}
}
