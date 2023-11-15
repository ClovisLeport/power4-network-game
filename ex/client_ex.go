package main

import (
	"log"
	"net"
)

func handleConnection(c net.Conn) {
	in_buf := make([]byte, 64, 256)
	_, err := c.Read(in_buf)
	if err != nil {
		log.Fatal("Read error : ", err)
	}
	log.Print("received ", string(in_buf))

	msg := "pong"
	_, err2 := c.Write([]byte(msg))
	if err2 != nil {
		log.Fatal("write error :", err)
	}
	log.Println("sent: ", msg)

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
