#!/bin/sh
set -e

# Run migrations if enabled (default: true)
if [ "${RUN_MIGRATIONS:-true}" = "true" ]; then
  echo "==> Running database migrations..."
  /app/migrate_bin -up
fi

# Run seeds if enabled (default: false)
if [ "${RUN_SEEDS:-false}" = "true" ]; then
  echo "==> Running database seeds..."
  /app/migrate_bin -seed
fi

echo "==> Starting Hotel Management API..."
exec "$@"
