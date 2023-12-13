package main

import (
	"log"
	"strconv"
)

func PlayerBegin(g *game) {

	// verifie si c'est le joueur 1 ou 2
	msg_rcv_num, _ := g.in.ReadString(byte('\n'))
	numPlayer := string(msg_rcv_num)
	log.Println(numPlayer)
	if numPlayer[0] == byte(2) {
		g.turn = p2Turn
		g.PlayerId = 2

	} else if numPlayer[0] == byte(1) {
		g.turn = p1Turn
		g.PlayerId = 1
	}
	log.Println(g.PlayerId)
	// verifie si les deux joueurs sont connéctés
	msg_rcv_2j, _ := g.in.ReadString(byte('\n'))
	var Is2Player string = string(msg_rcv_2j)
	if string(Is2Player[:1]) == "2j" {

		g.numberPlayer += 2
	}

	// met la couleur de l'autre joueur
	msg_rcv_color, _ := g.in.ReadString(byte('\n'))
	p2Color := string(msg_rcv_color)
	log.Print("couleur de p2 :")
	log.Println(p2Color)
	if string(p2Color[0]) == "C" {
		marks, _ := strconv.Atoi(string(p2Color[1]))
		g.p2Color = marks
	}

	log.Println(g.p1Color)

}
