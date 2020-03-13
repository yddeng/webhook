package codec

import (
	"github.com/golang/protobuf/proto"
)

type Message struct {
	name string
	data proto.Message
}

func NewMessage(data proto.Message) *Message {

	return &Message{
		name: proto.MessageName(data),
		data: data,
	}

}

func (this *Message) GetName() string {
	return this.name
}

func (this *Message) GetData() proto.Message {
	return this.data
}
