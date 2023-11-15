package main

import (
	"bufio"
	"log"
	"net"
	"time"

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

func handleClient(c net.Conn) {
	in := bufio.NewReader(c)
	out := bufio.NewWriter(c)

	msg := "ping\n"
	_, err := out.WriteString(msg)
	if err != nil {
		log.Fatal("Write error: ", err)
	}
	out.Flush()
	log.Println("Sent: ", msg)

	msg_rcv, err := in.ReadString(byte('\n'))
	if err != nil {
		log.Fatal("Read error: ", err)
	}
	log.Println("Received", msg_rcv)

}

func handleConnection(c net.Conn) {
	in_buf := make([]byte, 64, 256)
	_, err := c.Read(in_buf)
	if err != nil {
		log.Fatal("Read error: ", err)
	}
	log.Println("Received", string(in_buf))

	msg := "pong\n"
	_, err = c.Write([]byte(msg))
	if err != nil {
		log.Fatal("Write error: ", err)
	}
	log.Println("Sent: ", msg)
}
func createServ() {
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
	defer conn.Close()
	/*log.Println("Le client s'est connecté")*/
	handleClient(conn)

	time.Sleep(10 * time.Second)

}
func createClient() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Println("Dial error:", err)
		return
	}
	defer conn.Close()

	log.Println("Je suis connecté")
	handleConnection(conn)
}

// Création, paramétrage et lancement du jeu.
func main() {

	g := game{}

	ebiten.SetWindowTitle("Programmation système : projet puissance 4")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	createServ()
	createClient()
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}

}
