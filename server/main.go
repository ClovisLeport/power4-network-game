package main

import (
	"bufio"
	"log"
	"net"
	"time"
)

func server1(listener net.Listener, in1 *bufio.Reader, out1 *bufio.Writer, in2 *bufio.Reader, out2 *bufio.Writer, isFirstGame bool) {
	// envoie au joueur si l'autre est pret à faire une deuxième partie
	if !isFirstGame {
		go func() {
			msg_rcv1 := ""
			for msg_rcv1 != "N\n" {
				in1.ReadString(byte('\n'))
				msg_rcv1, _ = in1.ReadString(byte('\n'))
				log.Println(msg_rcv1)

			}
			log.Println("j'ai recu N 1" + msg_rcv1)
			out2.WriteString(msg_rcv1)
			out2.Flush()

		}()

		msg_rcv2 := ""
		for msg_rcv2 != "N\n" {
			msg_rcv2, _ = in2.ReadString(byte('\n'))
		}

		log.Println("j'ai recu N 2" + msg_rcv2)
		out1.WriteString(msg_rcv2)
		out1.Flush()

	}
	// permet de ne pas lancer le jeu tant que les deux couleur on pas été lancés
	bothColorDecide := 0

	// recupère la couleur des joueurs et renvoie
	go func() {
		msg_rcv1, err := in1.ReadString(byte('\n'))
		if err != nil {
			log.Fatal("client1 deconnécté")
		}
		log.Println("j'ai recu et j'envoie:" + msg_rcv1)
		out2.WriteString(msg_rcv1)
		out2.Flush()
		bothColorDecide++
	}()

	msg_rcv2, err := in2.ReadString(byte('\n'))
	if err != nil {
		log.Fatal("client2 deconnécté")
	}
	log.Println("j'ai recu : " + msg_rcv2)
	out1.WriteString(msg_rcv2)
	out1.Flush()

	PartieFini := true
	for PartieFini {
		if bothColorDecide == 1 {
			msg_rcv1, err := in1.ReadString(byte('\n'))

			if err != nil {
				log.Fatal("client1 deconnécté")
			}

			log.Println("j'ai recu 1: " + msg_rcv1)
			//log.Println("j'ai recu 1 : " + msg_rcv1)
			out2.WriteString(msg_rcv1)
			out2.Flush()
			if string(msg_rcv1[1]) == "W" || string(msg_rcv1[1]) == "L" || string(msg_rcv1[1]) == "E" {
				PartieFini = false

			}

			msg_rcv2, err := in2.ReadString(byte('\n'))
			if err != nil {
				log.Fatal("client2 deconnécté")
			}
			log.Println("j'ai recu 2 : " + msg_rcv2)

			//log.Println("j'ai recu 2: " + msg_rcv2)
			out1.WriteString(msg_rcv2)
			out1.Flush()

			if string(msg_rcv2[1]) == "W" || string(msg_rcv2[1]) == "L" || string(msg_rcv2[1]) == "E" {
				PartieFini = false

			}
		}

	}
	server1(listener, in1, out1, in2, out2, false)
	// defer conn.Close()
	// defer conn2.Close()

}
func connection(listener net.Listener) (*bufio.Reader, *bufio.Writer) {
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("accept error:", err)
	}
	in1 := bufio.NewReader(conn)
	out1 := bufio.NewWriter(conn)
	return in1, out1
}
func main() {
	// creation des clients
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Println("listen error:", err)
		return
	}
	defer listener.Close()

	//connection et  creation bufio writer et reader pour client 1
	in1, out1 := connection(listener)
	log.Println("Le client1 s'est connecté")

	// envoie au client 1 qu'il est le 1
	out1.WriteString("1\n")
	out1.Flush()

	//connection et  creation bufio writer et reader pour client 2
	in2, out2 := connection(listener)
	log.Println("Le client2 s'est connecté")

	// envoie au client 2 qu'il est le 2
	out2.WriteString("2\n")
	out2.Flush()

	time.Sleep(1 * time.Millisecond)
	// dit aux deux joueurs que les deux sont connéctés
	out1.WriteString("2j\n")
	out1.Flush()

	out2.WriteString("2j\n")
	out2.Flush()

	// jeu
	server1(listener, in1, out1, in2, out2, true)
}
