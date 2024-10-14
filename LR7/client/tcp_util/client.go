package tcp_util

import (
	"bufio"
	"log"
	"net"
	"os"
)

func Start(HOST string, PORT string) {
	//Start connection
	log.Printf("TCP клиент успешно стартовал: %s:%s", HOST, PORT)

	conn, err := net.Dial("tcp", HOST+":"+PORT)
	if err != nil {
		log.Println("Dial failed", err.Error())
		os.Exit(1)
	}

	line := " "
	for {
		print("Введите сообщение: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line = scanner.Text()
		sendRequest(line, conn)
		handleAnswer(conn)
		if line == "!exit" {
			break
		}

	}
	conn.Close()
}

func sendRequest(message string, conn net.Conn) {
	_, err := conn.Write([]byte(message))
	if err != nil {
		log.Println("Write data failed ", err.Error())
		os.Exit(1)
	}
}

func handleAnswer(conn net.Conn) {
	reader := bufio.NewReader(conn)
	msg, err := reader.ReadString('\n')
	if err != nil {
		println("Read data failed:", err.Error())
	}
	print(msg)
}
