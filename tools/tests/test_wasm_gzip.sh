#!/usr/bin/env bash
# Tests for tools/wasm-gzip.sh, which writes the sim.wasm.gz the site deploys and gates
# the transfer budget on it (Makefile `artifacts`, sim.yml's artifact job). Plain bash, no
# bats, the shape of tools/tests/test_remeasure.sh. The input is a compressible stand-in
# for the module, not a real wasm: the script never reads the format.
#
#   bash tools/tests/test_wasm_gzip.sh
#
# Prints "ok N tests" and exits 0 when every assertion holds; names each failure and exits
# 1 otherwise.
set -uo pipefail

HERE=$(cd "$(dirname "$0")" && pwd)
readonly HERE
readonly SCRIPT=$HERE/../wasm-gzip.sh

SCRATCH=$(mktemp -d)
trap 'rm -rf "$SCRATCH"' EXIT
PASSED=0
FAILED=0

assert() {
	local what=$1
	shift
	if "$@"; then
		PASSED=$((PASSED + 1))
	else
		FAILED=$((FAILED + 1))
		echo "FAIL: $what" >&2
	fi
}
equal() { [[ $1 == "$2" ]]; }
contains() { [[ $1 == *"$2"* ]]; }

# A few hundred KB of text with repeats a better deflate can find and gzip -9's window
# mostly cannot: the same shape as the module's name tables.
module=$SCRATCH/sim.wasm
for i in $(seq 1 20000); do printf 'github.com/wowsims/classic/sim/core.(*Unit).Spell%05d\n' $((i * 7919 % 20011)); done >"$module"
out=$SCRATCH/sim.wasm.gz

STATUS=0
OUT=$("$SCRIPT" "$module" "$out" 4.5 2>&1) || STATUS=$?
assert "within budget: exits 0 ($OUT)" equal "$STATUS" 0
assert "writes a gzip the system gzip inflates to the module" cmp -s <(gzip -dc "$out") "$module"
assert "reports the bytes and the budget share" contains "$OUT" "% of the 4.50 MiB budget"
assert "no temporary file is left" test ! -e "$out.tmp"
gzip9=$(gzip -9 -n -c "$module" | wc -c | tr -d ' ')
ours=$(wc -c <"$out" | tr -d ' ')
assert "smaller than gzip -9 ($ours vs $gzip9)" test "$ours" -lt "$gzip9"
first=$(shasum -a 256 <"$out")
"$SCRIPT" "$module" "$out" 4.5 >/dev/null 2>&1
assert "the same module gives the same bytes" equal "$(shasum -a 256 <"$out")" "$first"

STATUS=0
OUT=$("$SCRIPT" "$module" "$SCRATCH/tiny.gz" 0.01 2>&1) || STATUS=$?
assert "over budget: exits 1" equal "$STATUS" 1
assert "over budget: says so" contains "$OUT" "over the 0.01 MiB budget"

STATUS=0
OUT=$("$SCRIPT" "$SCRATCH/missing.wasm" "$SCRATCH/x.gz" 4.5 2>&1) || STATUS=$?
assert "missing module: exits 1" equal "$STATUS" 1
assert "missing module: names it" contains "$OUT" "missing.wasm is missing or empty"

STATUS=0
OUT=$("$SCRIPT" "$module" "$SCRATCH/x.gz" lots 2>&1) || STATUS=$?
assert "bad budget: exits 1" equal "$STATUS" 1

STATUS=0
OUT=$("$SCRIPT" "$module" 2>&1) || STATUS=$?
assert "wrong arguments: exits 1 with usage" contains "$STATUS:$OUT" "1:wasm-gzip: usage"

# A pigz whose stream does not inflate back is refused before the budget is read.
mkdir -p "$SCRATCH/bin"
printf '#!/usr/bin/env bash\nprintf "not a gzip"\n' >"$SCRATCH/bin/pigz"
chmod +x "$SCRATCH/bin/pigz"
STATUS=0
OUT=$(PATH="$SCRATCH/bin:$PATH" "$SCRIPT" "$module" "$SCRATCH/bad.gz" 4.5 2>&1) || STATUS=$?
assert "a stream that does not inflate: exits 1" equal "$STATUS" 1
assert "a stream that does not inflate: says so" contains "$OUT" "does not inflate back"
assert "a stream that does not inflate: writes nothing" test ! -e "$SCRATCH/bad.gz"

# The callers: the budget and the deployed file both come from this script, and nothing
# compresses the module a second, different way.
ROOT=$HERE/../..
# The single quotes are the point: these are the Makefile's own make-variable spellings.
# shellcheck disable=SC2016
assert "make artifacts writes and gates sim.wasm.gz with the script" \
	grep -qF './tools/wasm-gzip.sh $(ARTIFACT_DIR)/sim.wasm $(ARTIFACT_DIR)/sim.wasm.gz $(WASM_BUDGET_MB)' "$ROOT/Makefile"
# shellcheck disable=SC2016
assert "publish-wasm deploys that file as is" \
	grep -qF 'cp $(ARTIFACT_DIR)/sim.wasm.gz "$(WEB_SIM_DIR)/$$id/sim.wasm.gz"' "$ROOT/Makefile"
assert "sim.yml's budget step runs the script" \
	grep -qF './tools/wasm-gzip.sh artifacts/sim.wasm artifacts/sim.wasm.gz' "$ROOT/.github/workflows/sim.yml"
assert "no caller gzips sim.wasm itself" \
	bash -c "! grep -nE 'gzip -9[^|]*sim\.wasm' '$ROOT/Makefile' '$ROOT/.github/workflows/sim.yml' '$ROOT/.github/workflows/web.yml'"

if [[ $FAILED -gt 0 ]]; then
	echo "$FAILED failed, $PASSED passed" >&2
	exit 1
fi
echo "ok $PASSED tests"
