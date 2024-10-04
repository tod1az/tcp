package main

import (
	"bufio"
	"fmt"
	"os"
	implementations "tcp/implementation"
)

var PORT = "6969"
var HOST = "localhost"

func main() {
	fmt.Println("1. Start TCP Server")
	fmt.Println("2. Start TCP Client")
	var option string
	for option == "" {
		_, err := fmt.Scan(&option)
		if err != nil {
			continue
		}

		switch option {
		case "1":
			implementations.ServerImplementation(HOST, PORT)
		case "2":
			name := getName()
			implementations.ClientImplementation(name, HOST, PORT)
		default:
			fmt.Println("Opcion fuera de rango, pruebe de nuevo")
		}
	}
}

func getName() string {
	var name string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese su nombre")
	if scanner.Scan() {
		name = scanner.Text()
	}
	return name
}
