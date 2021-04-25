package socket

import (
	"context"
	"dudu/internal/entity"
	"errors"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
)

type IRoom interface {
	//Strat()
	Join(entity.User, IConnection) error
	Exit(entity.User)
	BroadcastMsg(entity.Message)
}

type Room struct {
	//房间名
	name string
	//服务器
	server IServer
	//当前房间所有链接
	users sync.Map
	//广播消息、上下线提醒
	chanMsg chan entity.Message
	//
	max, size uint32
	//
	ctx context.Context
}

func NewRoom(name string, server IServer, max uint32, ctx context.Context) IRoom {
	r := &Room{
		name:    name,
		server:  server,
		users:   sync.Map{},
		chanMsg: make(chan entity.Message, 10000),
		max:     max,
		ctx:     ctx,
	}
	//启动一个广播协程
	go r.Broadcast()
	return r

}

func (r *Room) Join(user entity.User, conn IConnection) error {
	if r.size >= r.max {
		return errors.New("room full ...")
	}
	r.users.Store(user, conn)
	atomic.AddUint32(&r.size, 1)
	msg := entity.NewMessage([]byte(fmt.Sprintf("welcome! %s", user.Name)), entity.WithUser(user), entity.WithCommand("send"))
	go r.BroadcastMsg(msg)
	return nil
}

func (r *Room) Exit(user entity.User) {
	r.users.Delete(user)
	atomic.AddUint32(&r.size, ^uint32(1-1))
}

func (r *Room) BroadcastMsg(msg entity.Message) {

	r.chanMsg <- msg
}

func (r *Room) Broadcast() {
	for {
		select {
		case msg := <-r.chanMsg:
			r.users.Range(func(key, value interface{}) bool {
				log.Println("recv  :", key, string(msg.GetBody()))
				u, _ := key.(entity.User)
				conn, _ := value.(IConnection)
				go func(u entity.User, conn IConnection) {
					uu := msg.GetMessageUser()
					if u.Id != uu.Id {
						if err := conn.Writer(NewBag(msg)); err != nil {
							log.Println("broadcast send error:", err)
						}
						log.Println("broadcast send succ:", uu.Id, ":", msg.GetMessageId())
					}
				}(u, conn)

				return true
			})
		case <-r.ctx.Done():
			return
		}
	}
}

// func (r *Room) Strat() {

// }

func (r *Room) Close() {
	close(r.chanMsg)
}
