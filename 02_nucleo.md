Propósito: fijar las entidades del juego y las condiciones que nunca se rompen.

## Entidades

| Entidad | Qué es | Estado que carga |
|---|---|---|
| Paleta | Rectángulo vertical que devuelve la bola. Hay dos: jugador y rival. | Posición vertical, alto, velocidad máxima de desplazamiento. |
| Bola | Círculo o cuadro que se desplaza y rebota. Hay una sola. | Posición, dirección, rapidez. |
| Marcador | Puntos de cada lado en la partida en curso. | Dos enteros, de 0 a 2. |
| Nivel | Ajuste de dificultad del rival. Hay exactamente tres. | Identificador y los valores que lo definen. |
| Partida | Una serie desde el saque inicial hasta que alguien llega a dos. | Marcador, nivel elegido, si sigue viva. |

El programa está siempre en uno de dos estados: **selección** o **juego**. No
hay un tercer estado de fin de partida: terminar la partida es volver a
selección con el marcador final a la vista.

## Invariantes

- Ninguna paleta sale del área de juego por arriba ni por abajo.
- Mientras el estado es juego, la bola nunca está detenida ni se mueve en
  vertical pura: siempre avanza hacia alguno de los dos lados.
- El marcador nunca pasa de 2 en ningún lado. Al llegar a 2, la partida termina
  en ese mismo instante.
- Solo hay una partida viva a la vez, y siempre con exactamente un nivel
  elegido.
- Al empezar una partida el marcador es 0–0. No se arrastra nada de la
  partida anterior salvo lo que se muestra en pantalla.

## Lenguaje ubicuo

El vocabulario está en `01_contexto.md` y es el mismo en documentación, código
y mensajes de commit. Dos precisiones que sí importan:

- **Anotar** es que la bola cruce el borde lateral, no que rebote.
- **Ganar** es llegar a dos puntos. No existe empate: la partida no termina
  por ninguna otra vía que un marcador en 2.
