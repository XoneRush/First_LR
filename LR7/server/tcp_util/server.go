package tcp_util

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func Start(HOST string, PORT string, cancel context.CancelFunc) {
	//старт
	var wg sync.WaitGroup

	listen, err := net.Listen("tcp", HOST+":"+PORT)
	if err != nil {
		log.Println("listening error ", err)
		cancel()
		return
	}

	log.Printf("TCP сервер успешно стартовал: %s:%s", HOST, PORT)

	//закрытие подключения
	defer listen.Close()

	//прослушивание запросов
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println("Error reading:", err)
			break
		}
		//Создание нового потока для обработки множества подключений
		wg.Add(1)
		go handleRequest(conn, wg)
	}
	wg.Wait()

}

func handleRequest(conn net.Conn, wg sync.WaitGroup) {
	wg.Done()
	for {
		buffer := make([]byte, 1024)
		_, err := conn.Read(buffer)
		if err != nil {
			log.Println(err)
			break
		}
		msg := string(buffer[:])
		if msg == "!exit" {
			break
		}

		time := time.Now().Format(time.ANSIC)
		fmt.Println("message: ", msg, " time: ", time)
		sendAnswer(conn, string("Сообщение пришло на сервер:  "+msg))

	}
	conn.Close()
}

func sendAnswer(conn net.Conn, message string) {
	conn.Write([]byte(message + "\n"))
}
