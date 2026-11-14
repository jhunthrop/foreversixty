#!/usr/bin/env bash
# searches.sh: Arms and Fury rotation and talent searches on the Windfury and Rend engine, one at a time, niced to 15,
# GOMAXPROCS=5. Raid 60 on each guide's raidBuild and bare 60 on the leveling build, both factions for the talent
# searches (the keep lists are design/reviews/talent-search/raid-keeps.json's). Run from the site's sim directory:
#   tools/remote-test.sh --engine-sha <fork sha> --dir sim HEAD -- bash ../design/reviews/2026-11-13-sim-accuracy/warriors-70938/searches.sh
set -uo pipefail
make -C .. simdb >/dev/null || exit 1
out=$(mktemp -d); bin=$(mktemp -d)
mkdir -p "$out/talents" "$out/rotation"
go build -o "$bin/talent-search" ./cmd/talent-search && go build -o "$bin/rotation-search" ./cmd/rotation-search || exit 1
keeps() { python3 -c "import json,sys;print(','.join(json.load(open('../design/reviews/talent-search/raid-keeps.json'))['keep'].get(sys.argv[1],[])))" "$1"; }
run() { echo "### $*"; nice -n 15 env GOMAXPROCS=5 "$@" 2>&1 | grep -v '^$' | tail -2; }
for spec in warrior-arms warrior-fury; do
	run "$bin/rotation-search" -repo-root .. -out "$out/rotation" -spec "$spec" -preset raid
	run "$bin/rotation-search" -repo-root .. -out "$out/rotation" -spec "$spec"
	for f in horde alliance; do
		run "$bin/talent-search" -repo-root .. -out "$out/talents/raid-$f" -spec "$spec" -preset raid -faction "$f" -keep "$(keeps $spec)"
		run "$bin/talent-search" -repo-root .. -out "$out/talents/bare-$f" -spec "$spec" -preset bare -faction "$f"
	done
done
for f in "$out"/*/*.md "$out"/*/*/*.md; do [ -f "$f" ] && { echo "########## ${f#$out/}"; sed -n 1,30p "$f"; }; done
rm -rf "$out" "$bin"
