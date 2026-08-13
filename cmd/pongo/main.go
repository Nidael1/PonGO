package main

import (
	"log"

	"github.com/Nidael1/PonGO/internal/juego"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("PonGO")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)
	if err := ebiten.RunGame(juego.Nuevo()); err != nil {
		log.Fatal(err)
	}
}
