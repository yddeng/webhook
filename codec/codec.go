package codec

import (
	"fmt"
	"github.com/yddeng/dnet/util"
	"io"
)

// 编解码器
// 消息 -- 格式: 消息头(消息len＋消息cmd+消息ID), 消息体

const (
	lenSize  = 2                // 消息长度（消息体的长度）
	idSize   = 2                // 消息ID（消息体的编码ID，对应的反序列化结构）
	headSize = lenSize + idSize // 消息头长度
	buffSize = 65535            // 缓存容量(与lenSize有关，2字节最大65535）
)

type Codec struct {
	*Decoder
}

func NewCodec() *Codec {
	return &Codec{
		Decoder: &Decoder{readBuf: util.NewBuffer(buffSize)},
	}
}

type Decoder struct {
	readBuf *util.Buffer
	dataLen uint16
	msgID   uint16
}

//解码
func (decoder *Codec) Decode(reader io.Reader) (interface{}, error) {
	for {
		msg, err := decoder.unPack()

		//fmt.Println(msg, err)
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

	if decoder.dataLen == 0 {
		if decoder.readBuf.Len() < headSize {
			return nil, nil
		}

		decoder.dataLen, _ = decoder.readBuf.ReadUint16BE()
		decoder.msgID, _ = decoder.readBuf.ReadUint16BE()

	}

	if decoder.readBuf.Len() < int(decoder.dataLen) {
		return nil, nil
	}

	data, _ := decoder.readBuf.ReadBytes(int(decoder.dataLen))

	msg, err := Unmarshal(decoder.msgID, data)
	if err != nil {
		return nil, err
	}

	//将消息长度置为0，用于下一次验证
	decoder.dataLen = 0
	return msg, nil
}

//编码
func (encoder *Codec) Encode(o interface{}) ([]byte, error) {

	msgID, data, err := Marshal(o)
	if err != nil {
		return nil, err
	}

	dataLen := len(data)
	if dataLen > buffSize-headSize {
		return nil, fmt.Errorf("encode dataLen is too large,len: %d", dataLen)
	}

	msgLen := dataLen + headSize
	buff := util.NewBuffer(msgLen)

	//msgLen+cmd+msgID
	//写入data长度
	buff.WriteUint16BE(uint16(dataLen))
	//msgID
	buff.WriteUint16BE(msgID)
	//data数据
	buff.WriteBytes(data)

	//fmt.Println("encode", len(data), msgID, msg.GetSerialNo(), data, buff.Peek(), buff.Len())

	return buff.Peek(), nil
}
