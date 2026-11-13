#!/usr/bin/env bash
# deltas.sh: every spec's published raid 60 entry (alliance) under the engine
# this checkout is pinned to, for the 70938 engine-follow pin's delta table.
# 10,000 iterations, seed 1, one sim at a time, niced to 15, GOMAXPROCS=5.
# Run from the site's sim directory (tools/remote-test.sh --dir sim, with
# --engine-sha for the candidate engine). Prints "spec<TAB>mean<TAB>error".
set -euo pipefail
make -C .. simdb >/dev/null
bin=$(mktemp -d)/spec-breakdown
go build -o "$bin" ./cmd/spec-breakdown
build=$(python3 -c 'import json;print(json.load(open("../web/src/data/active-build.json"))["build"])')
for file in ../data/builds/"$build"/bis/*.json; do
	spec=$(basename "$file" .json)
	out=$(nice --adjustment=15 env GOMAXPROCS=5 "$bin" -repo-root .. -spec "$spec" -preset raid -band 60 -faction alliance -seed 1 -iterations 10000 -json 2>/dev/null) || { echo "$spec	error"; continue; }
	python3 -c 'import json,sys; r=json.loads(sys.argv[2]); print("%s\t%.1f\t%.1f" % (sys.argv[1], r["DPS"], r["Error"]))' "$spec" "$out"
done
