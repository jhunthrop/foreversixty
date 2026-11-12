#!/usr/bin/env bash
# jobs2.sh (the raid imbue cells and the searches again after the presets note was repaired; the bare cells and the rows are in jobs.sh): the shamans-60 lane's measurements at the pin, one sim at a time,
# niced to 15, GOMAXPROCS=5. Run from the site's sim directory
# (tools/remote-test.sh --dir sim --pin-engine).
#
#   A. Elemental's Lightning Bolt rank rows at 50, 51, 56, 57 and bare 60 on
#      the hotfixed overload rows (MR43): the row's rank against the newest
#      rank, both factions' guide races, seeds 1 and 50001, 10,000 iterations.
#   B. Enhancement's imbue at bare 50, bare 60 and raid 60 (both factions):
#      Windfury Weapon against Rockbiter and Flametongue Weapon, the raid's
#      other consumables kept.
#   C. The raid and bare talent searches and rotation searches for both
#      specs on the fixed rotation-search (it wears the raid build under the
#      raid preset now).
set -euo pipefail
make -C .. simdb >/dev/null
bin=$(mktemp -d)
go build -o "$bin/spec-breakdown" ./cmd/spec-breakdown
go build -o "$bin/talent-search" ./cmd/talent-search
go build -o "$bin/rotation-search" ./cmd/rotation-search
sb() { # label, then spec-breakdown args
	local label=$1; shift
	printf '%s\t' "$label"
	nice --adjustment=15 env GOMAXPROCS=5 "$bin/spec-breakdown" -repo-root .. -iterations 10000 -json "$@" 2>/dev/null |
		python3 -c 'import json,sys; d=json.load(sys.stdin); print(round(d["DPS"],2), round(d["Error"],2), round(d.get("SecondsOOM",0),1))' || echo FAILED
}
echo "## B. Enhancement imbue, raid 60"
raid=major_frenzy_potion,food_smoked_desert_dumpling,flask_of_natural_aggression,elixir_of_the_mongoose,juju_power,r_o_i_d_s,juju_might
for imbue in windfury_weapon rockbiter_weapon flametongue_weapon; do
	for faction in alliance horde; do
		for seed in 1 50001; do
			sb "enh raid 60 $faction seed $seed $imbue" -spec shaman-enhancement -preset raid -band 60 -faction $faction -seed $seed -consumes main_hand_imbue:$imbue,$raid
		done
	done
done
echo "## C. Searches"
out=$(mktemp -d); mkdir -p "$out/talents" "$out/rotation"
run() { nice --adjustment=15 env GOMAXPROCS=5 "$@" 2>&1 | grep -v '^$' | tail -2; }
run "$bin/talent-search" -repo-root .. -out "$out/talents" -spec shaman-enhancement -preset raid -keep 'Stormstrike,Maelstrom Weapon,Rage of the Farseer,Elemental Focus,Improved Ghost Wolf,Tidal Focus,Shamanistic Focus'
run "$bin/talent-search" -repo-root .. -out "$out/talents" -spec shaman-enhancement -keep 'Stormstrike,Maelstrom Weapon,Rage of the Farseer,Improved Ghost Wolf,Tidal Focus,Shamanistic Focus'
run "$bin/talent-search" -repo-root .. -out "$out/talents" -spec shaman-elemental -preset raid -keep 'Lightning Mastery,Call of Thunder,Elemental Focus,Trial by Fire,Elemental Warding,Call of Flame,Tidal Focus'
run "$bin/talent-search" -repo-root .. -out "$out/talents" -spec shaman-elemental -keep 'Lightning Mastery,Call of Thunder,Elemental Focus,Trial by Fire,Elemental Warding,Call of Flame,Tidal Focus,Improved Reincarnation'
run "$bin/rotation-search" -repo-root .. -out "$out/rotation" -spec shaman-enhancement -preset raid
run "$bin/rotation-search" -repo-root .. -out "$out/rotation" -spec shaman-enhancement
run "$bin/rotation-search" -repo-root .. -out "$out/rotation" -spec shaman-elemental -preset raid
run "$bin/rotation-search" -repo-root .. -out "$out/rotation" -spec shaman-elemental
for f in "$out"/*/*.md; do echo "=== $(basename "$(dirname "$f")")/$(basename "$f")"; cat "$f"; done
