#!/usr/bin/env bash
# orders.sh: the order of the last ten points of the two new level-60
# leveling builds (talent-search winners), on the bare band-50 gear at the
# level, both factions' guide races, seeds 1 and 50001, 10,000 iterations,
# one sim at a time, niced to 15, GOMAXPROCS=5. Run from the site's sim
# directory (tools/remote-test.sh --dir sim --engine-sha <sha>).
set -euo pipefail
make -C .. simdb >/dev/null
bin=$(mktemp -d)/spec-breakdown
go build -o "$bin" ./cmd/spec-breakdown
sb() { # label spec level talents
	printf '%s\t' "$1"
	for faction in alliance horde; do for seed in 1 50001; do
		nice --adjustment=15 env GOMAXPROCS=5 "$bin" -repo-root .. -iterations 10000 -json -spec "$2" -preset bare -band 50 -level "$3" -faction $faction -seed $seed -talents "$4" 2>/dev/null |
			python3 -c 'import json,sys; d=json.load(sys.stdin); print(round(d["DPS"],2), round(d["Error"],2), round(d.get("SecondsOOM",0),1), end="  ")' || printf 'FAILED  '
	done; done
	echo
}
# Enhancement: Elemental tree digits Convection,Concussion,Devastation,...; Enhancement tree ...ImpGhostWolf(12),SpiritWeapons,MentalQuickness,ImprovedStormstrike(15),Maelstrom,Farseer; Restoration IHW,Tidal,Mindfulness.
E33=055030310051202051   # the guide's 33 Enhancement points (Improved Stormstrike 0)
E35=055030310051202251   # with Improved Stormstrike 2
echo "## Enhancement 50: Convection 3 (published) against Convection 1 + Improved Stormstrike 2"
sb "enh 50 conv3" shaman-enhancement 50 "3000000000000000-$E33-0500000000000000"
sb "enh 50 conv1 impss2" shaman-enhancement 50 "1000000000000000-$E35-0500000000000000"
echo "## Enhancement 53 (44 points): Conv 3 + ImpSS 2 + Conv 1 | Conv 5 + Dev 1 | Conv 3 + Dev 3"
sb "enh 53 conv4 impss2" shaman-enhancement 53 "4000000000000000-$E35-0500000000000000"
sb "enh 53 conv5 dev1" shaman-enhancement 53 "5010000000000000-$E33-0500000000000000"
sb "enh 53 conv3 dev3" shaman-enhancement 53 "3030000000000000-$E33-0500000000000000"
echo "## Enhancement 57 (48 points): Conv 5 Dev 3 ImpSS 2 | Conv 5 Dev 3 Mindfulness 2 | Conv 5 Dev 1 ImpSS 2 Mindfulness 2"
sb "enh 57 conv5 dev3 impss2" shaman-enhancement 57 "5030000000000000-$E35-0500000000000000"
sb "enh 57 conv5 dev3 mind2" shaman-enhancement 57 "5030000000000000-$E33-0520000000000000"
sb "enh 57 conv5 dev1 impss2 mind2" shaman-enhancement 57 "5010000000000000-$E35-0520000000000000"
echo "## Enhancement 60 check: the winner"
sb "enh 60 winner" shaman-enhancement 59 "5030000000000000-$E35-0530000000000000"
ELE=5500331220001351
echo "## Elemental 55 (46 points): Thundering 5 + AK 5 | Thundering 5 + AK 3 + Mindfulness 2"
sb "ele 55 ak5" shaman-elemental 55 "$ELE-055000000000000000-0500000000000000"
sb "ele 55 ak3 mind2" shaman-elemental 55 "$ELE-053000000000000000-0520000000000000"
echo "## Elemental 58 (49 points): AK 5 + Shamanistic Focus 1 + Mindfulness 2 | AK 5 + Mindfulness 2 + Improved Reincarnation 1 | AK 5 + Mindfulness 3"
sb "ele 58 ak5 sf1 mind2" shaman-elemental 58 "$ELE-055000010000000000-0520000000000000"
sb "ele 58 ak5 mind2 ir1" shaman-elemental 58 "$ELE-055000000000000000-0520010000000000"
sb "ele 58 ak5 mind3" shaman-elemental 58 "$ELE-055000000000000000-0530000000000000"
