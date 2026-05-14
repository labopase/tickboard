#!/bin/sh
set -eu

PROJECT_ROOT="$(cd "$(dirname "$0")/.." && pwd)"
COMPOSE_FILE="$PROJECT_ROOT/infrastructures/docker/docker-compose.yml"
ENV_FILE="$PROJECT_ROOT/infrastructures/docker/.env"

if [ $# -eq 0 ]; then
  set -- up -d
fi

docker compose -f "$COMPOSE_FILE" --env-file "$ENV_FILE" "$@"