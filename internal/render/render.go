package render

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	anchoPaleta  = 10
	altoPaleta   = 80
	tamBola      = 10
	anchoVentana = 640
	altoVentana  = 480
)

var (
	blanco = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	gris   = color.RGBA{R: 80, G: 80, B: 80, A: 255}
	negro  = color.RGBA{R: 0, G: 0, B: 0, A: 255}
)

func DibujarJuego(
	screen *ebiten.Image,
	jugX, jugY float64,
	rivX, rivY float64,
	bolX, bolY float64,
	pts [2]int,
) {
	screen.Fill(negro)

	// Línea central punteada
	for y := 0; y < altoVentana; y += 20 {
		ebitenutil.DrawRect(screen, anchoVentana/2-1, float64(y), 2, 10, gris)
	}

	// Paletas
	ebitenutil.DrawRect(screen, jugX, jugY, anchoPaleta, altoPaleta, blanco)
	ebitenutil.DrawRect(screen, rivX, rivY, anchoPaleta, altoPaleta, blanco)

	// Bola
	ebitenutil.DrawRect(screen, bolX, bolY, tamBola, tamBola, blanco)

	// Marcador
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d", pts[0]), anchoVentana/2-30, 10)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d", pts[1]), anchoVentana/2+20, 10)
}

func DibujarSeleccion(screen *ebiten.Image, nivel int, ultimo [2]int) {
	screen.Fill(negro)

	ebitenutil.DebugPrintAt(screen, "P O N G O", anchoVentana/2-40, 80)

	if ultimo[0] > 0 || ultimo[1] > 0 {
		ebitenutil.DebugPrintAt(screen,
			fmt.Sprintf("Ultimo marcador: %d - %d", ultimo[0], ultimo[1]),
			anchoVentana/2-70, 140)
	}

	ebitenutil.DebugPrintAt(screen, "Elige nivel:", anchoVentana/2-50, 210)
	ebitenutil.DebugPrintAt(screen, "  1  Facil", anchoVentana/2-40, 240)
	ebitenutil.DebugPrintAt(screen, "  2  Parejo", anchoVentana/2-40, 260)
	ebitenutil.DebugPrintAt(screen, "  3  Dificil", anchoVentana/2-40, 280)

	marcador := " "
	if nivel > 0 {
		marcador = fmt.Sprintf("> %d", nivel)
	}
	ebitenutil.DebugPrintAt(screen, marcador, anchoVentana/2-75, 240+((nivel-1)*20))
}
