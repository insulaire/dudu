package socket

type ICommand interface {
	Add(key string, handle IHandler)
	Get(string) IHandler
}

type commandHandler struct {
	handles map[string]IHandler
}

func NewCommand() ICommand {
	return &commandHandler{handles: map[string]IHandler{}}
}

func (cmd *commandHandler) Add(key string, handle IHandler) {
	cmd.handles[key] = handle
}
func (cmd *commandHandler) Get(key string) IHandler {
	if v, ok := cmd.handles[key]; ok {
		return v
	}
	return nil
}
