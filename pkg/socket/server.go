package socket

import (
	"context"
	"dudu/utils"
	"fmt"
	"log"
	"net"
	"sync"
)

type IServer interface {
	start()
	Stop()
	Serve()

	AddRoom(string) IRoom
	ExistRoom(string) (IRoom, bool)

	GetHandler(string) IHandler
}

type Server struct {
	Name      string
	Version   string
	Host      string
	Port      int
	msg       chan IBag
	rooms     sync.Map
	max, size uint32
	ctx       context.Context
	cancel    context.CancelFunc
	command   ICommand
	users     sync.Map
}

type ServerOption func(*Server)

func WithCommand(cmd ICommand) ServerOption {
	return func(s *Server) {
		s.command = cmd
	}
}
func WithCancelContext() ServerOption {
	ctx, cancel := context.WithCancel(context.Background())
	return func(s *Server) {
		s.ctx = ctx
		s.cancel = cancel
	}
}

func NewServer(opts ...ServerOption) IServer {
	server := &Server{
		Host:    utils.GlabalObject.Host,
		Port:    utils.GlabalObject.Port,
		Version: "tcp4",
		Name:    utils.GlabalObject.Name,
		msg:     make(chan IBag, 100),
		ctx:     context.Background(),
		command: NewCommand(),
		rooms:   sync.Map{},
		users:   sync.Map{},
	}
	for _, opt := range opts {
		opt(server)
	}
	return server
}

func (s *Server) start() {
	fmt.Println("Starting...")
	addr, err := net.ResolveTCPAddr(s.Version, fmt.Sprintf("%s:%d", s.Host, s.Port))
	if err != nil {
		log.Println(err)
		return
	}
	listener, err := net.ListenTCP(s.Version, addr)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Printf("Connection at %s \n", conn.RemoteAddr().String())
		c := NewConnection(s, conn)

		go c.Start()
	}

}
func (s *Server) Stop() {
	close(s.msg)
	s.cancel()
}
func (s *Server) Serve() {
	s.start()
	select {}
}

func (s *Server) RemoveConnection(uint32) {

}

func (s *Server) AddRoom(name string) IRoom {
	room := NewRoom(name, s, 2, s.ctx)
	s.rooms.Store(name, room)
	return room
}

func (s *Server) ExistRoom(name string) (IRoom, bool) {
	if v, ok := s.rooms.Load(name); ok {
		room, _ := v.(IRoom)
		return room, ok
	}
	return nil, false
}

func (s *Server) GetHandler(key string) IHandler {
	return s.command.Get(key)
}
