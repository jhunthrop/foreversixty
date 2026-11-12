set -u
run() { echo "### $*"; nice --adjustment=15 env GOMAXPROCS=6 go run ./cmd/spec-breakdown -repo-root .. -iterations 10000 -seed 1 "$@" 2>&1 | grep -E "DPS|out of mana|Flametongue|OtherActionAttack#1|Rockbiter" ; }
for L in 10 14 16 18 19 24 29; do
  for imbue in rockbiter_weapon flametongue_weapon; do
    run -spec shaman-enhancement -preset bare -band 20 -level $L -faction horde -consumes main_hand_imbue:$imbue
  done
done
