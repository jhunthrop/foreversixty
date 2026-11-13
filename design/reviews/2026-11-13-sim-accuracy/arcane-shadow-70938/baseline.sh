#!/usr/bin/env bash
# baseline.sh: the published Arcane and Shadow entries at bare 50, bare 60 and raid 60, both factions,
# 10,000 iterations, seed 1, one sim at a time, niced to 15, GOMAXPROCS=5. Run on the box from sim/:
#   [SPECS="..."] [RAID_ONLY=1] tools/remote-test.sh --pin-engine [--engine-sha <sha>] --dir sim HEAD -- bash ../design/reviews/2026-11-13-sim-accuracy/arcane-shadow-70938/baseline.sh [extra spec-breakdown flags]
set -euo pipefail
echo "active build: $(cat ../web/src/data/active-build.json | tr -d "\n"); engine $(sed -n 's/.*Version = "\(.*\)"/\1/p' enginever/version.go); $(grep "replace github.com/wowsims/classic" go.mod | tail -1); simdb sha $(sha256sum ../data/builds/$(sed -n 's/.*"build"[^"]*"\([^"]*\)".*/\1/p' ../web/src/data/active-build.json)/simdb.bin | cut -c1-12)"
make -C .. simdb >/dev/null
bin=$(mktemp -d)
go build -o "$bin/spec-breakdown" ./cmd/spec-breakdown
run() {
	echo "=== $*"
	nice --adjustment=15 env GOMAXPROCS=5 "$bin/spec-breakdown" -iterations "${ITER:-10000}" -seed "${SEED:-1}" "$@"
}
for spec in ${SPECS:-mage-arcane priest-shadow}; do
	for f in alliance horde; do
		run -spec $spec -preset raid -band 60 -faction $f "${@}"
		[ -n "${RAID_ONLY:-}" ] && continue
		run -spec $spec -preset bare -band 60 -faction $f "${@}"
		run -spec $spec -preset bare -band 50 -faction $f "${@}"
	done
done
