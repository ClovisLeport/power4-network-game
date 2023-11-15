package main

import (
	"bufio"
	"log"
	"net"
	"time"
)

func handleClient(c net.Conn) {
	in := bufio.NewReader(c)
	out := bufio.NewWriter(c)

	msg := "ping\n"
	_, err := out.WriteString(msg)
	if err != nil {
		log.Fatal("Write error: ", err)
	}
	out.Flush()
	log.Println("Sent: ", msg)

	msg_rcv, err := in.ReadString(byte('\n'))
	if err != nil {
		log.Fatal("Read error: ", err)
	}
	log.Println("Received", msg_rcv)
}

func main() {
	listener, err := net.Listen("unix", "test.sock")
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
	defer conn.Close()
	/*log.Println("Le client s'est connecté")*/
	handleClient(conn)

	time.Sleep(10 * time.Second)

}
