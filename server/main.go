package main

import (
	"bufio"
	"log"
	"net"
	"time"
)



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
	// creation bufio writer et reader pour client 1
	in1 := bufio.NewReader(conn)
	out1 := bufio.NewWriter(conn)
	
	// envoie au client 1 qu'il est le 1
	out1.WriteString("1")


	conn2, err := listener.Accept()
	if err != nil {
		log.Println("accept error:", err)
		return
	}
	log.Println("Le client2 s'est connecté")
	// creation bufio writer et reader pour client 2
	in2 := bufio.NewReader(conn2)
	out2 := bufio.NewWriter(conn2)
	// envoie au client 1 qu'il est le 2
	out1.WriteString("2")

	time.Sleep(1 * time.Second)

	// dit aux deux joueurs que les deux sont connéctés
	out1.WriteString("2j")
	out2.WriteString("2j")


	// recupère la couleur des joueurs et renvoie

	msg_rcv1, err := in1.ReadString(byte('\n'))
	log.Print("couleur choisi")
	log.Println(msg_rcv1)
	out2.WriteString(msg_rcv1)

	msg_rcv2, err := in2.ReadString(byte('\n'))
	out1.WriteString(msg_rcv2)




	//PartieFini := true
	// go func() {
	// 	for (PartieFini) {
	// 		valueP1 := handleClientRead(conn)
	// 		if (valueP1[:1] == "w" || valueP1[:1] == "l" ) {
	// 			handleClientWrite(conn2, "w1")
	// 			PartieFini = false
	// 			break
	// 		}
	// 		handleClientWrite(conn2, valueP1)
	
	// 		valueP2 := handleClientRead(conn2)
	// 		if (valueP1[:1] == "w" || valueP1[:1] == "l" ) {
	// 			handleClientWrite(conn, "w1")
	// 			PartieFini = false
	// 			break
	// 		}
	// 		handleClientWrite(conn, valueP2)
	// 	}
	// }()
	
	// defer conn.Close()
	// defer conn2.Close()

}

func main() {
	server1()
}
