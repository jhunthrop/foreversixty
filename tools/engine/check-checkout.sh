#!/usr/bin/env bash
# check-checkout.sh: does a local build compile the engine the site pins?
#
#   tools/engine/check-checkout.sh <site-dir> <engine-replace>
#
# <engine-replace> is where github.com/wowsims/classic resolves for this build
# (`go list -m -f '{{with .Replace}}{{.Path}}{{end}}' github.com/wowsims/classic`
# from <site>/sim, which honours GOWORK). CI rewrites sim/go.mod's development
# replace to the fork at the pin (.github/actions/pin-engine), so there it is a
# module path, not a directory, and this says nothing. On a laptop it is the
# owner's fork checkout, which is often on another commit than
# <site>/sim/enginever/version.go names; then `make artifacts` builds a
# sim.wasm CI never would, and this warns on stderr with the way to build the
# pin instead. It never fails the build (the developer flow builds the fork
# checkout on purpose); exit 2 on bad arguments only.
set -euo pipefail

readonly VERSION_FILE=sim/enginever/version.go

main() {
	[[ $# -eq 2 ]] || { echo "usage: $0 <site-dir> <engine-replace>" >&2; exit 2; }
	local site=$1 engine=$2 pin head
	[[ -d $engine ]] || exit 0
	pin=$(sed -n 's/.*Version = "\(.*\)"/\1/p' "$site/$VERSION_FILE" 2>/dev/null || true)
	[[ -n $pin ]] || { echo "check-checkout: $site/$VERSION_FILE has no Version" >&2; exit 0; }
	head=$(git -C "$engine" rev-parse HEAD 2>/dev/null || true)
	if [[ $head != "$pin"* ]]; then
		echo "warning: the engine resolves to $engine at ${head:0:9}, not the pinned $pin." >&2
		echo "  This build differs from CI's. To build the pin: git -C $engine worktree add <dir> $pin," >&2
		echo "  then pass GOWORK=<a copy of go.work plus 'replace github.com/wowsims/classic => <dir>'>." >&2
	elif [[ -n $(git -C "$engine" status --porcelain 2>/dev/null) ]]; then
		echo "warning: the engine checkout at $engine is at the pin $pin but has uncommitted changes;" >&2
		echo "  this build differs from CI's." >&2
	fi
}

main "$@"
