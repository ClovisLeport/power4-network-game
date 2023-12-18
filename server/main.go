package main

import (
	"bufio"
	"log"
	"net"
	"time"
)

func server1(conn net.Conn, conn2 net.Conn, in1 *bufio.Reader, out1 *bufio.Writer, in2 *bufio.Reader, out2 *bufio.Writer, isFirstGame bool) {
	// envoie au joueur si l'autre est pret à faire une deuxième partie
	if !isFirstGame {
		in1.ReadString(byte('\n'))
		msg_rcv1, _ := in1.ReadString(byte('\n'))
		log.Println("j'envoie et j'ai recu : " + string(msg_rcv1))
		out2.WriteString(msg_rcv1)
		out2.Flush()

		msg_rcv2, _ := in2.ReadString(byte('\n'))
		out1.WriteString(msg_rcv2)
		out1.Flush()
	}
	// recupère la couleur des joueurs et renvoie
	msg_rcv1, _ := in1.ReadString(byte('\n'))
	out2.WriteString(msg_rcv1)
	out2.Flush()

	msg_rcv2, _ := in2.ReadString(byte('\n'))
	out1.WriteString(msg_rcv2)
	out1.Flush()

	PartieFini := true
	for PartieFini {
		msg_rcv1, _ := in1.ReadString(byte('\n'))
		log.Println("j'ai recu 1 : " + msg_rcv1)
		out2.WriteString(msg_rcv1)
		out2.Flush()
		if string(msg_rcv1[1]) == "W" || string(msg_rcv1[1]) == "L" || string(msg_rcv1[1]) == "E" {
			PartieFini = false
			break
		}

		msg_rcv2, _ := in2.ReadString(byte('\n'))
		log.Println("j'ai recu 2: " + msg_rcv2)
		out1.WriteString(msg_rcv2)
		out1.Flush()

		if string(msg_rcv2[1]) == "W" || string(msg_rcv2[1]) == "L" || string(msg_rcv2[1]) == "E" {
			PartieFini = false
			break

		}

	}
	log.Println("je sors")
	server1(conn, conn2, in1, out1, in2, out2, false)
	// defer conn.Close()
	// defer conn2.Close()

}

func main() {
	// creation des clients
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
	out1.WriteString("1\n")
	out1.Flush()
	conn2, err := listener.Accept()
	if err != nil {
		log.Println("accept error:", err)
		return
	}
	// creation bufio writer et reader pour client 2
	in2 := bufio.NewReader(conn2)
	out2 := bufio.NewWriter(conn2)

	// envoie au client 1 qu'il est le 2 qu'il est le 2
	out2.WriteString("2\n")
	out2.Flush()
	log.Println("Le client2 s'est connecté")
	time.Sleep(1 * time.Millisecond)
	// dit aux deux joueurs que les deux sont connéctés
	out1.WriteString("2j\n")
	out1.Flush()

	out2.WriteString("2j\n")
	out2.Flush()

	// jeu
	server1(conn, conn2, in1, out1, in2, out2, true)
}
