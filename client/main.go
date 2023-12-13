package main

import (
	"bufio"
	"log"
	"net"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font/opentype"
)

// Mise en place des polices d'écritures utilisées pour l'affichage.
func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	smallFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size: 30,
		DPI:  72,
	})
	if err != nil {
		log.Fatal(err)
	}

	largeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size: 50,
		DPI:  72,
	})
	if err != nil {
		log.Fatal(err)
	}
}

// Création d'une image annexe pour l'affichage des résultats.
func init() {
	offScreenImage = ebiten.NewImage(globalWidth, globalHeight)
}

// Création, paramétrage et lancement du jeu.
func main() {

	g := game{}

	ebiten.SetWindowTitle("Programmation système : projet puissance 4")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	// connexion au server
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		log.Println("Dial error:", err)
		return
	}
	log.Println("Je suis connecté")

	// creation des bufio reader et writer pour le client
	g.in = bufio.NewReader(conn)
	g.out = bufio.NewWriter(conn)

	// fonction qui verifie que les deux joueurs sont connéctés et ont choisi leur couleurs

	go PlayerBegin(&g)

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}

}
