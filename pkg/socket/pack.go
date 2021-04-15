package socket

import (
	"bytes"
	"dudu/internal/entity"
	"encoding/binary"
	"encoding/json"
)

type IPack interface {
	Pack(IBag) ([]byte, error)
	UnPack([]byte) (IBag, error)
	GetHeaderLength() uint32
	UnPackMessage(buf []byte) (entity.Message, error)
}

type Pack struct {
	HeaderLength uint32
}

func NewPack() IPack {
	return &Pack{HeaderLength: 8}
}

func (p *Pack) GetHeaderLength() uint32 {
	return p.HeaderLength
}

func (p *Pack) Pack(msg IBag) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	b, err := json.Marshal(msg.GetBody())
	if err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, uint32(len(b))); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, msg.GetMessageId()); err != nil {
		return nil, err
	}
	buf.Write(b)
	return buf.Bytes(), nil
}

func (p *Pack) UnPack(buf []byte) (IBag, error) {
	reader := bytes.NewReader(buf)
	msg := &Bag{}
	if err := binary.Read(reader, binary.LittleEndian, &msg.len); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.LittleEndian, &msg.msgId); err != nil {
		return nil, err
	}
	return msg, nil
}

func (p *Pack) UnPackMessage(buf []byte) (entity.Message, error) {
	//reader := bytes.NewReader(buf)
	msg := entity.Message{}
	json.Unmarshal(buf, &msg)
	// if err := binary.Read(reader, binary.LittleEndian, &msg); err != nil {
	// 	return nil, err
	// }
	return msg, nil
}
