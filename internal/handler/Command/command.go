package command

import (
	"dudu/internal/handler"
)

type Command struct {
	handles map[string]handler.IHandler
}

func (cmd *Command) AddCommand(key string, handle handler.IHandler) {
	cmd.handles[key] = handle
}

func (cmd *Command) DoCommand(key string, value string) {
	if v, ok := cmd.handles[key]; ok {
		v.Handle()
	}
}
