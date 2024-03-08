package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Новое подключение:", conn.RemoteAddr())

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Printf("Получено сообщение от %s: %s\n", conn.RemoteAddr(), message)
	}
	fmt.Printf("Соединение с %s закрыто.\n", conn.RemoteAddr())
}

func main() {
	listener, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		fmt.Println("Ошибка при прослушивании порта:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Сервер запущен. Ожидание подключений на localhost:3000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ошибка при подключении:", err)
			continue
		}

		go handleConnection(conn)
	}
}
