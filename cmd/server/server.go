package main

import (
	"dudu/internal/handler"
	"dudu/pkg/socket"
	"dudu/utils"
	"fmt"
)

func main() {
	cmd := socket.NewCommand()
	cmd.Add("send", &handler.Send{})
	cmd.Add("addroom", &handler.AddRoom{})
	cmd.Add("quit", &handler.QuitRoom{})
	s := socket.NewServer(socket.WithCancelContext(), socket.WithCommand(cmd))
	printlog()
	s.Serve()
}

func printlog() {
	fmt.Println("Server Name:", utils.GlabalObject.Name)
	fmt.Println("Version:", utils.GlabalObject.Version)
	fmt.Printf("Listening at %s:%d \n", utils.GlabalObject.Host, utils.GlabalObject.Port)
}
