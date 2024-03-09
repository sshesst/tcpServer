package main

import (
	"fmt"
	"tcpServer/pkg/server"
)

func main() {
	var port string

	fmt.Print("Введите порт для запуска: ")
	fmt.Scanln(&port)
	if port == "" {
		port = "3000" // Порт по умолчанию
	}
	server.RunServer(port)
}
