package main

import (
	"log"
	"net"
	"strconv"
	"bufio"
)


func PlayerBegin(g *game,c net.Conn,in *bufio.Reader, out *bufio.Writer) {


	// verifie si c'est le joueur 1 ou 2
	c.Read(in)
	numPlayer:= string(in)
	if ( numPlayer== "2") {
		g.turn = p2Turn
		g.PlayerId = 2
	}else if ( numPlayer== "1") {
		g.turn = p2Turn
		g.PlayerId = 1
	}

	// verifie si les deux joueurs sont connéctés

	c.Read(in)
	var Is2Player string = string(in)
	if (string(Is2Player[:2]) == "2j") {
		log.Println()
		log.Println(g.gameState)
		g.numberPlayer+=2
	}

	// met la couleur de l'autre joueur

	c.Read(in)
	p2Color:= string(in)
	log.Print("couleur de p2 :")
	log.Println(p2Color)
	if (string(p2Color[0]) == "C") {
		marks, _ := strconv.Atoi(string(p2Color[1]))
		g.p2Color = marks
	}

	log.Println(g.p1Color)

}