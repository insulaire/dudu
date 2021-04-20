package entity

import (
	"github.com/google/uuid"
)

type Message struct {
	Len     uint32
	User    User
	MsgId   uint32
	Body    []byte
	Command string
}

var defaultMessageOptions = Message{
	MsgId:   uuid.New().ID(),
	User:    User{},
	Command: "send",
}

type MessageOption func(*Message)

func WithUser(user User) MessageOption {
	return func(msg *Message) {
		msg.User = user
	}
}

func WithCommand(command string) MessageOption {
	return func(msg *Message) {
		msg.Command = command
	}
}

func WithMsgId(msg_id uint32) MessageOption {
	return func(msg *Message) {
		msg.MsgId = msg_id
	}
}

func NewMessage(body []byte, opts ...MessageOption) Message {
	msg := defaultMessageOptions
	for _, v := range opts {
		v(&msg)
	}
	msg.Body = body
	msg.Len = uint32(len(body))
	return msg
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
