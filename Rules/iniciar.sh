#!/usr/bin/env bash
# Arranque de PonGO: inicializa el repositorio y sube el esquema raíz.
# Se detiene ante cualquier error y no fuerza nada.
set -e

REPO="https://github.com/Nidael1/PonGO.git"
COMMIT="[sprint 1.0][issue #1] esquema raiz"

cd "$(dirname "$0")"

if [ -d .git ]; then
  echo "Esta carpeta ya es un repositorio git. No se vuelve a inicializar."
else
  git init -b main
fi

git add -A

if git diff --cached --quiet; then
  echo "No hay cambios que confirmar."
else
  git commit -m "$COMMIT"
fi

if git remote get-url origin >/dev/null 2>&1; then
  echo "El remoto 'origin' ya existe: $(git remote get-url origin)"
else
  git remote add origin "$REPO"
fi

git push -u origin main
echo "Listo."
