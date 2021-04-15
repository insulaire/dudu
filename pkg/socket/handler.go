package socket

import "dudu/internal/entity"

type IMessageHandler interface {
	DoHandler(entity.Message)
	AddHandler()
}
