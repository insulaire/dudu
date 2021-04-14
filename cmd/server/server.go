package main

import (
	"dudu/pkg/socket"
	"dudu/utils"
	"fmt"
)

func main() {
	s := socket.NewServer()
	printlog()
	s.Serve()
}

func printlog() {
	fmt.Println("Server Name:", utils.GlabalObject.Name)
	fmt.Println("Version:", utils.GlabalObject.Version)
	fmt.Printf("Listening at %s:%d \n", utils.GlabalObject.Host, utils.GlabalObject.Port)
}
