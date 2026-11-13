#!/usr/bin/env bash
# jobs3.sh: the searches' two leads at 10,000 iterations, one sim at a time, nice 15,
# GOMAXPROCS=5. Run from the site's sim directory on the box:
#   tools/remote-test.sh --dir sim --engine-sha <sha> HEAD -- bash ../design/reviews/2026-11-13-sim-accuracy/enhancement-70938/jobs3.sh
#   A. The Flame Shock line after Stormstrike gated on mana (rot-fs-after-ss-mana*.json)
#      against the card and the ungated line: raid and bare 60 seeds 1 and 50001,
#      bare 30 to 50 seed 1, both factions.
#   B. The raid talent winner (Mindfulness 3 into Reverberation 2 and Ancestral Knowledge 1)
#      against the raid build, both wearing a Flame Shock line: raid 60, seeds 1 and 50001.
set -uo pipefail
R=../design/reviews/2026-11-13-sim-accuracy/enhancement-70938
make -C .. simdb >/dev/null || exit 1
bin=$(mktemp -d)/spec-breakdown
go build -o "$bin" ./cmd/spec-breakdown || exit 1
sb() { # label, then spec-breakdown args
	local label=$1; shift
	printf '%s\t' "$label"
	nice -n 15 env GOMAXPROCS=5 "$bin" -repo-root .. -spec shaman-enhancement -iterations 10000 -json "$@" 2>/dev/null |
		python3 -c 'import json,sys; d=json.load(sys.stdin); print("%.1f\t%.1f\t%.1f" % (d["DPS"], d["Error"], d.get("SecondsOOM", 0)))' || echo FAILED
}
echo "## A. Flame Shock"
for v in fs-after-ss-mana30 fs-after-ss-mana50 fs-after-ss-mana70; do
	rot=(); [[ $v != card ]] && rot=(-rotation "$R/rot-$v.json")
	for faction in alliance horde; do
		for seed in 1 50001; do
			sb "$v raid 60 $faction $seed" -preset raid -band 60 -faction $faction -seed $seed "${rot[@]}"
			sb "$v bare 60 $faction $seed" -preset bare -band 60 -faction $faction -seed $seed "${rot[@]}"
		done
		for band in 50 40 30; do sb "$v bare $band $faction 1" -preset bare -band $band -faction $faction -seed 1 "${rot[@]}"; done
	done
done
echo "## B. Raid talents"
for t in guide:0533001000000000-053030310051202051-0530000000000000 \
	reverb-ak:0535001000000000-054030310051202051-0500000000000000; do
	for v in fs-after-ss fs-after-ss-mana50; do
	for faction in alliance horde; do
		for seed in 1 50001; do
			sb "${t%%:*} $v raid 60 $faction $seed" -preset raid -band 60 -faction $faction -seed $seed -talents "${t#*:}" -rotation "$R/rot-$v.json"
		done
	done
	done
done
