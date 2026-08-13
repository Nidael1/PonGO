Propósito: declarar el minisprint en curso, sus tareas y cuándo se da por cerrado.

Sprint 1.0 — abierto el 2026-08-13.

## Objetivo del minisprint

Tener PonGO jugable de punta a punta en macOS: ventana, las dos paletas, la
bola, marcador a dos puntos, los tres niveles seleccionables y los tres
sonidos. Al terminar el sprint, el objetivo de fondo del proyecto queda
comprobado: la instalación local de Go compila y ejecuta un programa gráfico
completo.

## Alcance — qué queda fuera

Fuera de este sprint, y del proyecto entero, todo lo listado en las fronteras
de `01_contexto.md`. Fuera de este sprint pero vivo en el backlog: los ajustes
finos que solo se detectan jugando (issue #11) y el afinado del timbre de los
sonidos (issue #12).

La calibración de los niveles (issue #10) sí entra en este sprint. Sin ella los
tres niveles existen pero no cumplen su definición.

## Checklist

- [x] #1 esquema raíz
- [x] #2 inicializar módulo Go y agregar Ebitengine
- [ ] #3 ventana y bucle de juego vacío
- [ ] #4 paleta del jugador con flechas
- [ ] #5 bola con rebotes en bordes y paletas
- [ ] #6 paleta del rival
- [ ] #7 marcador y fin de partida a dos puntos
- [ ] #8 pantalla de selección de nivel con marcador final
- [ ] #9 sonidos sintetizados de rebote, punto y fin
- [ ] #10 calibrar los valores de los tres niveles

## Criterios de aceptación

- El binario compila con `go build ./...` sin advertencias y arranca abriendo
  una ventana en macOS.
- Las flechas mueven la paleta izquierda y esta no sale del área.
- La bola rebota contra bordes y paletas, y no se queda atrapada dentro de una
  paleta en veinte partidas seguidas.
- Al cruzar la bola un borde lateral, el marcador de ese lado sube en uno.
- Al llegar cualquier marcador a 2, la partida termina y aparece la pantalla de
  selección con el marcador final visible.
- Desde la selección se elige nivel y arranca una partida nueva en 0–0, sin
  cerrar el programa.
- Los tres sonidos suenan en su momento y se distinguen entre sí de oído.
- Calibración (issue #10), medida jugando diez partidas por nivel:
  - Nivel 1: el jugador gana al menos 8 de 10.
  - Nivel 2: el jugador gana entre 4 y 6 de 10.
  - Nivel 3: el jugador gana 2 de 10 o menos, pero gana al menos una.
  El criterio se mide con el jugador jugando en serio, no dejándose perder. Si
  un nivel queda fuera de rango, se ajusta y se vuelve a medir.
