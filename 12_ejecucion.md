Propósito: dejar todo lo necesario para clonar el repo y tenerlo corriendo.

## Requisitos

- macOS.
- Go 1.22 o superior. La versión exacta la fija `go.mod`; esa es la que manda.
- Herramientas de línea de comandos de Xcode instaladas
  (`xcode-select --install`). Ebitengine compila contra frameworks del sistema
  en macOS y sin ellas el build falla.
- Nada más. Sin Docker, sin base de datos, sin servicios externos.

## Cómo correr en local

```
git clone https://github.com/Nidael1/PonGO.git
cd PonGO
go mod download
go run ./cmd/pongo
```

La primera compilación tarda: Ebitengine arrastra código nativo. Las
siguientes son rápidas.

## Cómo probar

```
go vet ./...
go test ./...
```

Y luego jugar. En este proyecto jugar no es opcional: es el método de
verificación principal.

## Estrategia de pruebas

- **Qué se prueba siempre:** el paquete `rival`. Es el único con lógica que se
  puede afirmar sin ventana ni teclado: dada una posición de bola, una posición
  de paleta y un nivel, la decisión de movimiento es determinista y
  verificable. También se prueba el conteo del marcador y la condición de fin
  de partida a dos puntos.
- **Qué NO se prueba y por qué:** el render, el audio y la lectura de teclado.
  Probarlos exige levantar una ventana y comparar píxeles o muestras de sonido;
  el costo es alto y el fallo se detecta de inmediato al ejecutar. Tampoco se
  prueba la física del rebote con casos automatizados: se valida jugando,
  contra los criterios de aceptación de `07_sprint_actual.md`.
- **Mínimo exigido para aceptar un commit ajeno:** `go vet ./...` y
  `go test ./...` en verde, y el juego arranca y se puede terminar una partida
  completa sin que el programa se cierre solo. Si el commit toca `rival`,
  además se vuelven a medir los niveles según el criterio de calibración.

## Cómo construir

```
go build -o pongo ./cmd/pongo
./pongo
```

El binario resultante corre en la máquina donde se compiló. No se compila para
otras plataformas: solo macOS, por frontera del proyecto.
