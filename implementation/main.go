package implementations

import (
	"fmt"
	"tcp/client"
	"tcp/server"
)

func ServerImplementation(host, port string) {
	s := server.CreateServer(host, port)
	listener, err := s.CreateListener()
	if err != nil {
		fmt.Println(err.Error())
	}
	s.ListenForMessages(listener)
}

func ClientImplementation(name, host, port string) {
	c, err := client.CreateClient(name, host, port)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.SendName()
	go c.ReadMessages()
	go c.WriteMessages()
	select {}
}
