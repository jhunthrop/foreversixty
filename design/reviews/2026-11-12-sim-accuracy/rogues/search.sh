#!/usr/bin/env bash
# search.sh: the rogue lane's talent searches on the current engine, one at a
# time, niced to 15, GOMAXPROCS=5 (two other lanes share the box). Raid 60 on
# each spec's raid build, bare 60 on the leveling build; the keep lists hold
# every talent the card casts or leans on. Run from the site's sim directory
# (tools/remote-test.sh --dir sim); prints every report.
set -euo pipefail
make -C .. simdb >/dev/null
out=$(mktemp -d)
bin=$(mktemp -d)/talent-search
go build -o "$bin" ./cmd/talent-search
search() {
	nice --adjustment=15 env GOMAXPROCS=5 "$bin" -repo-root .. -out "$out" "$@" 2>&1 | grep -v '^$' | tail -2
}
search -spec rogue-assassination -preset raid -keep 'Cold Blood,Mutilate,Venom,Blood Splatter,Seal Fate'
search -spec rogue-combat -preset raid -keep 'Adrenaline Rush,Blade Flurry'
search -spec rogue-subtlety -preset raid -keep 'Ghostly Strike,Premeditation,Hemorrhage,Preparation,Death Mark,Boot Knife'
search -spec rogue-assassination -keep 'Cold Blood,Mutilate,Venom,Seal Fate'
search -spec rogue-combat -keep 'Adrenaline Rush,Blade Flurry'
search -spec rogue-subtlety -keep 'Ghostly Strike,Premeditation,Hemorrhage,Preparation,Death Mark,Boot Knife'
for f in "$out"/*.md; do echo "=== $(basename "$f")"; cat "$f"; done
