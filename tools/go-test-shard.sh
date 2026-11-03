#!/usr/bin/env bash
# tools/go-test-shard.sh <shard> <total> <weights>: reads Go import paths (go list ./...) on
# stdin and prints the ones shard <shard> of <total> tests, so CI can split one module's
# `go test -p 1` run across parallel jobs (api.yml's test matrix) with every package in
# exactly one shard. Packages are dealt heaviest first to the lightest shard (longest
# processing time first), weighted by <weights>: "<import path> <seconds>" lines, # comments
# allowed, from a recent CI log. A package the file does not list weighs 1, so a new package
# still lands in a shard; refresh the file when a shard runs long.
#
#   go list ./... | ../tools/go-test-shard.sh 2 4 test-weights.txt
set -euo pipefail

usage() {
	echo "usage: go list ./... | tools/go-test-shard.sh <shard> <total> <weights file>" >&2
	exit 64
}

(($# == 3)) || usage
readonly SHARD=$1 TOTAL=$2 WEIGHTS=$3
if ! [[ $SHARD =~ ^[1-9][0-9]*$ && $TOTAL =~ ^[1-9][0-9]*$ ]] || ((SHARD > TOTAL)); then
	usage
fi
[[ -r $WEIGHTS ]] || {
	echo "go-test-shard: cannot read $WEIGHTS" >&2
	exit 1
}

awk 'FNR == NR { if ($1 !~ /^#/ && NF == 2) weight[$1] = $2; next }
	NF { print (($1 in weight) ? weight[$1] : 1), $1 }' "$WEIGHTS" - |
	sort -k1,1nr -k2,2 |
	awk -v shard="$SHARD" -v total="$TOTAL" 'BEGIN { for (i = 1; i <= total; i++) load[i] = 0 }
	{
		best = 1
		for (i = 2; i <= total; i++) if (load[i] < load[best]) best = i
		load[best] += $1
		if (best == shard) print $2
	}'
