set -uo pipefail
make -C .. simdb >/dev/null
bin=$(mktemp -d)/spec-breakdown
go build -o "$bin" ./cmd/spec-breakdown
run() { # spec preset band level faction iterations
	nice -n 15 env GOMAXPROCS=5 "$bin" -repo-root .. -spec "$1" -preset "$2" -band "$3" -level "$4" -faction "$5" -seed 1 -iterations "$6" -json 2>/dev/null |
		python3 -c 'import json,sys; d=json.load(sys.stdin); print(round(d["DPS"],3), round(d["Error"],3), round(d.get("SecondsOOM",0),1), end="")' || echo -n "FAILED"
	echo "	$1 $2 $3 level $4 $5"
}
for cell in "priest-shadow bare 20 10" "priest-shadow bare 20 15" "priest-shadow bare 20 20" "priest-shadow bare 20 25" "priest-shadow bare 20 27" "priest-shadow bare 30 30" "priest-shadow bare 30 35" "priest-shadow bare 40 40" "priest-shadow bare 50 50" "priest-shadow bare 60 60"; do
	set -- $cell
	for faction in alliance horde; do run "$1" "$2" "$3" "$4" "$faction" 5000; done
done
run priest-shadow raid 60 60 alliance 5000
run priest-shadow raid 60 60 horde 5000
for cell in "mage-frost bare 20 20" "mage-arcane bare 20 20" "mage-fire bare 20 20" "warlock-affliction bare 20 20" "warlock-destruction bare 20 20" "warlock-demonology bare 20 20" "mage-frost bare 40 40" "warlock-affliction bare 40 40"; do
	set -- $cell; run "$1" "$2" "$3" "$4" alliance 3000
done
