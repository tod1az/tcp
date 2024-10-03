package server

import (
	"fmt"
	"net"
	"strings"
)

var PORT = "6969"
var HOST = "localhost"
var TYPE = "tcp"

type User struct {
	name string
	conn net.Conn
}

var users = make(map[net.Addr]User)

func StartServer() {

	listener, err := net.Listen("tcp", HOST+":"+PORT)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer listener.Close()
	fmt.Printf("Listening on port: %v \n", PORT)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}
		users[conn.RemoteAddr()] = User{conn: conn, name: ""}

		go handleRequest(conn)
	}
}
func handleRequest(conn net.Conn) {

	addr := conn.RemoteAddr()
	currentUser := users[addr]
	for {
		buffer := make([]byte, 1042)
		n, err := conn.Read(buffer)

		if err != nil {
			if err.Error() == "EOF" {
				if currentUser.name != "" {
					fmt.Printf("%v, se ha desconectado \n", currentUser.name)
				}
				break
			}
			fmt.Println(err.Error())
			continue
		}

		message := string(buffer[:n])
		if currentUser.name == "" {
			currentUser.name = strings.Split(message, ":")[0]
			users[addr] = currentUser
		}

		fmt.Println("Mensaje Recibido: " + message)
		for key, value := range users {
			if key != addr {
				_, err := value.conn.Write([]byte(message))
				if err != nil {
					fmt.Println(err.Error())
				}
			}
		}
	}
}
