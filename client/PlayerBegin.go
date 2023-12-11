package main

import (
	"log"
	"net"
	"strconv"
)
func handleConnectionWrite(c net.Conn, msg string) {

	_, err := c.Write([]byte(msg))
	if err != nil {
		log.Fatal("Write error: ", err)
	}
	log.Println("Sent: ", msg)
}

func handleConnectionRead(c net.Conn) (value string) {
	in_buf := make([]byte, 64, 256)
	_, err := c.Read(in_buf)
	if err != nil {
		log.Fatal("Read error: ", err)

	}
	log.Println("Received", string(in_buf))
	return string(in_buf)

}

func PlayerBegin(g game,c net.Conn) {


	// verifie si c'est le joueur 1 ou 2
	in_buf_j := make([]byte, 64, 256)
	c.Read(in_buf_j)
	numPlayer:= string(in_buf_j)
	if ( numPlayer== "2") {
		g.turn = p2Turn
		g.PlayerId = 2
	}else if ( numPlayer== "1") {
		g.turn = p2Turn
		g.PlayerId = 1
	}

	// verifie si les deux joueurs sont connéctés
	in_buf_2j := make([]byte, 64, 256)
	c.Read(in_buf_2j)
	var Is2Player string = string(in_buf_2j)
	if (string(Is2Player[:2]) == "2j") {
		log.Println()
		g.numberPlayer+=2
	}


	// met la couleur de l'autre joueur
	in_buf_c := make([]byte, 64, 256)
	c.Read(in_buf_c)
	p2Color:= string(in_buf_c)
	if (string(p2Color[0]) == "C") {
		marks, _ := strconv.Atoi(string(p2Color[1]))
		g.p2Color = marks
	}

	log.Println(g.p1Color)

}