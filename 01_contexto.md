Propósito: definir qué es PonGO, para quién se construye y hasta dónde llega.

## Qué resuelve y para quién

PonGO es un Pong clásico de escritorio para macOS, escrito en Go con
Ebitengine. Lo construye y lo juega una sola persona.

El objetivo real no es el juego: es comprobar que la instalación local de Go
compila y ejecuta un programa gráfico completo, con ventana, entrada de
teclado, audio y bucle de render. El Pong es la excusa mínima para ejercitar
esas cuatro cosas de punta a punta.

Se juega contra la máquina. La partida es corta a propósito —dos puntos— para
poder llegar al final y volver a empezar en menos de un minuto.

## Perfil y por qué

Perfil 2 — práctica, no monetizable.

Nadie paga por PonGO y no hay intención de venderlo. Sin cliente y sin ingreso
que proteger, se prioriza estabilidad y solución simple por encima de la
solución ideal. La documentación registra únicamente lo que no es evidente
para alguien que ya conoce el tema; todo lo genérico de un Pong se declara
genérico y no se explica.

## Fronteras — qué NO hace este proyecto

- No corre en Windows ni en Linux. Solo macOS.
- No hay red: no hay servidor, ni partidas remotas, ni segundo jugador humano.
- No hay archivos de audio. El sonido se sintetiza en código.
- No hay recursos externos de ningún tipo: el binario no depende de archivos
  sueltos junto a él.
- No hay persistencia. El marcador vive en memoria y muere al cerrar.
- No hay menú de configuración, ni opciones gráficas, ni remapeo de teclas.
- No se distribuye a otras máquinas ni se firma ni se empaqueta como `.app`.
- No hay base de datos, ni API HTTP, ni autenticación. Ver
  `ADR-0001-sin-stack-cliente-servidor.md`.

## Alcance funcional

- Como jugador, quiero mover mi paleta con las flechas arriba y abajo, para
  devolver la bola.
- Como jugador, quiero enfrentar una paleta controlada por la máquina, para
  tener contra quién jugar sin segunda persona.
- Como jugador, quiero elegir entre tres niveles de dificultad antes de jugar,
  para ajustar el reto.
- Como jugador, quiero ver el marcador durante la partida, para saber cómo voy.
- Como jugador, quiero que la partida termine al llegar alguien a dos puntos,
  para que una ronda dure poco.
- Como jugador, quiero escuchar un sonido al rebotar la bola, al anotarse un
  punto y al terminar la partida, para tener respuesta audible.
- Como jugador, quiero que al terminar la partida aparezca la pantalla de
  selección de nivel con el marcador final, para volver a jugar sin reiniciar
  el programa.

## Vocabulario interno

| Término | Significa |
|---|---|
| Jugador | La paleta izquierda, controlada con las flechas. |
| Rival | La paleta derecha, controlada por el programa. |
| Punto | Lo que gana quien no dejó pasar la bola por su lado. |
| Partida | Serie que termina cuando alguien llega a dos puntos. |
| Nivel | Uno de tres ajustes de dificultad del rival. Se elige antes de cada partida. |
| Selección | Pantalla de elección de nivel. Es también la pantalla de fin de partida. |
