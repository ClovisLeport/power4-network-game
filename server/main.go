package main

import (
	"bufio"
	"log"
	"net"
	"time"
)

func handleClientRead(c net.Conn) {
	in := bufio.NewReader(c)

	msg_rcv, err := in.ReadString(byte('\n'))
	if err != nil {
		log.Fatal("Read error: ", err)
	}
	log.Println("Received", msg_rcv)
}

func handleClientWrite(c net.Conn, msg string) {

	out := bufio.NewWriter(c)

	_, err := out.WriteString(msg)
	if err != nil {
		log.Fatal("Write error: ", err)
	}
	out.Flush()
	//log.Println("Sent: ", msg)

}

func server1() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Println("listen error:", err)
		return
	}
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		log.Println("accept error:", err)
		return
	}
	log.Println("Le client1 s'est connecté")
	handleClientWrite(conn, "tu es joueur 1")

	conn2, err := listener.Accept()

	if err != nil {
		log.Println("accept error:", err)
		return
	}
	log.Println("Le client2 s'est connecté")
	handleClientWrite(conn2, "tu es joueur 2")
	time.Sleep(1 * time.Second)
	// dit aux deux joueurs que les deux sont connéctés
	handleClientWrite(conn, "2j")
	handleClientWrite(conn2, "2j")

	defer conn.Close()
	defer conn2.Close()

}

func main() {
	server1()
}
