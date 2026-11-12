#!/usr/bin/env bash
# baseline.sh: the shamans-60 lane's spec breakdowns at the pin. Enhancement
# and Elemental at raid 60 (both factions), bare 60 and bare 50, 10,000
# iterations, seed 1, one sim at a time, niced to 15, GOMAXPROCS=5. Run from
# the site's sim directory (tools/remote-test.sh --dir sim); prints a JSON
# line per cell plus the action table.
set -euo pipefail
make -C .. simdb >/dev/null
bin=$(mktemp -d)/spec-breakdown
go build -o "$bin" ./cmd/spec-breakdown
run() { # spec preset band faction [extra...]
	local spec=$1 preset=$2 band=$3 faction=$4; shift 4
	echo "### $spec $preset $band $faction $*"
	nice --adjustment=15 env GOMAXPROCS=5 "$bin" -repo-root .. -spec "$spec" -preset "$preset" -band "$band" -faction "$faction" -seed 1 -iterations 10000 "$@" 2>&1 | grep -vE '^\s*$'
}
for spec in shaman-enhancement shaman-elemental; do
	run "$spec" raid 60 horde
	run "$spec" raid 60 alliance
	run "$spec" bare 60 horde
	run "$spec" bare 50 horde
done
