package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type Client struct {
	Name string
	Conn net.Conn
}

func CreateClient(name, host, port string) (*Client, error) {
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		return nil, err
	}
	return &Client{Name: name, Conn: conn}, nil
}

func (c *Client) SendName() {
	_, err := c.Conn.Write([]byte(c.Name))
	if err != nil {
		fmt.Print("Error leyendo su mensaje, intente de nuevo:")
	}
}
func (c *Client) ReadMessages() {
	for {
		buffer := make([]byte, 1024)
		ln, err := c.Conn.Read(buffer)
		if err != nil {
			fmt.Println(err.Error())
			if err.Error() == "EOF" {
				c.Conn.Close()
				fmt.Println("Conexion terminada")
				break
			}
		}
		fmt.Println(string(buffer[:ln]))
	}
}
func (c *Client) WriteMessages() {
	var message string
	for {
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			message = scanner.Text()
		}

		if message == "close" {
			c.Conn.Close()
			fmt.Println("Conexion cerrada")
			os.Exit(1)
		}
		_, err := c.Conn.Write([]byte(message))
		if err != nil {
			fmt.Print("Error leyendo su mensaje, intente de nuevo:")
			continue
		}
	}
}
