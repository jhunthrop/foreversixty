set -u
run() { echo "### $*"; nice --adjustment=15 env GOMAXPROCS=6 go run ./cmd/spec-breakdown -repo-root .. -iterations 10000 -seed 1 "$@" 2>&1 | grep -E "DPS|out of mana|Earth Shock \(" ; }
for L in 25 26 27; do
  for r in 3 4; do
    run -spec shaman-enhancement -preset bare -band 20 -level $L -faction horde -race orc -rank 10414=$r
    run -spec shaman-enhancement -preset bare -band 20 -level $L -faction alliance -race dwarf -rank 10414=$r
  done
done
