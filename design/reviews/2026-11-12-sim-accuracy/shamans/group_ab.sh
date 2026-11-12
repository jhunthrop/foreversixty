set -u
R=../design/reviews/2026-11-12-sim-accuracy/shamans
run() { echo "### $*"; nice --adjustment=15 env GOMAXPROCS=6 go run ./cmd/spec-breakdown -repo-root .. -iterations 10000 -seed 1 "$@" 2>&1 | grep -E "DPS|out of mana|Earth Shock \(10414\) \||Stormstrike \(17364\) \||Lightning Bolt \(15208\) \||Windfury Weapon|OtherActionAttack#1|^\| Maelstrom|Flurry \(16257\)" ; }
run -spec shaman-enhancement -preset raid -band 60 -faction horde
for p in 0 15 25 35 50; do run -spec shaman-enhancement -preset raid -band 60 -faction horde -rotation $R/enhancement-earth-shock-$p.json; done
for p in 0 15 25 35 50; do run -spec shaman-enhancement -preset bare -band 60 -faction horde -rotation $R/enhancement-earth-shock-$p.json; done
for p in 0 15 25 35 50; do run -spec shaman-enhancement -preset bare -band 40 -faction horde -rotation $R/enhancement-earth-shock-$p.json; done
