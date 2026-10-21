# tools/dev-stack/lib.sh
#
# Shared settings for the local guild stack. Sourced by up.sh, seed.sh, smoke.sh and
# down.sh; it sets no shell options of its own (the scripts that source it do).
#
# The stack is throwaway on purpose: its Postgres cluster and built binaries live in the
# WSL filesystem under FS_DEV_STATE (not on /mnt/c, where Postgres cannot keep file
# permissions and builds are slow), and `down.sh --wipe` deletes all of it.

STACK_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "$STACK_DIR/../.." && pwd)"
LOG_DIR="$STACK_DIR/logs"

FS_DEV_STATE="${FS_DEV_STATE:-$HOME/.cache/fs-dev-stack}"
PGDATA_DIR="$FS_DEV_STATE/pgdata"
PGBIN="${PGBIN:-$(ls -d /usr/lib/postgresql/*/bin 2>/dev/null | sort -V | tail -1)}"
PG_PORT="${FS_DEV_PG_PORT:-5440}"
API_PORT="${FS_DEV_API_PORT:-8080}"
WEB_PORT="${FS_DEV_WEB_PORT:-4321}"

# The database the API and the seed share. The credentials are throwaway and the cluster
# listens on loopback only.
DB_URL="postgres://forever:forever@127.0.0.1:${PG_PORT}/forever_dev?sslmode=disable"

# The browser reaches the site and the API as "localhost" on both ports, so the session
# cookie (host-only, SameSite=Lax) is same-site between them. 127.0.0.1 for one and
# localhost for the other would be cross-site and the cookie would not be sent.
WEB_ORIGIN="http://localhost:${WEB_PORT}"
API_ORIGIN="http://localhost:${API_PORT}"

# The seeded sign-ins (api/cmd/seedguild/lifecycle.go).
GM_BATTLETAG="hunthrop#1894"
OFFICER_BATTLETAG="Ashenveil#2201"
MEMBER_BATTLETAG="Quillfeather#3301"
VISITOR_BATTLETAG="Wanderer#4401"

mkdir -p "$LOG_DIR" "$FS_DEV_STATE"

api_healthy() {
  curl -fs -o /dev/null --max-time 2 "http://127.0.0.1:${API_PORT}/health"
}

pg_up() {
  "$PGBIN/pg_isready" -q -h 127.0.0.1 -p "$PG_PORT"
}
