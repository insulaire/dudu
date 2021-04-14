package handler

type IHandler interface {
	Before()
	Handle()
	After()
}

type BaseHandler struct {
}

func (h *BaseHandler) Before() {

}

func (h *BaseHandler) Handle() {

}

func (h *BaseHandler) After() {

}
