#!/usr/bin/env bash
# searches2.sh: the second talent-search pass. The first pass's Fire finalists
# all re-spent the guide's threat points (two of Burning Soul, one of Arcane
# Subtlety), which the guide keeps on purpose, and the Balance raid finalists
# re-spent Subtlety (threat) and Nature's Focus (pushback on a boss that hits
# the raid); this pass keeps them and asks what else beats the guides. One
# sim at a time, niced to 15, GOMAXPROCS=5. Run from the site's sim
# directory (tools/remote-test.sh --dir sim --pin-engine).
set -euo pipefail
make -C .. simdb >/dev/null
out=$(mktemp -d)
bin=$(mktemp -d)
mkdir -p "$out/talents"
go build -o "$bin/talent-search" ./cmd/talent-search
run() { nice --adjustment=15 env GOMAXPROCS=5 "$@" 2>&1 | grep -v '^$' | tail -2; }
fire='Pyroblast,Heating Up,Combustion,Improved Scorch,Ignite,Improved Fire Blast,Arcane Blast,Missile Barrage,Burning Soul,Arcane Subtlety'
run "$bin/talent-search" -repo-root .. -out "$out/talents" -spec mage-fire -preset raid -keep "$fire"
run "$bin/talent-search" -repo-root .. -out "$out/talents" -spec mage-fire -keep "$fire"
run "$bin/talent-search" -repo-root .. -out "$out/talents" -spec druid-balance -preset raid -keep "Insect Swarm,Eclipse,Moonkin Form,Nature's Swiftness,Subtlety,Nature's Focus"
run "$bin/talent-search" -repo-root .. -out "$out/talents" -spec druid-balance -keep "Insect Swarm,Eclipse,Moonkin Form,Nature's Swiftness,Subtlety"
for f in "$out"/*/*.md; do echo "=== $(basename "$(dirname "$f")")/$(basename "$f")"; cat "$f"; done
