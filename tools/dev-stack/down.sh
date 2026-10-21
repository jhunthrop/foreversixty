#!/usr/bin/env bash
# tools/dev-stack/down.sh [--wipe]
#
# Stops the API and Postgres. --wipe also deletes the throwaway cluster and the built
# binaries, so the next up.sh and seed.sh start from nothing. Safe to run when nothing is
# running.
set -euo pipefail
. "$(dirname "${BASH_SOURCE[0]}")/lib.sh"

if [ -f "$FS_DEV_STATE/api.pid" ]; then
  pid="$(cat "$FS_DEV_STATE/api.pid")"
  if kill -0 "$pid" 2>/dev/null; then
    kill "$pid"
    for _ in $(seq 1 20); do
      kill -0 "$pid" 2>/dev/null || break
      sleep 0.5
    done
    if kill -0 "$pid" 2>/dev/null; then kill -9 "$pid"; fi
  fi
  rm -f "$FS_DEV_STATE/api.pid"
fi
# An API started from an earlier shell whose pid file is gone.
pkill -f "$FS_DEV_STATE/api" 2>/dev/null || true

if [ -f "$PGDATA_DIR/PG_VERSION" ] && [ -n "$PGBIN" ]; then
  "$PGBIN/pg_ctl" -D "$PGDATA_DIR" -m fast -w stop >/dev/null 2>&1 || true
fi
if [ "${1:-}" = "--wipe" ]; then
  rm -rf "$FS_DEV_STATE"
  echo "down.sh: stopped and wiped $FS_DEV_STATE"
else
  echo "down.sh: stopped (data kept in $FS_DEV_STATE; --wipe removes it)"
fi
