package main

import (
	// "math/rand"
	//"log"

	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	// "log"
	//"net"

	"strconv"
	//"bufio"
)

// Mise à jour de l'état du jeu en fonction des entrées au clavier.
func (g *game) Update() error {
	g.stateFrame++
	switch g.gameState {
	case waitState:
		if g.waitOtherPlayer() {
			g.gameState++
		}
	case titleState:
		if g.titleUpdate() {
			g.gameState++
		}
	case colorSelectState:
		if g.colorSelectUpdate() {
			g.gameState++
		}
	case waitColorState:
		if g.waitColorUpdate() {
			g.gameState++
		}
	case playState:

		g.tokenPosUpdate()

		if g.PlayerId == 1 {
			if !g.isFirstTour {
				if g.haveListen1 == false {
					g.haveListen1 = true
					// routine qui attend le message de l'autre joueur
					go func() {
						msg_rcv_Pos, _ := g.in.ReadString(byte('\n'))
						pOthPos := string(msg_rcv_Pos)
						log.Println("j'ai recu : " + pOthPos)
						xValue, errX := strconv.Atoi(string(pOthPos[0]))
						if errX != nil {
							fmt.Println("Erreur de conversion en nombre entier.")
						}
						g.updateGrid(p2Token, xValue)
						g.turn = 1
						if string(pOthPos[1]) == "W" {
							g.result = p1wins
							g.gameState++
						} else if string(pOthPos[1]) == "L" {
							g.result = p2wins
							g.gameState++
						} else if string(pOthPos[1]) == "E" {
							g.result = equality
							g.gameState++
						}
					}()
				}
			}
			if g.turn == 1 {
				_, _ = g.p1Update()
			}

		} else if g.PlayerId == 2 {

			//recupere la position de l'autre joueur et met à jour
			if g.haveListen2 == false {
				g.haveListen2 = true
				// routine qui attend le message de l'autre joueur
				go func() {
					msg_rcv_Pos, _ := g.in.ReadString(byte('\n'))
					pOthPos := string(msg_rcv_Pos)
					log.Println("j'ai recu : " + pOthPos)
					xValue, _ := strconv.Atoi(string(pOthPos[0]))

					g.updateGrid(p1Token, xValue)
					g.turn = 2
					if string(pOthPos[1]) == "W" {
						g.result = p1wins
						g.gameState++
					} else if string(pOthPos[1]) == "L" {
						g.result = p2wins
						g.gameState++
					} else if string(pOthPos[1]) == "E" {
						g.result = equality
						g.gameState++
					}

				}()

			}
			if g.turn == 2 {
				_, _ = g.p2Update()
			}
		}

	case resultState:
		if g.resultUpdate() {
			g.reset1()
			g.reset()
			g.gameState = waitState
		}
	}

	return nil
}

// Mise à jour de l'état du jeu à l'écran titre.
func (g *game) titleUpdate() bool {

	g.stateFrame = g.stateFrame % globalBlinkDuration
	return inpututil.IsKeyJustPressed(ebiten.KeyEnter)
}

// fonction qui verifie que les 2j sont connéctés
func (g *game) waitOtherPlayer() bool {

	if !g.isFirstGame {
		go PlayerBeginNew(g)
		g.isFirstGame = true
	}
	if g.numberPlayer == 2 {

		return true
	}
	return false
}

func (g *game) waitColorUpdate() bool {

	if g.p2Color != -1 && g.p1Color != -1 {
		return true
	}
	return false
}

// Mise à jour de l'état du jeu lors de la sélection des couleurs.
func (g *game) colorSelectUpdate() bool {
	col := g.selectedColor % globalNumColorCol
	line := g.selectedColor / globalNumColorLine

	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		col = (col + 1) % globalNumColorCol
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		col = (col - 1 + globalNumColorCol) % globalNumColorCol
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		line = (line + 1) % globalNumColorLine
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		line = (line - 1 + globalNumColorLine) % globalNumColorLine
	}

	g.selectedColor = line*globalNumColorLine + col

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		if g.PlayerId == 2 {
			g.p2Color = g.selectedColor
		} else {
			g.p1Color = g.selectedColor
		}
		g.out.WriteString("C" + strconv.Itoa(g.selectedColor) + "\n")
		g.out.Flush()
		return true

	}

	return false
}

// Gestion de la position du prochain pion à jouer par le joueur 1.
func (g *game) tokenPosUpdate() {
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		g.tokenPosition = (g.tokenPosition - 1 + globalNumTilesX) % globalNumTilesX
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		g.tokenPosition = (g.tokenPosition + 1) % globalNumTilesX
	}
}

// Gestion du moment où le prochain pion est joué par le joueur 1.
func (g *game) p1Update() (int, int) {
	//récupere le pion du joueur 2

	//le joueur 1 place son pion et l'envoie au joueur 2
	lastXPositionPlayed := -1
	lastYPositionPlayed := -1
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		if updated, yPos := g.updateGrid(p1Token, g.tokenPosition); updated {
			g.turn = 2
			lastXPositionPlayed = g.tokenPosition
			lastYPositionPlayed = yPos

			// verification si qq un a gagné
			if lastXPositionPlayed >= 0 {
				finished, result := g.checkGameEnd(lastXPositionPlayed, lastYPositionPlayed)
				if finished {
					if result == p2wins {
						g.out.WriteString(strconv.Itoa(lastXPositionPlayed) + "W\n")
						g.out.Flush()
					} else if result == p1wins {
						g.out.WriteString(strconv.Itoa(lastXPositionPlayed) + "L\n")
						g.out.Flush()
					} else {
						g.out.WriteString(strconv.Itoa(lastXPositionPlayed) + "E\n")
						g.out.Flush()
					}

					g.result = result
					g.gameState++
				}
			}
			g.out.WriteString(strconv.Itoa(lastXPositionPlayed) + "R\n")
			g.out.Flush()
			g.isFirstTour = false
			g.haveListen1 = false
		}
	}

	return lastXPositionPlayed, lastYPositionPlayed

}

// Gestion de la position du prochain pion joué par le joueur 2 et
// du moment où ce pion est joué.
func (g *game) p2Update() (int, int) {

	// le joueur 2 met son pion et envoie au joueur 1
	lastXPositionPlayed := -1
	lastYPositionPlayed := -1
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		if updated, yPos := g.updateGrid(p2Token, g.tokenPosition); updated {
			g.turn = 1
			lastXPositionPlayed = g.tokenPosition
			lastYPositionPlayed = yPos
			// verification si qq un a gagné
			if lastXPositionPlayed >= 0 {
				finished, result := g.checkGameEnd(lastXPositionPlayed, lastYPositionPlayed)
				if finished {
					if result == p2wins {
						g.result = p1wins
						g.out.WriteString(strconv.Itoa(lastXPositionPlayed) + "L\n")
						g.out.Flush()
					} else if result == p1wins {
						g.result = p2wins
						g.out.WriteString(strconv.Itoa(lastXPositionPlayed) + "W\n")
						g.out.Flush()
					} else {
						g.out.WriteString(strconv.Itoa(lastXPositionPlayed) + "E\n")
						g.out.Flush()
					}

					g.gameState++
				}
			}
			g.out.WriteString(strconv.Itoa(lastXPositionPlayed) + "R\n")
			g.out.Flush()
			g.haveListen2 = false
		}
	}

	return lastXPositionPlayed, lastYPositionPlayed

}

// Mise à jour de l'état du jeu à l'écran des résultats.
func (g game) resultUpdate() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return true
	}
	return false
}

// Mise à jour de la grille de jeu lorsqu'un pion est inséré dans la
// colonne de coordonnée (x) position.
func (g *game) updateGrid(token, position int) (updated bool, yPos int) {
	for y := globalNumTilesY - 1; y >= 0; y-- {
		if g.grid[position][y] == noToken {
			updated = true
			yPos = y
			g.grid[position][y] = token
			return
		}
	}
	return
}
func (g *game) reset1() {
	g.p1Color = -1
	g.p2Color = -1
	g.numberPlayer = 1
	g.isFirstGame = false
	g.turn = 1
	g.result = 0
	g.haveListen1 = false
	g.haveListen2 = false
	g.isFirstTour = true
}

// Vérification de la fin du jeu : est-ce que le dernier joueur qui
// a placé un pion gagne ? est-ce que la grille est remplie sans gagnant
// (égalité) ? ou est-ce que le jeu doit continuer ?
func (g game) checkGameEnd(xPos, yPos int) (finished bool, result int) {

	tokenType := g.grid[xPos][yPos]

	// horizontal
	count := 0
	for x := xPos; x < globalNumTilesX && g.grid[x][yPos] == tokenType; x++ {
		count++
	}
	for x := xPos - 1; x >= 0 && g.grid[x][yPos] == tokenType; x-- {
		count++
	}

	if count >= 4 {
		if tokenType == p1Token {
			return true, p1wins
		}
		return true, p2wins
	}

	// vertical
	count = 0
	for y := yPos; y < globalNumTilesY && g.grid[xPos][y] == tokenType; y++ {
		count++
	}

	if count >= 4 {
		if tokenType == p1Token {
			return true, p1wins
		}
		return true, p2wins
	}

	// diag haut gauche/bas droit
	count = 0
	for x, y := xPos, yPos; x < globalNumTilesX && y < globalNumTilesY && g.grid[x][y] == tokenType; x, y = x+1, y+1 {
		count++
	}

	for x, y := xPos-1, yPos-1; x >= 0 && y >= 0 && g.grid[x][y] == tokenType; x, y = x-1, y-1 {
		count++
	}

	if count >= 4 {
		if tokenType == p1Token {
			return true, p1wins
		}
		return true, p2wins
	}

	// diag haut droit/bas gauche
	count = 0
	for x, y := xPos, yPos; x >= 0 && y < globalNumTilesY && g.grid[x][y] == tokenType; x, y = x-1, y+1 {
		count++
	}

	for x, y := xPos+1, yPos-1; x < globalNumTilesX && y >= 0 && g.grid[x][y] == tokenType; x, y = x+1, y-1 {
		count++
	}

	if count >= 4 {
		if tokenType == p1Token {
			return true, p1wins
		}
		return true, p2wins
	}

	// egalité ?
	if yPos == 0 {
		for x := 0; x < globalNumTilesX; x++ {
			if g.grid[x][0] == noToken {
				return
			}
		}
		return true, equality
	}

	return
}
