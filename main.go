package main

import (
	"fmt"
	"tcp/client"
	"tcp/server"
)

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
			server.StartServer()

		case "2":
			client.StartClient()
		default:
			fmt.Println("Opcion fuera de rango, pruebe de nuevo")
		}
	}
}
