package handler

import (
	"dudu/internal/entity"
	"dudu/pkg/socket"
)

type Send struct {
	socket.BaseHandler
}

func (s *Send) Handle(c socket.IConnection, msg *entity.Message) error {
	return c.Send(msg)
}
