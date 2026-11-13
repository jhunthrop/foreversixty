#!/usr/bin/env bash
# run.sh: the raid-searches lane. Re-runs the raid 60 rotation and talent
# searches, on the tools that wear the published enchants (18d046e19), for
# every DPS spec but the rogues, both factions, one search at a time, niced
# to 15, GOMAXPROCS=5 (other lanes share the box). Order follows the live
# tier list's visibility: its top and bottom first. Talent keeps come from
# design/reviews/talent-search/raid-keeps.json. Run from the site's sim
# directory (tools/remote-test.sh --pin-engine --dir sim). Every report is
# printed between "=== FILE <name>" markers so the caller can split them.
#   run.sh [spec ...]
set -euo pipefail
make -C .. simdb >/dev/null
go vet ./cmd/rotation-search/ ./cmd/talent-search/ ./cmd/heal-builds/
go test ./cmd/rotation-search/ ./cmd/talent-search/ ./cmd/heal-builds/
bin=$(mktemp -d)
for c in rotation-search talent-search; do go build -o "$bin/$c" "./cmd/$c"; done
run() { nice --adjustment=15 env GOMAXPROCS=5 "$@"; }
keeps=../design/reviews/talent-search/raid-keeps.json
first="hunter-marksmanship druid-feral warrior-fury hunter-survival mage-frost druid-balance shaman-elemental warlock-affliction"
rest="hunter-beast-mastery warrior-arms paladin-retribution priest-shadow shaman-enhancement mage-arcane mage-fire warlock-demonology warlock-destruction"
specs=${*:-$first $rest}
out=$(mktemp -d)
for s in $specs; do
	keep=$(python3 -c 'import json,sys; print(",".join(json.load(open(sys.argv[1]))["keep"].get(sys.argv[2], [])))' "$keeps" "$s")
	for f in alliance horde; do
		mkdir -p "$out/$f/rotation" "$out/$f/talents"
		run "$bin/rotation-search" -repo-root .. -spec "$s" -faction "$f" -preset raid -out "$out/$f/rotation" 2>&1 | grep 'wrote\|rror' || true
		run "$bin/talent-search" -repo-root .. -spec "$s" -faction "$f" -preset raid -keep "$keep" -out "$out/$f/talents" 2>&1 | grep 'wrote\|rror' || true
		for r in "$out/$f"/rotation/*.md "$out/$f"/talents/*.md; do
			[[ -f $r ]] || continue
			echo "=== FILE $f/$(basename "$(dirname "$r")")/$(basename "$r")"; cat "$r"; rm "$r"
		done
	done
done
