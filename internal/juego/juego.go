package juego

import (
	"github.com/Nidael1/PonGO/internal/audio"
	"github.com/Nidael1/PonGO/internal/entrada"
	"github.com/Nidael1/PonGO/internal/render"
	"github.com/Nidael1/PonGO/internal/rival"
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
	estadoPausado
)

type Juego struct {
	estado         estadoJuego
	nivel          int
	jugador        Paleta
	palRival       Paleta
	bola           Bola
	marcador       Marcador
	ultimoMarcador [2]int
	audio          *audio.Audio
}

func Nuevo() *Juego {
	return &Juego{
		estado: estadoSeleccion,
		nivel:  1,
		audio:  audio.Nueva(),
	}
}

func (j *Juego) Update() error {
	switch j.estado {
	case estadoSeleccion:
		j.updateSeleccion()
	case estadoEnJuego:
		if entrada.LeerPausa() {
			j.estado = estadoPausado
		} else {
			j.updateEnJuego()
		}
	case estadoPausado:
		if entrada.LeerPausa() {
			j.estado = estadoEnJuego
		}
	}
	return nil
}

func (j *Juego) Draw(screen *ebiten.Image) {
	switch j.estado {
	case estadoSeleccion:
		render.DibujarSeleccion(screen, j.nivel, j.ultimoMarcador)
	case estadoEnJuego, estadoPausado:
		render.DibujarJuego(
			screen,
			j.jugador.X, j.jugador.Y,
			j.palRival.X, j.palRival.Y,
			j.bola.X, j.bola.Y,
			[2]int{j.marcador.Jugador, j.marcador.Rival},
		)
		if j.estado == estadoPausado {
			render.DibujarPausa(screen)
		}
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
	j.moverRival()
	j.moverBola()
	j.verificarPunto()
}


func (j *Juego) moverRival() {
	bolaCentroY := j.bola.Y + TamBola/2
	rivalCentroY := j.palRival.Y + AltoPaleta/2
	dir := rival.Decidir(bolaCentroY, rivalCentroY, j.nivel)
	j.palRival.Y += float64(dir) * rival.Velocidades[j.nivel]
	j.palRival.Y = clamp(j.palRival.Y, 0, AltoVentana-AltoPaleta)
}

func (j *Juego) verificarPunto() {
	if j.bola.X+TamBola < 0 {
		j.marcador.Rival++
		j.audio.PlayPunto()
		if !j.verificarFin() {
			j.lanzarBola(1)
		}
	}
	if j.bola.X > AnchoVentana {
		j.marcador.Jugador++
		j.audio.PlayPunto()
		if !j.verificarFin() {
			j.lanzarBola(-1)
		}
	}
}

// verificarFin retorna true si la partida terminó.
func (j *Juego) verificarFin() bool {
	if j.marcador.Jugador >= PuntosParaGanar || j.marcador.Rival >= PuntosParaGanar {
		j.ultimoMarcador = [2]int{j.marcador.Jugador, j.marcador.Rival}
		j.audio.PlayFin()
		j.estado = estadoSeleccion
		return true
	}
	return false
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
