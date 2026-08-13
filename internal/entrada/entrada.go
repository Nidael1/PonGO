package entrada

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Intencion int

const (
	Quieto Intencion = iota
	Subir
	Bajar
)

func LeerPaleta() Intencion {
	arriba := ebiten.IsKeyPressed(ebiten.KeyArrowUp)
	abajo := ebiten.IsKeyPressed(ebiten.KeyArrowDown)
	if arriba && abajo {
		return Quieto
	}
	if arriba {
		return Subir
	}
	if abajo {
		return Bajar
	}
	return Quieto
}

// LeerPausa devuelve true si se acaba de presionar la barra espaciadora.
func LeerPausa() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeySpace)
}

// LeerNivel devuelve 1-3 si se acaba de presionar esa tecla, 0 si no.
func LeerNivel() int {
	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		return 1
	}
	if inpututil.IsKeyJustPressed(ebiten.Key2) {
		return 2
	}
	if inpututil.IsKeyJustPressed(ebiten.Key3) {
		return 3
	}
	return 0
}
