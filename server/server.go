package server

import (
	"fmt"
	"net"
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

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {

	addr := conn.RemoteAddr()
	users[addr] = User{conn: conn, name: ""}
	currentUser := users[addr]
	for {
		buffer := make([]byte, 1042)
		n, err := conn.Read(buffer)

		if err != nil {
			if err.Error() == "EOF" {
				if currentUser.name != "" {
					SendToAllUsers(addr, currentUser.name+", se ha desconectado \n")
					conn.Close()
					delete(users, addr)
				}
				break
			}
			fmt.Println(err.Error())
			continue
		}

		message := string(buffer[:n])
		if currentUser.name == "" {
			currentUser.name = message
			users[addr] = currentUser
			SendToAllUsers(addr, message+", se ha conectado")
		} else {
			fmt.Println("Mensaje Recibido: " + message)
			SendToAllUsers(addr, currentUser.name+": "+message)
		}

	}
}

func SendToAllUsers(currentAddr net.Addr, message string) {
	for key, value := range users {
		if key != currentAddr {
			_, err := value.conn.Write([]byte(message))
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}
