#!/usr/bin/env bash
# tools/dev-stack/up.sh [--foreground]
#
# Starts the local guild stack from WSL: a throwaway Postgres on 127.0.0.1:5440, then the
# API on 127.0.0.1:8080 with its migrations applied, the addon data served from
# data/builds, and the dev sign-in (POST /v1/dev/login) turned on. Safe to run twice:
# a running Postgres or API is left alone.
#
# By default the API runs in the background and this script exits once /health answers.
# --foreground keeps the script alive (it follows the API log), which is what the Browser
# pane's `api-local` launch configuration needs so it can stop the stack by stopping the
# process.
set -euo pipefail
. "$(dirname "${BASH_SOURCE[0]}")/lib.sh"

foreground=0
[ "${1:-}" = "--foreground" ] && foreground=1

if [ -z "$PGBIN" ] || [ ! -x "$PGBIN/initdb" ]; then
  echo "up.sh: no PostgreSQL server binaries found (looked in /usr/lib/postgresql/*/bin)." >&2
  exit 1
fi

# 1. Postgres.
if [ ! -f "$PGDATA_DIR/PG_VERSION" ]; then
  echo "up.sh: creating the throwaway Postgres cluster in $PGDATA_DIR"
  "$PGBIN/initdb" -D "$PGDATA_DIR" -U postgres --auth=trust -E UTF8 >"$LOG_DIR/initdb.log"
fi
if ! pg_up; then
  echo "up.sh: starting Postgres on 127.0.0.1:$PG_PORT"
  "$PGBIN/pg_ctl" -D "$PGDATA_DIR" -l "$LOG_DIR/postgres.log" -w \
    -o "-p $PG_PORT -k $FS_DEV_STATE -c listen_addresses=127.0.0.1 -c fsync=off" start >/dev/null
fi
psql_admin() { "$PGBIN/psql" -h 127.0.0.1 -p "$PG_PORT" -U postgres -v ON_ERROR_STOP=1 -qtA "$@"; }
if [ "$(psql_admin -d postgres -c "select 1 from pg_roles where rolname = 'forever'")" != "1" ]; then
  psql_admin -d postgres -c "create role forever login password 'forever' superuser"
fi
if [ "$(psql_admin -d postgres -c "select 1 from pg_database where datname = 'forever_dev'")" != "1" ]; then
  psql_admin -d postgres -c "create database forever_dev owner forever"
fi

# 2. The API. Its startup applies every migration to forever_dev.
if api_healthy; then
  echo "up.sh: the API is already answering on 127.0.0.1:$API_PORT"
else
  echo "up.sh: building the API and the seed tool"
  (cd "$REPO_ROOT/api" && go build -o "$FS_DEV_STATE/api" ./cmd/api && go build -o "$FS_DEV_STATE/seedguild" ./cmd/seedguild)
  echo "up.sh: starting the API on 127.0.0.1:$API_PORT (log: tools/dev-stack/logs/api.log)"
  # FS_DEV_LOGIN needs LISTEN_HOST to be a loopback address, or the API refuses to start
  # (api/internal/config). The mail key is a placeholder: no mail is sent from here.
  # Battle.net and Stripe are left unset on purpose; the walkthrough doc says what that
  # means for sign-in and billing.
  env DATABASE_URL="$DB_URL" RESEND_API_KEY=dev-unused \
    PUBLIC_BASE_URL="$WEB_ORIGIN" API_BASE_URL="$API_ORIGIN" \
    TREE_DATA_DIR="$REPO_ROOT/data/builds" SESSION_COOKIE_DOMAIN=none TRUSTED_PROXY_HOPS=0 \
    PORT="$API_PORT" LISTEN_HOST=127.0.0.1 FS_DEV_LOGIN=1 \
    setsid nohup "$FS_DEV_STATE/api" >>"$LOG_DIR/api.log" 2>&1 </dev/null &
  echo $! >"$FS_DEV_STATE/api.pid"
  for _ in $(seq 1 90); do
    api_healthy && break
    sleep 1
  done
  if ! api_healthy; then
    echo "up.sh: the API did not answer /health within 90 s; last log lines:" >&2
    tail -n 20 "$LOG_DIR/api.log" >&2
    exit 1
  fi
fi
echo "up.sh: ready. API http://127.0.0.1:$API_PORT, Postgres 127.0.0.1:$PG_PORT/forever_dev"

if [ "$foreground" = 1 ]; then
  exec tail -n 0 -F "$LOG_DIR/api.log"
fi
