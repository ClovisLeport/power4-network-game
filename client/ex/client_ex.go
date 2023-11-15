package main

import (
	"log"
	"net"
)

func handleConnection(c net.Conn) {
	in_buf := make([]byte, 64, 256)
	_, err := c.Read(in_buf)
	if err != nil {
		log.Fatal("Read error: ", err)
	}
	log.Println("Received", string(in_buf))

	msg := "pong\n"
	_, err = c.Write([]byte(msg))
	if err != nil {
		log.Fatal("Write error: ", err)
	}
	log.Println("Sent: ", msg)
}

func main() {

	conn, err := net.Dial("unix", "test.sock")
	if err != nil {
		log.Println("Dial error:", err)
		return
	}
	defer conn.Close()

	log.Println("Je suis connecté")
	handleConnection(conn)

}
