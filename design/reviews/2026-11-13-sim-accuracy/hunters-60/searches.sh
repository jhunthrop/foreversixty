#!/usr/bin/env bash
# searches.sh: talent and rotation searches for the three hunters at 60 on the lane engine,
# one at a time, nice 15, GOMAXPROCS=5. Raid 60 on the raid build (keep lists from
# design/reviews/talent-search/raid-keeps.json), bare 60 on the leveling build. Run from sim/:
#   tools/remote-test.sh --dir sim --engine-sha <sha> HEAD -- bash ../design/reviews/2026-11-13-sim-accuracy/hunters-60/searches.sh [specs...]
set -uo pipefail
specs=${*:-hunter-marksmanship hunter-beast-mastery hunter-survival}
make -C .. simdb >/dev/null || exit 1
out=$(mktemp -d); bin=$(mktemp -d)
go build -o "$bin/ts" ./cmd/talent-search && go build -o "$bin/rs" ./cmd/rotation-search || exit 1
keeps() { python3 -c "import json,sys;print(\",\".join(json.load(open(\"../design/reviews/talent-search/raid-keeps.json\"))[\"keep\"].get(sys.argv[1],[])))" "$1"; }
run() { nice -n 15 env GOMAXPROCS=5 "$@" 2>&1 | grep -v "^$" | tail -2; }
for s in $specs; do
  [ -n "${RS_ONLY:-}" ] && { echo "### rs $s raid"; run "$bin/rs" -repo-root .. -out "$out/rs-raid" -spec "$s" -preset raid; echo "### rs $s bare"; run "$bin/rs" -repo-root .. -out "$out/rs-bare" -spec "$s"; continue; }
  for f in horde alliance; do
    echo "### ts $s raid $f"; run "$bin/ts" -repo-root .. -out "$out/ts-raid-$f" -spec "$s" -preset raid -faction "$f" -keep "$(keeps "$s")"
  done
  echo "### rs $s raid"; run "$bin/rs" -repo-root .. -out "$out/rs-raid" -spec "$s" -preset raid
  echo "### rs $s bare"; run "$bin/rs" -repo-root .. -out "$out/rs-bare" -spec "$s"
done
for f in "$out"/*/*.md; do echo "=== $(basename "$(dirname "$f")")/$(basename "$f")"; sed -n 1,400p "$f"; done
