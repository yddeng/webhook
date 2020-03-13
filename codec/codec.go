package codec

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/yddeng/dutil/buffer"
	"io"

	"reflect"
)

const (
	nameSize = 1                   // 协议名长度 //uint8
	bodySize = 2                   // 协议内容长度（消息体的编码ID，对应的反序列化结构）//uint16
	headSize = nameSize + bodySize // 消息头长度
	buffSize = 65535
)

type Codec struct {
	readBuf *buffer.Buffer
	name    string
	nameLen uint8
	bodyLen uint16
}

func NewCodec() *Codec {
	return &Codec{
		readBuf: buffer.NewBuffer(buffSize),
	}
}

//解码
func (decoder *Codec) Decode(reader io.Reader) (interface{}, error) {
	for {
		msg, err := decoder.unPack()

		if msg != nil {
			return msg, nil

		} else if err == nil {
			_, err1 := decoder.readBuf.ReadFrom(reader)
			if err1 != nil {
				return nil, err1
			}
		} else {
			return nil, err
		}
	}
}

func (decoder *Codec) unPack() (interface{}, error) {
	if decoder.bodyLen == 0 {
		if decoder.readBuf.Len() < headSize {
			return nil, nil
		}

		decoder.nameLen, _ = decoder.readBuf.ReadUint8BE()
		decoder.bodyLen, _ = decoder.readBuf.ReadUint16BE()
	}

	if decoder.readBuf.Len() < int(decoder.nameLen)+int(decoder.bodyLen) {
		return nil, nil
	}

	name, _ := decoder.readBuf.ReadString(int(decoder.nameLen))
	body, _ := decoder.readBuf.ReadBytes(int(decoder.bodyLen))

	//将消息长度置为0，用于下一次验证
	decoder.bodyLen = 0

	pmsg, err := Unmarshal(name, body)
	if err != nil {
		return nil, err
	}

	msg := &Message{
		name: name,
		data: pmsg.(proto.Message),
	}

	return msg, nil
}

//编码
func (encoder *Codec) Encode(o interface{}) ([]byte, error) {
	var name string
	var data []byte
	var nameLen, bodyLen int
	var err error

	msg, ok := o.(*Message)
	if !ok {
		return nil, fmt.Errorf("invailed type:%s", reflect.TypeOf(o).String())
	}

	name = msg.GetName()
	data, err = Marshal(msg.GetData())
	if err != nil {
		return nil, err
	}

	nameLen = len(name)
	bodyLen = len(data)
	if bodyLen+nameLen > buffSize-headSize {
		return nil, fmt.Errorf("encode dataLen is too large,len: %d", bodyLen+nameLen)
	}

	totalLen := headSize + nameLen + bodyLen
	buff := buffer.NewBuffer(totalLen)
	//namelen
	buff.WriteUint8BE(uint8(nameLen))
	//bodylen
	buff.WriteUint16BE(uint16(bodyLen))
	//name
	buff.WriteString(name)
	//body
	buff.WriteBytes(data)

	return buff.Peek(), nil
}

func Marshal(data interface{}) ([]byte, error) {
	ret, err := proto.Marshal(data.(proto.Message))
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func Unmarshal(name string, data []byte) (msg interface{}, err error) {
	tt := proto.MessageType(name)
	//反序列化的结构
	msg = reflect.New(tt.Elem()).Interface()
	err = proto.Unmarshal(data, msg.(proto.Message))
	if err != nil {
		return nil, err
	}
	return msg, nil
}
