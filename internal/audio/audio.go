package audio

import (
	"bytes"
	"math"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

const sampleRate = 44100

// Audio agrupa el contexto y los datos PCM pre-generados de los tres sonidos.
type Audio struct {
	ctx        *audio.Context
	reboteData []byte
	puntoData  []byte
	finData    []byte
}

func Nueva() *Audio {
	ctx := audio.NewContext(sampleRate)
	return &Audio{
		ctx:        ctx,
		reboteData: generar(880, 0.04),  // tono agudo, muy corto
		puntoData:  generar(523, 0.20),  // Do5, medio
		finData:    generar(220, 0.55),  // La3, largo y grave
	}
}

func (a *Audio) PlayRebote() { a.play(a.reboteData) }
func (a *Audio) PlayPunto()  { a.play(a.puntoData) }
func (a *Audio) PlayFin()    { a.play(a.finData) }

func (a *Audio) play(data []byte) {
	p, err := a.ctx.NewPlayer(bytes.NewReader(data))
	if err != nil {
		return
	}
	p.Play()
}

// generar produce PCM stereo 16-bit little-endian con una onda sinusoidal
// a la frecuencia dada, con envolvente de decaimiento lineal.
func generar(freq, durSegs float64) []byte {
	n := int(sampleRate * durSegs)
	buf := make([]byte, n*4)
	for i := 0; i < n; i++ {
		t := float64(i) / sampleRate
		v := math.Sin(2 * math.Pi * freq * t)
		env := 1.0 - float64(i)/float64(n)
		s := int16(v * env * 0x5fff)
		buf[i*4] = byte(s)
		buf[i*4+1] = byte(s >> 8)
		buf[i*4+2] = byte(s)
		buf[i*4+3] = byte(s >> 8)
	}
	return buf
}
