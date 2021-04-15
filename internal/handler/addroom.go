package handler

import (
	"dudu/internal/entity"
	"dudu/pkg/socket"
)

type AddRoom struct {
	socket.BaseHandler
}

func (s *AddRoom) Handle(c socket.IConnection, msg *entity.Message) error {
	return c.AddRoom(msg)
}
