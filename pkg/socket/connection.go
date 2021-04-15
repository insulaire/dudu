package socket

import (
	"dudu/internal/entity"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

type IConnection interface {
	Start()
	Writer(IBag) error
	// Send()
	// AddRoom(IRoom) error
	// QuitRoom()
}

type Connection struct {
	//tcp连接
	conn *net.TCPConn
	//加入的服务
	server IServer
	//加入的房间 默认为nil
	room IRoom
	// //处理消息
	handlerFunc IMessageHandler

	Reader chan entity.Message
}

func NewConnection(Server IServer, conn *net.TCPConn) IConnection {
	return &Connection{
		conn:   conn,
		server: Server,
	}
}

func (c *Connection) Start() {
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
			if err := c.Command(msg); err != nil {
				msg := entity.NewMessage([]byte(err.Error()))
				c.Writer(NewBag(msg))
			}
		}
	}
}

func (c *Connection) Command(msg entity.Message) error {
	switch msg.GetCommand() {
	case "addroom":
		if room, ok := c.server.ExistRoom(string(msg.GetBody())); ok {
			c.room = room
		} else {
			c.room = c.server.AddRoom(string(msg.GetBody()))
		}
		return c.room.Join(msg.GetMessageUser(), c)
	case "quit":
		c.room.Exit(msg.GetMessageUser())
		c.room = nil
		return errors.New("quit succ")
	case "send":
		log.Println("send :", string(msg.GetBody()))
		if c.room == nil {
			return errors.New("not in room")
		}
		c.Send(msg)
		break
	default:
		return errors.New("command not found")
	}
	return nil
}

func (c *Connection) Send(msg entity.Message) {
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
