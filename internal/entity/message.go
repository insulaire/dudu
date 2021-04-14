package entity

import (
	"github.com/google/uuid"
)

type IMessage interface {
	GetMessageId() uint32
	GetMessageUser() User
	GetCommand() string
	GetBody() []byte
}

type Message struct {
	Len     uint32
	User    User
	MsgId   uint32
	Body    []byte
	Command string
}

func NewMessage(user User, command string, body []byte) IMessage {
	return &Message{
		Len:     uint32(len(body)),
		MsgId:   uuid.New().ID(),
		Body:    body,
		User:    user,
		Command: command,
	}
}

func (msg *Message) GetMessageId() uint32 {
	return msg.MsgId
}

func (msg *Message) GetBody() []byte {
	return msg.Body
}

func (msg *Message) GetMessageUser() User {
	return msg.User
}

func (msg *Message) GetCommand() string {
	return msg.Command
}
