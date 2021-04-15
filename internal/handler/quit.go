package handler

import (
	"dudu/internal/entity"
	"dudu/pkg/socket"
)

type QuitRoom struct {
	socket.BaseHandler
}

func (s *QuitRoom) Handle(c socket.IConnection, msg *entity.Message) error {
	return c.QuitRoom(msg)
}
