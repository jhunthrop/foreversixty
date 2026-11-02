#!/usr/bin/env bash
# tools/wasm-compress.sh <sim.wasm> <out.gz|out.br> <budget MiB>
#
# Compresses the browser module into one of the two files the site deploys, proves it, and
# enforces that file's transfer budget on THAT file. The Makefile's `artifacts` target and
# sim.yml's artifact job both call it once per file, so each budget always measures the
# bytes a visitor downloads rather than a second, different compression of the same module.
# The output's extension picks the codec:
#
#   .gz  pigz's zopfli mode (-11): a plain gzip stream, the one the page's
#        DecompressionStream('gzip') inflates (web/src/lib/sim/wasm-fetch.ts) wherever the
#        Worker does not stand in front of the file (astro preview, every e2e run). About
#        4.7% smaller than `gzip -9` on the 2026-11-02 module. -I 3 keeps most of that for a
#        fifth of the default 15 iterations' time. -n keeps the header free of a name and
#        timestamp, so the same module always gives the same bytes.
#   .br  brotli -q 11 with a 16 MiB window (-w 24): about 28% smaller than the gzip on the
#        2026-11-02 module. Production's Worker (web/src/worker.ts) answers a request for the
#        .gz with these bytes and Content-Encoding: br, so the browser decodes them natively
#        and the page receives the module itself.
#
# pigz 2.8 and brotli 1.1 are on GitHub's ubuntu-24.04 image; on a Mac, `brew install pigz
# brotli`. The result is decoded with the system tool and compared with the module before
# the budget is read: a stream the standard decoder cannot read never reaches the deploy.
set -euo pipefail

readonly PIGZ_FLAGS=(-11 -I 3 -n -c)
readonly BROTLI_FLAGS=(-q 11 -w 24 -c)

die() {
	echo "wasm-compress: $*" >&2
	exit 1
}

need() {
	command -v "$1" >/dev/null || die "$1 is not installed (GitHub's ubuntu image ships it; on a Mac: brew install $1)"
}

[[ $# -eq 3 ]] || die "usage: wasm-compress.sh <sim.wasm> <out.gz|out.br> <budget MiB>"
readonly module=$1 out=$2 budget_mib=$3
[[ -s $module ]] || die "$module is missing or empty"
[[ $budget_mib =~ ^[0-9]+(\.[0-9]+)?$ ]] || die "budget '$budget_mib' is not a number of MiB"

case $out in
*.gz)
	need pigz
	compress() { pigz "${PIGZ_FLAGS[@]}" "$module"; }
	decompress() { gzip -dc "$1"; }
	;;
*.br)
	need brotli
	compress() { brotli "${BROTLI_FLAGS[@]}" "$module"; }
	decompress() { brotli -dc "$1"; }
	;;
*) die "$out is neither a .gz nor a .br file" ;;
esac

compress >"$out.tmp"
decompress "$out.tmp" 2>/dev/null | cmp -s - "$module" || {
	rm -f "$out.tmp"
	die "$out does not decompress back to $module"
}
mv "$out.tmp" "$out"

bytes=$(wc -c <"$out" | tr -d ' ')
awk -v f="$(basename "$out")" -v b="$bytes" -v m="$budget_mib" 'BEGIN {
	printf "%s: %d bytes, %.2f MiB, %.1f%% of the %.2f MiB budget\n", f, b, b/1048576, 100*b/(m*1048576), m
	exit (b <= m*1048576) ? 0 : 1
}' || die "$out is over the $budget_mib MiB budget"
