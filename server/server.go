package server

import (
	"fmt"
	"net"
)

type User struct {
	name string
	conn net.Conn
}

type Server struct {
	Port      string
	Host      string
	Users     map[net.Addr]User
	StopChann chan struct{}
}

func CreateServer(host, port string) *Server {
	return &Server{
		Host:      host,
		Port:      port,
		Users:     make(map[net.Addr]User),
		StopChann: make(chan struct{}),
	}
}

func (s *Server) CreateListener() (net.Listener, error) {
	listener, err := net.Listen("tcp", s.Host+":"+s.Port)
	if err != nil {
		return nil, err
	}
	return listener, nil
}

func (s *Server) ListenForMessages(listener net.Listener) {
	defer listener.Close()
	fmt.Printf("Listening on port: %v\n", s.Port)
	for {
		select {
		case <-s.StopChann:
			fmt.Println("Deteniendo el servidor")
			s.SendToAllUsers("Deteniendo el servidor")
			listener.Close()
			return

		default:
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println(err.Error())
			}

			go s.HandleRequest(conn)
		}
	}
}

func (s *Server) HandleRequest(conn net.Conn) {

	addr := conn.RemoteAddr()
	s.Users[addr] = User{conn: conn, name: ""}
	currentUser := s.Users[addr]
	for {
		buffer := make([]byte, 1042)
		n, err := conn.Read(buffer)

		if err != nil {
			if err.Error() == "EOF" {
				if currentUser.name != "" {
					s.SendToOtherUsers(addr, currentUser.name+", se ha desconectado\n")
					conn.Close()
					delete(s.Users, addr)
				}
				break
			}
			fmt.Println(err.Error())
			continue
		}

		message := string(buffer[:n])
		if currentUser.name == "" {
			currentUser.name = message
			s.Users[addr] = currentUser
			s.SendToOtherUsers(addr, message+", se ha conectado\n")
			fmt.Printf("Addres: %v Name: %v connected\n", addr, currentUser.name)
		} else {
			fmt.Println("Mensaje Recibido: " + message)
			s.SendToOtherUsers(addr, currentUser.name+": "+message)
		}

	}
}

func (s *Server) SendToOtherUsers(currentAddr net.Addr, message string) {
	for key, value := range s.Users {
		if key != currentAddr {
			_, err := value.conn.Write([]byte(message))
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}

func (s *Server) SendToAllUsers(message string) {
	for _, value := range s.Users {
		_, err := value.conn.Write([]byte(message))
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func (s *Server) StopServer() {
	close(s.StopChann)
}
