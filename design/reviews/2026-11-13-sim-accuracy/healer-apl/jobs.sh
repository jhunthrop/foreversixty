#!/usr/bin/env bash
# jobs.sh <jobfile>: one spec-breakdown -heal run per line of the job file ("<name> <flags...>"),
# printed in full; "@V" in a line's flags is this directory's variants/ folder. Run from sim/ on the
# box: tools/remote-test.sh --pin-engine --dir sim <ref> -- bash ../design/reviews/2026-11-13-sim-accuracy/healer-apl/jobs.sh <jobfile>
set -uo pipefail
here=$(cd "$(dirname "$0")" && pwd)
make -C .. simdb >/dev/null || exit 1
bin=$(mktemp -d)/spec-breakdown
go build -o "$bin" ./cmd/spec-breakdown || exit 1
while read -r name flags; do
	[[ -z $name || $name == \#* ]] && continue
	echo "##### $name"
	flags=${flags//@V/$here/variants}
	# shellcheck disable=SC2086
	nice -n 15 env GOMAXPROCS=5 "$bin" -repo-root .. -heal $flags 2>&1
done < "$here/$1"
