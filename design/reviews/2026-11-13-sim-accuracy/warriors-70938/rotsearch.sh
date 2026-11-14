#!/usr/bin/env bash
# rotsearch.sh: Arms and Fury rotation searches (raid 60 and bare 60) on the engine this checkout is pinned to,
# printing each report's probe rows whole. Run from the site's sim directory:
#   tools/remote-test.sh --engine-sha <fork sha> --dir sim HEAD -- bash ../design/reviews/2026-11-13-sim-accuracy/warriors-70938/rotsearch.sh
set -uo pipefail
make -C .. simdb >/dev/null || exit 1
out=$(mktemp -d); bin=$(mktemp -d)
go build -o "$bin/rotation-search" ./cmd/rotation-search || exit 1
for spec in warrior-arms warrior-fury; do
	nice --adjustment=15 env GOMAXPROCS=5 "$bin/rotation-search" -repo-root .. -out "$out" -spec "$spec" -preset raid >/dev/null 2>&1
	nice --adjustment=15 env GOMAXPROCS=5 "$bin/rotation-search" -repo-root .. -out "$out" -spec "$spec" >/dev/null 2>&1
done
for f in "$out"/*.md; do echo "########## $(basename "$f")"; grep -E "^[|] (remove|insert|move|swap|Kind)|Baseline" "$f"; done
rm -rf "$out" "$bin"
