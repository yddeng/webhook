package codec

import (
	"github.com/golang/protobuf/proto"
	"github.com/yddeng/dnet"
	"reflect"
)

type handler func(dnet.Session, proto.Message)

type Handler struct {
	handlers map[string]handler
}

func NewHandler() *Handler {
	return &Handler{
		handlers: map[string]handler{},
	}
}

func (this *Handler) Register(descriptor interface{}, callback handler) {
	msgName := reflect.TypeOf(descriptor).String()
	if nil == callback {
		return
	}
	_, ok := this.handlers[msgName]
	if ok {
		return
	}

	this.handlers[msgName] = callback
}

func (this *Handler) Dispatch(session dnet.Session, msg proto.Message) {
	if nil != msg {
		name := reflect.TypeOf(msg).String()
		handler, ok := this.handlers[name]
		if ok {
			handler(session, msg)
		}
	}
}
