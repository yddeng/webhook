package codec

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"reflect"
)

type Protocol struct {
	id2Type map[uint16]reflect.Type
	type2Id map[reflect.Type]uint16
}

var pProtocol *Protocol

func Register(id uint16, msg interface{}) {
	if pProtocol == nil {
		fmt.Errorf("protocol is nil,need init")
		return
	}

	tt := reflect.TypeOf(msg)

	if _, ok := pProtocol.id2Type[id]; ok {
		fmt.Errorf("%d already register to type:%s", id, tt)
		return
	}

	pProtocol.id2Type[id] = tt
	pProtocol.type2Id[tt] = id
}

func Marshal(data interface{}) (uint16, []byte, error) {
	if pProtocol == nil {
		return 0, nil, fmt.Errorf("protocol is nil,need init")
	}

	id, ok := pProtocol.type2Id[reflect.TypeOf(data)]
	if !ok {
		return 0, nil, fmt.Errorf("type: %s undefined", reflect.TypeOf(data))
	}

	ret, err := proto.Marshal(data.(proto.Message))
	if err != nil {
		return 0, nil, err
	}

	return id, ret, nil
}

func Unmarshal(msgID uint16, data []byte) (msg interface{}, err error) {
	if pProtocol == nil {
		return nil, fmt.Errorf("protocol is nil,need init")
	}

	tt, ok := pProtocol.id2Type[msgID]
	if !ok {
		err = fmt.Errorf("msgID: %d undefined", msgID)
		return
	}

	//反序列化的结构
	msg = reflect.New(tt.Elem()).Interface()
	err = proto.Unmarshal(data, msg.(proto.Message))
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func init() {
	pProtocol = &Protocol{
		id2Type: map[uint16]reflect.Type{},
		type2Id: map[reflect.Type]uint16{},
	}

	Register(1, &Heartbeat{})
	Register(2, &Event{})
	Register(3, &Login{})
}
