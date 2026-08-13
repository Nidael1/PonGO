package rival

import "testing"

func TestDecidir_sube_cuando_bola_arriba(t *testing.T) {
	dir := Decidir(100, 200, 2)
	if dir != -1 {
		t.Errorf("esperado -1, obtenido %d", dir)
	}
}

func TestDecidir_baja_cuando_bola_abajo(t *testing.T) {
	dir := Decidir(350, 200, 3)
	if dir != 1 {
		t.Errorf("esperado 1, obtenido %d", dir)
	}
}

func TestDecidir_quieto_en_zona_muerta_nivel1(t *testing.T) {
	// Con umbral=45 para nivel 1, diferencia de 20 no debe mover al rival.
	dir := Decidir(240, 220, 1)
	if dir != 0 {
		t.Errorf("esperado 0, obtenido %d", dir)
	}
}

func TestDecidir_nivel_mayor_es_mas_reactivo(t *testing.T) {
	// diff=10: nivel 1 (umbral 45) no reacciona, nivel 3 (umbral 3) sí.
	bolaCentro, rivalCentro := 250.0, 240.0
	if Decidir(bolaCentro, rivalCentro, 1) != 0 {
		t.Error("nivel 1 no debería reaccionar a diff=10")
	}
	if Decidir(bolaCentro, rivalCentro, 3) != 1 {
		t.Error("nivel 3 sí debería reaccionar a diff=10")
	}
}
