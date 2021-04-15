package main

import (
	"dudu/internal/entity"
	"dudu/pkg/socket"
	"dudu/utils"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp4", fmt.Sprintf("%s:%d", utils.GlabalObject.Host, utils.GlabalObject.Port))
	if err != nil {
		panic(err)
	}
	go r(conn)
	name := ""
	log.Println("input name : ,")
	fmt.Scanln(&name)
	user := entity.NewUser(name)

	cmd := ""
	str := ""
	for {
		fmt.Scan(&cmd, &str)
		switch cmd {
		case "send":
			if len(str) == 0 {
				continue
			}
			msg := entity.NewMessage([]byte(str), entity.WithUser(user), entity.WithCommand("send"))
			w(msg, conn)
			break
		case "addroom":
			if len(str) == 0 {
				continue
			}
			msg := entity.NewMessage([]byte(str), entity.WithUser(user), entity.WithCommand("addroom"))
			w(msg, conn)
			break
		case "quit":
			msg := entity.NewMessage([]byte(str), entity.WithUser(user), entity.WithCommand("quit"))
			w(msg, conn)
			break
		default:
			log.Println("ERROR:command notfound")
		}

	}
	//ch := <-signal.Stop()
	//w(entity.NewMessage(user, "send", []byte("bcd")), conn)
	//select {}
}
func w(msg entity.Message, conn net.Conn) {
	pk := socket.NewPack()
	buf, err := pk.Pack(socket.NewBag(msg))
	if err != nil {
		panic(err)
	}
	// v, _ := pk.UnPackMessage(buf[8:])
	// fmt.Println(v)
	conn.Write(buf)
}

func r(conn net.Conn) {
	pack := socket.NewPack()
	for {
		bufHead := make([]byte, pack.GetHeaderLength())
		_, err := io.ReadFull(conn, bufHead)
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
			_, err = io.ReadFull(conn, body)
			if err != nil {
				fmt.Println(err)
				return
			}
			msg, err := pack.UnPackMessage(body)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(msg.GetBody()))

		}
	}
}
