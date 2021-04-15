package socket

import (
	"dudu/internal/entity"
)

type IHandler interface {
	Before(*entity.Message) error
	Handle(IConnection, *entity.Message) error
	After(*entity.Message) error
}

type BaseHandler struct {
}

func (h *BaseHandler) Before(*entity.Message) error {
	return nil
}

func (h *BaseHandler) Handle(IConnection, *entity.Message) error {
	return nil
}

func (h *BaseHandler) After(*entity.Message) error {
	return nil
}
