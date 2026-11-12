set -u
run() { echo "### $*"; nice --adjustment=15 env GOMAXPROCS=6 go run ./cmd/spec-breakdown -repo-root .. -iterations 10000 -seed 1 "$@" 2>&1 | grep -E "DPS|out of mana" ; }
for race in tauren troll windshaper-skyborne; do
  for L in 25 26; do for r in 3 4; do run -spec shaman-enhancement -preset bare -band 20 -level $L -faction horde -race $race -rank 10414=$r; done; done
  for r in 4 5; do run -spec shaman-elemental -preset bare -band 30 -level 39 -faction horde -race $race -rank 15208=$r; done
done
