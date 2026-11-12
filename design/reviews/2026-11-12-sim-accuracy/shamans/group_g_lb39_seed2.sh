set -u
run() { echo "### $*"; nice --adjustment=15 env GOMAXPROCS=6 go run ./cmd/spec-breakdown -repo-root .. -iterations 10000 -seed 50001 "$@" 2>&1 | grep -E "DPS" ; }
for race in windshaper-skyborne tauren dwarf; do for r in 4 5; do run -spec shaman-elemental -preset bare -band 30 -level 39 -faction horde -race $race -rank 15208=$r; done; done
run -spec shaman-elemental -preset bare -band 30 -level 39 -faction alliance -race dwarf -rank 15208=4
run -spec shaman-elemental -preset bare -band 30 -level 39 -faction alliance -race dwarf -rank 15208=5
