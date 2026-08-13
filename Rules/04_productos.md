Propósito: decir de qué partes se compone el binario y qué puede depender de qué.

## Módulos

PonGO es un solo binario ejecutable. No hay productos separados. Dentro del
binario hay cinco piezas:

| Módulo | Responsabilidad |
|---|---|
| `juego` | Estado global (selección o juego), marcador, fin de partida, bucle. |
| `entrada` | Lectura del teclado. Traduce teclas a intención: subir, bajar, elegir nivel. |
| `rival` | Decide hacia dónde mueve la paleta derecha, según el nivel. |
| `render` | Dibuja paletas, bola, marcador y la pantalla de selección. |
| `audio` | Sintetiza y reproduce los tres sonidos. |

La separación existe por una razón concreta y no por ceremonia: `rival` es lo
único que se va a tocar repetidamente durante la calibración de los niveles
(issue #10). Aislarlo evita andar buscando números de dificultad regados entre
el bucle y el dibujo.

## Límites entre módulos

- `juego` es el único que conoce el estado completo. Los demás reciben lo que
  necesitan y devuelven un resultado.
- `rival` no dibuja, no lee teclado y no sabe el marcador. Recibe la posición
  de la bola y la suya, y devuelve hacia dónde moverse.
- `render` no modifica estado. Solo lee y dibuja.
- `entrada` no interpreta reglas: devuelve intención, no consecuencias.
- `audio` no decide cuándo suena algo; `juego` se lo pide.

## Dependencias permitidas

```
juego → entrada, rival, render, audio
render → (solo Ebitengine)
audio  → (solo Ebitengine)
entrada → (solo Ebitengine)
rival  → (nada)
```

Ninguna dependencia en sentido contrario. `rival` sin dependencias es
deliberado: es la pieza que se prueba sola.

Dependencias externas del proyecto: Ebitengine (`github.com/hajimehoshi/ebiten/v2`)
y la librería estándar. Nada más. No se usa `go-lib`, la librería interna
compartida: es infraestructura de servidor y aquí no hay servidor.
