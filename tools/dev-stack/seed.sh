#!/usr/bin/env bash
# tools/dev-stack/seed.sh
#
# Seeds the walkthrough data through api/cmd/seedguild --lifecycle: three accounts, the
# unclaimed guild "Lifecycle Test", its roster as the addon would have sent it, and one
# complete report. Starts the stack first if it is not up. Running it twice is a no-op the
# second time (seedguild checks for the guild before writing anything). `seed.sh --remove`
# takes the seeded rows out again.
set -euo pipefail
. "$(dirname "${BASH_SOURCE[0]}")/lib.sh"

"$STACK_DIR/up.sh"
# The tool reads the active build's BiS tables relative to the repository, so it runs
# from the repository's api directory. It is rebuilt every time (a cached build is
# instant) so an edit to the seed is never run from a stale binary.
cd "$REPO_ROOT/api"
go build -o "$FS_DEV_STATE/seedguild" ./cmd/seedguild
DATABASE_URL="$DB_URL" "$FS_DEV_STATE/seedguild" --lifecycle "$@"
