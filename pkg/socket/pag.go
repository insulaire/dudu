package socket

import (
	"dudu/internal/entity"

	"github.com/google/uuid"
)

type IBag interface {
	GetLength() uint32
	GetMessageId() uint32
	GetBody() entity.Message
}

type Bag struct {
	len   uint32
	msgId uint32
	body  entity.Message
}

func NewBag(body entity.Message) IBag {
	return &Bag{
		msgId: uuid.New().ID(),
		body:  body,
	}
}

func (msg *Bag) GetLength() uint32 {
	return msg.len
}

func (msg *Bag) GetMessageId() uint32 {
	return msg.msgId
}

func (msg *Bag) GetBody() entity.Message {
	return msg.body
}
