Propósito: describir cómo se comporta el juego, en especial el rival.

## Reglas de negocio

### Los tres niveles

Es la única regla propia de este proyecto. Los niveles **se definen por el
resultado esperado, no por su mecanismo**:

| Nivel | Resultado esperado |
|---|---|
| 1 — Fácil | El jugador le gana casi siempre. |
| 2 — Parejo | El jugador le gana alrededor de la mitad de las veces. |
| 3 — Difícil | Al jugador le cuesta trabajo ganarle. |

Los valores concretos que producen ese resultado —rapidez de la paleta rival,
retraso de su reacción, rapidez de la bola— **no están decididos**. Se calibran
jugando, contra los criterios de aceptación de `07_sprint_actual.md`.
Depende del issue #10.

Ninguna otra sección de esta documentación afirma un número de dificultad
mientras ese issue siga abierto.

Regla que sí queda fija ahora, porque es la que hace calibrable lo demás: los
tres niveles mueven las mismas palancas en la misma dirección. Un nivel más
alto nunca es *distinto*, es *más duro*. Si al calibrar resulta que el nivel 3
es más fácil que el 2 en alguna dimensión, es un bug, no una variante.

### Fin de partida

La partida termina en cuanto un lado llega a 2 puntos. Termine como termine
—ganando o perdiendo el jugador— el programa vuelve a la pantalla de selección
de nivel, y ahí se muestra el marcador final. Desde esa pantalla se elige nivel
y empieza una partida nueva desde 0–0.

### Control

El jugador mueve su paleta con las flechas arriba y abajo. No hay otra entrada
durante el juego más que la selección de nivel en la pantalla de selección.

### Sonido

Tres sonidos, sintetizados en código, sin archivos: rebote de la bola contra
una paleta o contra un borde, punto anotado, y fin de partida. Timbre y
duración son libres; lo único exigido es que los tres se distingan entre sí.

## Casos borde

- **Bola que entra casi horizontal.** Si tras un rebote el ángulo queda tan
  plano que la bola tarda demasiado en cruzar, se corrige el ángulo. Se
  detecta jugando; queda como issue #11 en el backlog, no bloquea el sprint.
- **Dos flechas a la vez.** Arriba y abajo presionadas simultáneamente: la
  paleta no se mueve.
- **Bola atrapada dentro de una paleta.** Si la detección de colisión permite
  que la bola quede solapada con la paleta y rebote en bucle, se separa la bola
  al detectar el impacto en lugar de solo invertir dirección.
- **Elegir nivel con partida en curso.** No puede ocurrir: durante el juego no
  hay entrada de selección.

## Qué se declara genérico y no se explica

Lo siguiente es un Pong estándar y se implementa como cualquiera esperaría, sin
documentación adicional:

- La bola rebota invirtiendo la componente correspondiente al chocar contra un
  borde horizontal o contra una paleta.
- El punto se anota cuando la bola cruza un borde lateral.
- Tras un punto, la bola vuelve al centro y se saca de nuevo.
- El marcador se dibuja arriba, uno de cada lado.
- El bucle de juego actualiza estado y luego dibuja, a la cadencia que impone
  Ebitengine.
