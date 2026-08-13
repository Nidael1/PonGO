Propósito: decir qué se registra y por qué casi nada.

## Qué se registra

Solo el error fatal de arranque: si Ebitengine no logra abrir la ventana o
inicializar el audio, el programa escribe el error a la salida de error
estándar y termina con código distinto de cero. Nada más.

No se usa Zap ni ninguna librería de logging. Es un programa de escritorio de
un solo usuario que corre en la máquina de quien lo escribió: el canal de
diagnóstico es la terminal desde la que se lanzó y la pantalla misma.

## Métricas

No hay. No existe operación desatendida que medir, ni usuario del que no se
tenga noticia directa. La única medición del proyecto es la calibración de los
niveles (issue #10) y se hace jugando y contando partidas a mano, no
instrumentando.

## Qué NO se registra

Nada del juego: ni posiciones, ni rebotes, ni marcadores, ni niveles elegidos,
ni duración de partidas. No se escribe ningún archivo de log en disco.
