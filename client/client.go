package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var PORT = "6969"
var HOST = "localhost"

func getName() string {
	var name string
	fmt.Println("Ingrese su nombre")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		name = scanner.Text()
	}
	return name
}

func StartClient() {

	var name = getName()
	conn, err := net.Dial("tcp", HOST+":"+PORT)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = conn.Write([]byte(name))
	if err != nil {
		fmt.Print("Error leyendo su mensaje, intente de nuevo:")
	}
	go writeMessage(conn)
	go readMessage(conn)
	select {}
}
func readMessage(conn net.Conn) {
	for {
		buffer := make([]byte, 1024)
		ln, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(err.Error())
			if err.Error() == "EOF" {
				conn.Close()
				fmt.Println("Conexion terminada")
				break
			}
		}
		fmt.Println(string(buffer[:ln]))
	}
}
func writeMessage(conn net.Conn) {

	var message string
	for {
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			message = scanner.Text()
		}

		if message == "close" {
			conn.Close()
			fmt.Println("Conexion cerrada")
			os.Exit(1)
			break
		}
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Print("Error leyendo su mensaje, intente de nuevo:")
			continue
		}
	}

}
