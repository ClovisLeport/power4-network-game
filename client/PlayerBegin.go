package main

import (
	"log"
	"strconv"
)

func PlayerBegin(g *game) {

	// verifie si c'est le joueur 1 ou 2
	if g.isFirstGame == true {

		msg_rcv_num, _ := g.in.ReadString(byte('\n'))
		numPlayer := string(msg_rcv_num)
		if string(numPlayer[0]) == "2" {
			g.PlayerId = 2
		} else {
			g.PlayerId = 1
		}
	} else {
		g.out.WriteString("N\n")
		g.out.Flush()
	}

	g.p1Color = -1
	g.p2Color = -1
	//c'est toujours le joueur 1 qui commence donc le premier à s'être connecté
	g.turn = 1
	g.haveListen1 = false
	g.haveListen2 = false
	g.isFirstTour = true

	// verifie si les deux joueurs sont connéctés

	msg_rcv_2j, _ := g.in.ReadString(byte('\n'))
	var Is2Player string = string(msg_rcv_2j)
	log.Println("j'ai recu N ? " + Is2Player)
	if string(Is2Player[:2]) == "2j" || string(Is2Player[0]) == "N" {
		g.numberPlayer++
	}

	// recoit et met la couleur de l'autre joueur
	msg_rcv_color, _ := g.in.ReadString(byte('\n'))
	pColor := string(msg_rcv_color)
	log.Println("j'ai recu la couleur ? " + pColor)

	// message commence par C pour dire que c'est la couleur qui est recu
	if string(pColor[0]) == "C" {
		marks, _ := strconv.Atoi(string(pColor[1]))
		if g.PlayerId == 1 {
			g.p2Color = marks
		} else {
			g.p1Color = marks
		}

	}

}
