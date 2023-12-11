package main

import (
	"bufio"
	"log"
	"net"
	"time"
)

func handleClientRead(c net.Conn) (value string) {
	in := bufio.NewReader(c)

	msg_rcv, err := in.ReadString(byte('\n'))
	if err != nil {
		log.Fatal("Read error: ", err)
	}
	log.Println("Received", msg_rcv)
	return string(msg_rcv)
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
	handleClientWrite(conn, "1")

	conn2, err := listener.Accept()

	if err != nil {
		log.Println("accept error:", err)
		return
	}
	log.Println("Le client2 s'est connecté")
	handleClientWrite(conn2, "2")
	time.Sleep(1 * time.Second)

	// dit aux deux joueurs que les deux sont connéctés
	handleClientWrite(conn, "2j")
	handleClientWrite(conn2, "2j")

	//// SA NE DEPASSE PAS CET ENDROIT !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!! je ne sais pas si ca a envoye les 2j
	// recupère la couleur des joueurs et renvoie
	in_j1 := bufio.NewReader(conn)
	msg_rcv1, err := in_j1.ReadString(byte('\n'))
	out1 := bufio.NewWriter(conn2)
	out1.WriteString(msg_rcv1)


	in_j2 := bufio.NewReader(conn2)
	msg_rcv2, err := in_j2.ReadString(byte('\n'))
	out2 := bufio.NewWriter(conn)
	out2.WriteString(msg_rcv2)

	PartieFini := true
	go func() {
		for (PartieFini) {
			valueP1 := handleClientRead(conn)
			if (valueP1[:1] == "w" || valueP1[:1] == "l" ) {
				handleClientWrite(conn2, "w1")
				PartieFini = false
				break
			}
			handleClientWrite(conn2, valueP1)
	
			valueP2 := handleClientRead(conn2)
			if (valueP1[:1] == "w" || valueP1[:1] == "l" ) {
				handleClientWrite(conn, "w1")
				PartieFini = false
				break
			}
			handleClientWrite(conn, valueP2)
		}
	}()
	
	defer conn.Close()
	defer conn2.Close()

}

func main() {
	server1()
}
