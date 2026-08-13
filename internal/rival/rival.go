package rival

// Velocidades máximas de la paleta rival por nivel (índice 0 sin uso).
// Valores provisionales; se calibran en issue #10.
var Velocidades = [4]float64{0, 1.5, 2.5, 3.8}

// umbral: zona muerta vertical. Si el centro de la bola está dentro de
// este margen respecto al centro de la paleta, el rival no se mueve.
// Nivel más alto = umbral menor = rival más reactivo.
var umbral = [4]float64{0, 45.0, 18.0, 3.0}

// Decidir devuelve -1 (subir), 0 (quieto) o 1 (bajar).
// bolaCentroY: coordenada Y del centro de la bola.
// rivalCentroY: coordenada Y del centro de la paleta rival.
func Decidir(bolaCentroY, rivalCentroY float64, nivel int) int {
	diff := bolaCentroY - rivalCentroY
	u := umbral[nivel]
	if diff > u {
		return 1
	}
	if diff < -u {
		return -1
	}
	return 0
}
