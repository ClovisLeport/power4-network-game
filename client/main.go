package main

import (
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

func createClient(conn net.Conn) (numJoueurRet string) {
	
	//defer conn.Close()
	log.Println("Je suis connecté")
	var numJoueur string = handleConnectionRead(conn)
	
	return numJoueur
}

// Création, paramétrage et lancement du jeu.
func main() {

	g := game{}

	ebiten.SetWindowTitle("Programmation système : projet puissance 4")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Println("Dial error:", err)
		return
	}
	createClient(conn)
	go func() {
		var Is2Player string = handleConnectionRead(conn)
		if (string(Is2Player[:2]) == "2j") {
			log.Println()
			g.numberPlayer+=2
		}
	}()
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
	



	

}
