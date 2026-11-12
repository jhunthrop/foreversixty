set -u
run() { echo "### $*"; nice --adjustment=15 env GOMAXPROCS=6 go run ./cmd/spec-breakdown -repo-root .. -iterations 10000 -seed 1 "$@" 2>&1 | grep -E "DPS|out of mana|Lightning Bolt \(" ; }
for L in 33 34; do for r in 3 4; do run -spec shaman-elemental -preset bare -band 30 -level $L -faction horde -race orc -rank 15208=$r; done; done
for r in 4 5; do
  run -spec shaman-elemental -preset bare -band 30 -level 39 -faction horde -race orc -rank 15208=$r
  run -spec shaman-elemental -preset bare -band 30 -level 39 -faction alliance -race dwarf -rank 15208=$r
  run -spec shaman-elemental -preset bare -band 40 -level 40 -faction horde -race orc -rank 15208=$r
done
