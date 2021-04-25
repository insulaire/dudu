package socket

import (
	"dudu/internal/entity"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

type IConnection interface {
	Start()
	Writer(IBag) error

	Send(*entity.Message) error
	AddRoom(*entity.Message) error
	QuitRoom(*entity.Message) error
}

type Connection struct {
	//tcp连接
	conn *net.TCPConn
	//加入的服务
	server IServer
	//加入的房间 默认为nil
	room IRoom

	Reader chan entity.Message

	//Writer chan entity.Message

	closed chan struct{}
}

func NewConnection(Server IServer, conn *net.TCPConn) IConnection {
	return &Connection{
		conn:   conn,
		server: Server,
	}
}

func (c *Connection) Start() {
	defer c.QuitRoom(nil)
	pack := NewPack()
	for {
		bufHead := make([]byte, pack.GetHeaderLength())
		_, err := io.ReadFull(c.conn, bufHead)
		if err != nil {
			fmt.Println(err)
			return
		}

		bag, err := pack.UnPack(bufHead)
		if err != nil {
			log.Println(err)
			return
		}
		if bag.GetLength() > 0 {
			body := make([]byte, bag.GetLength())
			_, err = io.ReadFull(c.conn, body)
			if err != nil {
				fmt.Println(err)
				return
			}
			msg, err := pack.UnPackMessage(body)
			if err != nil {
				fmt.Println(err)
				return
			}
			if handle := c.server.GetHandler(msg.Command); handle != nil {
				if err := c.DoCommand(&msg, handle); err != nil {
					msg := entity.NewMessage([]byte(err.Error()))
					c.Writer(NewBag(msg))
				}
			} else {
				msg := entity.NewMessage([]byte(fmt.Sprintf("command [%s] not found", msg.Command)))
				c.Writer(NewBag(msg))
			}
		}
	}
}

func Reader(c *Connection) {
	ticker := time.NewTicker(time.Minute)
	for {
		select {
		case <-ticker.C:
			//c.QuitRoom(nil)
			return
		case <-c.Reader:
			ticker.Reset(time.Minute)
		case <-c.closed:
			return
		}
	}
}

func (c *Connection) DoCommand(msg *entity.Message, handle IHandler) error {
	defer handle.After(msg)
	handle.Before(msg)
	return handle.Handle(c, msg)
}

func (c *Connection) BroadcastSend(msg entity.Message) {
	newMsg := entity.NewMessage(append([]byte(fmt.Sprintf("%s:", msg.User.Name)), msg.Body...), entity.WithUser(msg.User))
	c.room.BroadcastMsg(newMsg)
}

func (c *Connection) Writer(bag IBag) error {
	pack := NewPack()
	buf, err := pack.Pack(bag)
	if err != nil {
		return err
	}
	if _, err = c.conn.Write(buf); err != nil {
		return err
	}
	return nil
}

func (c *Connection) Send(msg *entity.Message) error {
	log.Println("send :", string(msg.GetBody()))
	if c.room == nil {
		return errors.New("not in room")
	}
	c.BroadcastSend(*msg)
	return nil
}

func (c *Connection) AddRoom(msg *entity.Message) error {
	var room IRoom
	if r, ok := c.server.ExistRoom(string(msg.GetBody())); ok {
		room = r
	} else {
		room = c.server.AddRoom(string(msg.GetBody()))
	}
	if err := room.Join(msg.GetMessageUser(), c); err != nil {
		return err
	}
	c.room = room
	return nil
}

func (c *Connection) QuitRoom(msg *entity.Message) error {
	c.room.Exit(msg.GetMessageUser())
	c.room = nil
	return nil
}
