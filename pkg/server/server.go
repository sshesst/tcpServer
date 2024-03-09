package server

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

func RunServer(port string) {
	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		fmt.Println("Ошибка при прослушивании порта:", err)
		return
	}
	defer listener.Close()

	fmt.Printf("Сервер запущен на localhost:%s\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ошибка при подключении:", err)
			continue
		}

		go handleConnection(conn)
	}
}
