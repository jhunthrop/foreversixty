#!/usr/bin/env bash
# jobs.sh <jobfile>: one spec-breakdown run per line ("<name> <flags...>"), one at a time,
# nice -n 15, GOMAXPROCS=5. Tanks print the tank fight (-tank), healers the heal fight (-heal),
# the rest JSON totals. Prints the checkout's provenance first: HEAD, the active build, the
# simdb.bin hashes (the build's and the embedded copy) and the engine the module resolves.
# Run from sim/ on the box:
#   tools/remote-test.sh --dir sim [--engine-sha <sha>] <ref> -- bash ../design/reviews/2026-11-13-sim-accuracy/bis-verification-70938/jobs.sh <jobfile>
set -uo pipefail
here=$(cd "$(dirname "$0")" && pwd)
make -C .. simdb >/dev/null || exit 1
build=$(python3 -c 'import json;print(json.load(open("../web/src/data/active-build.json"))["build"])')
echo "##### provenance"
echo "head $(git rev-parse HEAD)"
echo "build $build"
sha256sum "../data/builds/$build/simdb.bin" internal/simdb/simdb.bin
echo "engine $(go list -m -f '{{if .Replace}}{{.Replace.Path}} {{.Replace.Version}}{{else}}{{.Version}}{{end}}' github.com/wowsims/classic 2>/dev/null || grep -E '^replace' go.mod)"
bin=$(mktemp -d)/spec-breakdown
go build -o "$bin" ./cmd/spec-breakdown || exit 1
while read -r name flags; do
	[[ -z $name || $name == \#* ]] && continue
	case $name in
	*-tank-*) mode=-tank ;;
	*-heal-*) mode=-heal ;;
	*) mode=-json ;;
	esac
	echo "##### $name"
	# shellcheck disable=SC2086
	nice -n 15 env GOMAXPROCS=5 "$bin" -repo-root .. $mode $flags 2>&1
done < "$here/$1"
