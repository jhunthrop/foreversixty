#!/usr/bin/env bash
# bands.sh: every published entry Ignite reaches (Fire at every band, Arcane's raid 60 build, which takes
# Ignite 5/5), 10,000 iterations, seed 1, one at a time, niced to 15, GOMAXPROCS=5. Run from the site's sim
# directory on the box, once per engine:
#   tools/remote-test.sh --engine-sha <sha> --dir sim HEAD -- bash ../design/reviews/2026-11-13-sim-accuracy/crit-flags/bands.sh
set -euo pipefail
make -C .. simdb >/dev/null
bin=$(mktemp -d)
go build -o "$bin/spec-breakdown" ./cmd/spec-breakdown
run() {
	echo "=== $*"
	nice --adjustment=15 env GOMAXPROCS=5 "$bin/spec-breakdown" -iterations 10000 -seed 1 "$@" > "$bin/out.md"
	sed -n 1,12p "$bin/out.md"; grep -i "ignite" "$bin/out.md" || true
}
for f in horde alliance; do run -spec mage-fire -preset raid -band 60 -faction $f; done
for b in 60 50 40 30 20; do run -spec mage-fire -preset bare -band $b -faction horde; done
run -spec mage-fire -preset bare -band 60 -faction alliance
for f in horde alliance; do run -spec mage-arcane -preset raid -band 60 -faction $f; done
