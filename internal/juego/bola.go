package juego

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

func (j *Juego) moverBola() {}
