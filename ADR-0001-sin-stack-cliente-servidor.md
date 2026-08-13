# ADR-0001 — No aplicar el stack estándar cliente-servidor
Estado: aceptada
Fecha: 2026-08-13

## Contexto

El estándar de la fábrica supone aplicaciones cliente-servidor: arquitectura
hexagonal, PostgreSQL con pgx y sqlc, Gin para HTTP, JWT para autenticación,
Zap, Viper, Docker, GitHub Actions y la librería interna `go-lib`.

PonGO no es una aplicación cliente-servidor. Es un juego de escritorio que
corre en una sola máquina, sin red, sin usuarios, sin persistencia y sin nada
que servir. Aplicar el estándar tal cual significaría arrastrar puertos,
adaptadores, un pool de base de datos y un router HTTP para un programa cuyo
estado completo cabe en unas cuantas variables y muere al cerrar la ventana.

Además, el objetivo real del proyecto es comprobar que la instalación local de
Go compila y ejecuta un programa gráfico. Cualquier capa que se interponga
entre `go run` y una ventana abierta trabaja en contra de ese objetivo.

## Opciones consideradas

1. **Aplicar el estándar completo.** Hexagonal con el juego como dominio y
   Ebitengine como adaptador de entrada/salida. Coherente con el resto de la
   fábrica, pero el costo estructural no compra nada: no hay motor de base de
   datos que cambiar ni MVP que escalar, que son las dos razones por las que el
   estándar exige hexagonal.
2. **Ebitengine con separación mínima por responsabilidad.** Cinco paquetes con
   dependencias en un solo sentido, sin puertos ni adaptadores.
3. **Go sin librería gráfica, en terminal.** Menos dependencias todavía, pero
   no ejercita render ni audio, que es justo lo que se quiere comprobar.

## Decisión

Opción 2. Go con Ebitengine (`github.com/hajimehoshi/ebiten/v2`) y librería
estándar, nada más.

Del estándar de la fábrica **no aplican**: arquitectura hexagonal, PostgreSQL,
pgx, sqlc, Gin, JWT, Zap, Viper, Docker, GitHub Actions, `go-lib` y el
despliegue en Fly.io o DigitalOcean. Cada una está justificada en su archivo
correspondiente: `06_observabilidad.md`, `10_operativos.md` y
`11_seguridad.md`.

Del estándar **sí se conservan**: el lenguaje Go, la convención de commits
`[sprint N.N][issue #<n>]`, los minisprints, la lista completa de archivos del
esquema y la regla de autonomía de la documentación.

## Consecuencias

- El proyecto no comparte infraestructura con el resto de la fábrica. No hereda
  mejoras de `go-lib` ni las devuelve.
- La separación en cinco paquetes es una convención de este repo, no el
  estándar. Quien venga del resto de los proyectos no encontrará puertos ni
  adaptadores y no debe agregarlos.
- Si algún día PonGO necesitara red —marcadores compartidos, segundo jugador
  remoto— esta decisión se reemplaza con un ADR nuevo. Hoy eso está declarado
  como frontera en `01_contexto.md`.
- Ebitengine queda como dependencia única y como punto de falla único. Si
  dejara de compilar en una versión futura de macOS, el proyecto se detiene.
  Es aceptable: es un proyecto de práctica.
