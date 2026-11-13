#!/usr/bin/env bash
# jobs.sh [imbue|search|all]: the enhancement-70938 lane's measurements, one sim at a
# time, nice 15, GOMAXPROCS=5. Run from the site's sim directory on the box:
#   tools/remote-test.sh --dir sim --engine-sha <sha> HEAD -- bash ../design/reviews/2026-11-13-sim-accuracy/enhancement-70938/jobs.sh all
#   imbue: Windfury, Rockbiter and Flametongue Weapon at bare 20 to 60 (both factions)
#          and raid 60 (both factions, the raid's other consumables kept), seed 1.
#   search: the raid and bare 60 talent and rotation searches (Alliance, seed 7).
set -uo pipefail
what=${1:-all}
make -C .. simdb >/dev/null || exit 1
bin=$(mktemp -d)
for c in spec-breakdown talent-search rotation-search; do go build -o "$bin/$c" ./cmd/$c || exit 1; done
sb() { # label, then spec-breakdown args
	local label=$1; shift
	printf '%s\t' "$label"
	nice -n 15 env GOMAXPROCS=5 "$bin/spec-breakdown" -repo-root .. -iterations 10000 -seed 1 -json "$@" 2>/dev/null |
		python3 -c 'import json,sys; d=json.load(sys.stdin); print("%.1f\t%.1f\t%.1f" % (d["DPS"], d["Error"], d.get("SecondsOOM", 0)))' || echo FAILED
}
if [[ $what == imbue || $what == all ]]; then
	raid=major_frenzy_potion,food_smoked_desert_dumpling,flask_of_natural_aggression,elixir_of_the_mongoose,juju_power,r_o_i_d_s,juju_might
	for faction in alliance horde; do
		for imbue in windfury_weapon rockbiter_weapon flametongue_weapon; do
			sb "raid 60 $faction $imbue" -spec shaman-enhancement -preset raid -band 60 -faction $faction -consumes main_hand_imbue:$imbue,$raid
			for band in 60 50 40 30 20; do
				[[ $band == 20 && $imbue == windfury_weapon ]] && continue
				sb "bare $band $faction $imbue" -spec shaman-enhancement -preset bare -band $band -faction $faction -consumes main_hand_imbue:$imbue
			done
		done
	done
fi
if [[ $what == search || $what == all ]]; then
	out=$(mktemp -d); mkdir -p "$out/talents" "$out/rotation"
	run() { nice -n 15 env GOMAXPROCS=5 "$@" 2>&1 | grep -v '^$' | tail -2; }
	run "$bin/talent-search" -repo-root .. -out "$out/talents" -spec shaman-enhancement -preset raid -keep 'Stormstrike,Maelstrom Weapon,Rage of the Farseer,Elemental Focus,Improved Ghost Wolf,Tidal Focus,Shamanistic Focus'
	run "$bin/talent-search" -repo-root .. -out "$out/talents" -spec shaman-enhancement -keep 'Stormstrike,Maelstrom Weapon,Rage of the Farseer,Improved Ghost Wolf,Tidal Focus,Shamanistic Focus'
	run "$bin/rotation-search" -repo-root .. -out "$out/rotation" -spec shaman-enhancement -preset raid
	run "$bin/rotation-search" -repo-root .. -out "$out/rotation" -spec shaman-enhancement
	for f in "$out"/*/*.md; do echo "=== $(basename "$(dirname "$f")")/$(basename "$f")"; cat "$f"; done
fi
