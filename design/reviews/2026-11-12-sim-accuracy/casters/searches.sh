#!/usr/bin/env bash
# searches.sh: the Fire, Frost and Balance lane's talent and rotation searches
# on the pinned engine, one at a time, niced to 15, GOMAXPROCS=5 (two other
# sim lanes share the box). Raid 60 on each guide's raidBuild (rotation-search
# reads it since the Arcane lane's fix) and bare 60 on the leveling build. The
# keep lists hold every modelled talent the card casts or leans on. Run from
# the site's sim directory (tools/remote-test.sh --dir sim); prints every
# report.
set -euo pipefail
make -C .. simdb >/dev/null
out=$(mktemp -d)
bin=$(mktemp -d)
mkdir -p "$out/talents" "$out/rotation"
go build -o "$bin/talent-search" ./cmd/talent-search
go build -o "$bin/rotation-search" ./cmd/rotation-search
run() { nice --adjustment=15 env GOMAXPROCS=5 "$@" 2>&1 | grep -v '^$' | tail -2; }
for spec in mage-fire mage-frost druid-balance; do
	case $spec in
	mage-fire) keep='Pyroblast,Heating Up,Combustion,Improved Scorch,Ignite,Improved Fire Blast,Arcane Blast,Missile Barrage' ;;
	mage-frost) keep="Ice Lance,Fingers of Frost,Ice Barrier,Missile Barrage,Winter's Chill,Cold Snap" ;;
	druid-balance) keep="Insect Swarm,Eclipse,Moonkin Form,Nature's Swiftness" ;;
	esac
	run "$bin/rotation-search" -repo-root .. -out "$out/rotation" -spec "$spec" -preset raid
	run "$bin/rotation-search" -repo-root .. -out "$out/rotation" -spec "$spec"
	run "$bin/talent-search" -repo-root .. -out "$out/talents" -spec "$spec" -preset raid -keep "$keep"
	run "$bin/talent-search" -repo-root .. -out "$out/talents" -spec "$spec" -keep "$keep"
done
for f in "$out"/*/*.md; do echo "=== $(basename "$(dirname "$f")")/$(basename "$f")"; cat "$f"; done
