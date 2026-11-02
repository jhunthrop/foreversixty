#!/usr/bin/env bash
# tools/wasm-gzip.sh <sim.wasm> <out.gz> <budget MiB>
#
# Compresses the browser module into the gzip file the site deploys, proves it, and
# enforces the transfer budget on THAT file. The Makefile's `artifacts` target and
# sim.yml's artifact job both call it, so the budget always measures the bytes a visitor
# downloads rather than a second, different compression of the same module.
#
# The compressor is pigz's zopfli mode (-11): its output is a plain gzip stream, the one
# the page's DecompressionStream('gzip') already inflates (web/src/lib/sim/wasm-fetch.ts),
# about 4.7% smaller than `gzip -9` on the 2026-11-02 module (4,276,658 against 4,481,935
# bytes). -I 3 keeps most of that for a fifth of the default 15 iterations' time. pigz 2.8
# is on GitHub's ubuntu-24.04 image; on a Mac, `brew install pigz`. -n keeps the header
# free of a name and timestamp, so the same module always gives the same bytes.
#
# The result is inflated with the system gzip and compared with the module before the
# budget is read: a stream the standard inflater cannot read never reaches the deploy.
set -euo pipefail

readonly PIGZ_FLAGS=(-11 -I 3 -n -c)

die() {
	echo "wasm-gzip: $*" >&2
	exit 1
}

[[ $# -eq 3 ]] || die "usage: wasm-gzip.sh <sim.wasm> <out.gz> <budget MiB>"
readonly module=$1 out=$2 budget_mib=$3
[[ -s $module ]] || die "$module is missing or empty"
[[ $budget_mib =~ ^[0-9]+(\.[0-9]+)?$ ]] || die "budget '$budget_mib' is not a number of MiB"
command -v pigz >/dev/null || die "pigz is not installed (GitHub's ubuntu image ships it; on a Mac: brew install pigz)"

pigz "${PIGZ_FLAGS[@]}" "$module" >"$out.tmp"
gzip -dc "$out.tmp" | cmp -s - "$module" || {
	rm -f "$out.tmp"
	die "$out does not inflate back to $module"
}
mv "$out.tmp" "$out"

bytes=$(wc -c <"$out" | tr -d ' ')
awk -v b="$bytes" -v m="$budget_mib" 'BEGIN {
	printf "sim.wasm.gz: %d bytes, %.2f MiB, %.1f%% of the %.2f MiB budget\n", b, b/1048576, 100*b/(m*1048576), m
	exit (b <= m*1048576) ? 0 : 1
}' || die "$out is over the $budget_mib MiB budget"
