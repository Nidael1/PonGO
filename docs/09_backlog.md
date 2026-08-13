Propósito: guardar lo que no está en el sprint actual y el orden en que se construye.

## Pendientes

- [ ] #11 corregir ángulos de rebote demasiado planos, si aparecen al jugar
- [ ] #12 afinar timbre y duración de los tres sonidos
- [ ] #13 pausa con la barra espaciadora
- [ ] #14 tabla del último ganador en pantalla de selección

Ninguno de los tres bloquea el sprint 1.0. El #13 no está comprometido: es una
idea, y solo entra a un sprint si el director lo decide.

## Bugs

No hay bugs registrados. No hay código todavía.

## Orden y dependencias

1. **Arranque** — issues #1, #2, #3. Sin dependencias.
2. **Movimiento** — issues #4, #5. Depende de 1: no hay nada que mover sin
   ventana ni bucle.
3. **Rival** — issue #6. Depende de 2: el rival reacciona a la bola.
4. **Partida** — issues #7, #8. Depende de 3: no hay punto sin dos paletas.
5. **Sonido** — issue #9. Depende de 4: los tres sonidos se disparan en
   eventos que solo existen con partida completa.
6. **Calibración** — issue #10. Depende de 5, o mejor dicho de todo lo
   anterior: se calibra jugando el juego terminado.
7. **Ajustes** — issues #11, #12. Dependen de 6: se detectan calibrando.
