package juego

import "math"

type Bola struct {
	X, Y   float64
	DX, DY float64
}

func (j *Juego) lanzarBola(dirX float64) {
	j.bola = Bola{
		X:  AnchoVentana/2 - TamBola/2,
		Y:  AltoVentana/2 - TamBola/2,
		DX: VelocidadBola * dirX,
		DY: VelocidadBola * 0.6,
	}
}

func (j *Juego) moverBola() {
	j.bola.X += j.bola.DX
	j.bola.Y += j.bola.DY

	// Rebote pared superior
	if j.bola.Y <= 0 {
		j.bola.Y = 0
		j.bola.DY = math.Abs(j.bola.DY)
	}
	// Rebote pared inferior
	if j.bola.Y+TamBola >= AltoVentana {
		j.bola.Y = AltoVentana - TamBola
		j.bola.DY = -math.Abs(j.bola.DY)
	}

	// Rebote paleta jugador
	if colisionPaleta(j.bola, j.jugador) {
		j.bola.X = j.jugador.X + AnchoPaleta
		j.bola.DX = math.Abs(j.bola.DX)
		j.ajustarAngleDY(&j.bola, j.jugador.Y)
	}

	// Rebote paleta rival
	if colisionPaleta(j.bola, j.palRival) {
		j.bola.X = j.palRival.X - TamBola
		j.bola.DX = -math.Abs(j.bola.DX)
		j.ajustarAngleDY(&j.bola, j.palRival.Y)
	}
}

// ajustarAngleDY ajusta DY según dónde golpeó la bola en la paleta.
func (j *Juego) ajustarAngleDY(b *Bola, paletaY float64) {
	centro := paletaY + AltoPaleta/2
	offset := (b.Y + TamBola/2 - centro) / (AltoPaleta / 2)
	speed := math.Sqrt(b.DX*b.DX + b.DY*b.DY)
	b.DY = offset * speed * 0.85
	// Garantizar mínimo vertical para no quedar horizontal puro
	if math.Abs(b.DY) < 0.5 {
		if b.DY >= 0 {
			b.DY = 0.5
		} else {
			b.DY = -0.5
		}
	}
}

func colisionPaleta(b Bola, p Paleta) bool {
	return b.X < p.X+AnchoPaleta &&
		b.X+TamBola > p.X &&
		b.Y < p.Y+AltoPaleta &&
		b.Y+TamBola > p.Y
}
