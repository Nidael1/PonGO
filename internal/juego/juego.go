package juego

import "github.com/hajimehoshi/ebiten/v2"

const (
	AnchoVentana = 640
	AltoVentana  = 480
)

type Juego struct{}

func Nuevo() *Juego { return &Juego{} }

func (j *Juego) Update() error { return nil }

func (j *Juego) Draw(screen *ebiten.Image) {}

func (j *Juego) Layout(_, _ int) (int, int) { return AnchoVentana, AltoVentana }
