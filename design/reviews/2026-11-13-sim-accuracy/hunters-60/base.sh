#!/usr/bin/env bash
# base.sh: full spec-breakdown output for the three hunters at 50 and 60, bare and raid, both factions.
# Run from sim/ on the box: tools/remote-test.sh --dir sim [--engine-sha <sha>] HEAD -- bash ../design/reviews/2026-11-13-sim-accuracy/hunters-60/base.sh [iterations] [specs...]
set -u
iters=${1:-10000}; shift || true
specs=${*:-hunter-beast-mastery hunter-marksmanship hunter-survival}
make -C .. simdb >/dev/null || exit 1
bin=$(mktemp -d)/sb; go build -o "$bin" ./cmd/spec-breakdown || exit 1
for s in $specs; do for f in horde alliance; do
  for pb in "raid 60" "bare 60" "bare 50"; do set -- $pb
    echo "######## $s $1 $2 $f"
    nice -n 15 env GOMAXPROCS=5 "$bin" -repo-root .. -iterations "$iters" -seed "${SEED:-1}" -spec "$s" -preset "$1" -band "$2" -faction "$f" 2>&1
  done
done; done
