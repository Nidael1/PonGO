Propósito: decir dónde corre esto y qué hay que respaldar.

## Entornos

Uno solo: la Mac de quien lo desarrolla. No hay entorno de desarrollo separado
del de producción porque no hay producción. El programa se compila y se corre
en la misma máquina en la que se escribe.

## Variables de configuración

No hay. Ni variables de entorno, ni archivo de configuración, ni Viper. Todo lo
ajustable del juego —tamaños, rapidez, valores de los niveles— son constantes
en el código, y se cambian recompilando. Es lo correcto aquí: el único que las
ajusta es quien tiene el código abierto.

## Despliegue

No hay. No se sube a Fly.io ni a DigitalOcean, no hay Docker, no hay GitHub
Actions. "Desplegar" es `go run ./cmd/pongo` en la máquina local; ver
`12_ejecucion.md`.

El binario tampoco se distribuye: no se firma, no se notariza y no se empaqueta
como `.app`. Eso está declarado como frontera en `01_contexto.md`.

## Respaldos

El repositorio en GitHub es el respaldo: `https://github.com/Nidael1/PonGO`.
No hay estado que respaldar fuera del código, porque el juego no persiste nada.
Cada commit es un punto seguro de retorno; esa es la razón de la convención de
un commit por tarea terminada.
