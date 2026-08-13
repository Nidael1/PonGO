package juego

import (
	"github.com/Nidael1/PonGO/internal/entrada"
	"github.com/Nidael1/PonGO/internal/render"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	AnchoVentana     = 640
	AltoVentana      = 480
	AnchoPaleta      = 10
	AltoPaleta       = 80
	TamBola          = 10
	XJugador         = 20.0
	XRival           = AnchoVentana - XJugador - AnchoPaleta
	VelocidadJugador = 4.0
	VelocidadBola    = 3.5
	PuntosParaGanar  = 2
)

type estadoJuego int

const (
	estadoSeleccion estadoJuego = iota
	estadoEnJuego
)

type Juego struct {
	estado         estadoJuego
	nivel          int
	jugador        Paleta
	palRival       Paleta
	bola           Bola
	marcador       Marcador
	ultimoMarcador [2]int
}

func Nuevo() *Juego {
	return &Juego{
		estado: estadoSeleccion,
		nivel:  1,
	}
}

func (j *Juego) Update() error {
	switch j.estado {
	case estadoSeleccion:
		j.updateSeleccion()
	case estadoEnJuego:
		j.updateEnJuego()
	}
	return nil
}

func (j *Juego) Draw(screen *ebiten.Image) {
	switch j.estado {
	case estadoSeleccion:
		render.DibujarSeleccion(screen, j.nivel, j.ultimoMarcador)
	case estadoEnJuego:
		render.DibujarJuego(
			screen,
			j.jugador.X, j.jugador.Y,
			j.palRival.X, j.palRival.Y,
			j.bola.X, j.bola.Y,
			[2]int{j.marcador.Jugador, j.marcador.Rival},
		)
	}
}

func (j *Juego) Layout(_, _ int) (int, int) { return AnchoVentana, AltoVentana }

func (j *Juego) updateSeleccion() {
	n := entrada.LeerNivel()
	if n > 0 {
		j.nivel = n
		j.iniciarPartida()
	}
}

func (j *Juego) iniciarPartida() {
	j.marcador = Marcador{}
	j.jugador = Paleta{X: XJugador, Y: AltoVentana/2 - AltoPaleta/2}
	j.palRival = Paleta{X: XRival, Y: AltoVentana/2 - AltoPaleta/2}
	j.lanzarBola(1)
	j.estado = estadoEnJuego
}

func (j *Juego) updateEnJuego() {
	j.moverJugador()
	j.moverBola()
}

func (j *Juego) moverJugador() {
	intent := entrada.LeerPaleta()
	switch intent {
	case entrada.Subir:
		j.jugador.Y -= VelocidadJugador
	case entrada.Bajar:
		j.jugador.Y += VelocidadJugador
	}
	j.jugador.Y = clamp(j.jugador.Y, 0, AltoVentana-AltoPaleta)
}

func clamp(v, min, max float64) float64 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}
