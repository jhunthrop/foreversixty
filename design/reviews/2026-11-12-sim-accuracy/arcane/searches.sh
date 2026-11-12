#!/usr/bin/env bash
# searches.sh: the Arcane lane's talent and rotation searches on the lane's
# engine, one at a time, niced to 15, GOMAXPROCS=5 (two other lanes share
# the box). Raid 60 on the raid build and bare 60 on the leveling build; the
# keep lists hold every talent the card casts or leans on (Arcane Precision
# before 40 is held for the owner and is not a 60 question). Run from the
# site's sim directory (tools/remote-test.sh --dir sim --engine-sha <sha>);
# prints every report.
set -euo pipefail
make -C .. simdb >/dev/null
out=$(mktemp -d)
bin=$(mktemp -d)
mkdir -p "$out/talents" "$out/rotation"
go build -o "$bin/talent-search" ./cmd/talent-search
go build -o "$bin/rotation-search" ./cmd/rotation-search
run() { nice --adjustment=15 env GOMAXPROCS=5 "$@" 2>&1 | grep -v '^$' | tail -2; }
run "$bin/talent-search" -repo-root .. -out "$out/talents" -spec mage-arcane -preset raid -keep 'Arcane Blast,Missile Barrage,Presence of Mind,Arcane Power,Arcane Wrath,Pyroblast,Heating Up,Improved Fire Blast,Ignite,Arcane Subtlety,Precision'
run "$bin/talent-search" -repo-root .. -out "$out/talents" -spec mage-arcane -keep 'Arcane Blast,Missile Barrage,Presence of Mind,Arcane Power,Precision'
run "$bin/rotation-search" -repo-root .. -out "$out/rotation" -spec mage-arcane -preset raid
run "$bin/rotation-search" -repo-root .. -out "$out/rotation" -spec mage-arcane
for f in "$out"/*/*.md; do echo "=== $(basename "$(dirname "$f")")/$(basename "$f")"; cat "$f"; done
