Propósito: fijar las convenciones de trabajo que se aplican en cada commit.

## Commits

Uno por actividad o tarea terminada, nunca por día. El commit es el punto
seguro de respaldo.

Formato obligatorio:

```
[sprint N.N][issue #<n>] <descripción>
```

Ejemplo: `[sprint 1.0][issue #4] paleta del jugador con flechas`

- `N.N` es el minisprint en curso, el que declara `07_sprint_actual.md`.
- `#<n>` es el número de issue de la tarea en el checklist. No se inventa: si
  la tarea no está en el checklist ni en el backlog, primero se agrega ahí.
- La descripción va en minúsculas, en presente, sin punto final.

Un commit cierra una tarea del checklist. Al hacerlo se marca `- [x]` en
`07_sprint_actual.md` en ese mismo commit.

## Ramas

`main` directo. Un solo desarrollador, sin revisión de terceros y sin
despliegue: ramas de trabajo no aportan nada y sí estorban.

Si en algún momento entra un colaborador, la convención cambia a una rama por
issue con nombre `issue-<n>-<slug>` y eso se registra aquí.

## Naming

- Paquetes en minúscula, una palabra, sin guiones ni guiones bajos.
- Nombres de archivo en minúscula con guion bajo si hace falta.
- Los nombres del código usan el vocabulario de `01_contexto.md` en español
  cuando nombran cosas del dominio: `Paleta`, `Bola`, `Marcador`, `Nivel`.
  Lo demás sigue la convención habitual de Go.
- Nada de abreviaturas inventadas.

## Estructura de carpetas

```
/
├── cmd/pongo/main.go      arranque: crea el juego y lo corre
├── internal/juego/        estado, bucle, marcador
├── internal/entrada/      teclado
├── internal/rival/        paleta de la máquina
├── internal/render/       dibujo
├── internal/audio/        sonido sintetizado
├── go.mod
├── MANIFEST.md            y el resto de la documentación raíz
└── ADR-000N-*.md
```

La documentación vive en la raíz del repo, no en `docs/`. Es lo primero que se
ve al abrir el proyecto y eso es intencional.
