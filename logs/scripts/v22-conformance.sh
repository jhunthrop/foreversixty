#!/usr/bin/env bash
# Sweep a directory of combat logs and fail unless every file parsed
# cleanly under a verified layout.
#
# Usage: logs/scripts/v22-conformance.sh <log-dir> [out.jsonl]
#
# The v22 corpus this was written against is 89 files and 7.3 GB and is
# not in the repository; it lives in the session scratchpad. The script is
# committed so the sweep is reproducible against any directory of logs.
set -euo pipefail

dir="${1:?usage: v22-conformance.sh <log-dir> [out.jsonl]}"
out="${2:-/tmp/v22-conformance.jsonl}"

cd "$(dirname "$0")/.."
go build -o /tmp/forever-logs ./cmd/forever-logs
/tmp/forever-logs conformance -json "$dir" > "$out"

python3 - "$out" <<'PY'
import json, sys, collections
rows = [json.loads(line) for line in open(sys.argv[1])]
unknown = collections.Counter()
errors = lines = fights = 0
bad_layout = []
for r in rows:
    for k, v in (r.get("unknown_events") or {}).items():
        unknown[k] += v
    errors += r["parse_errors"]
    lines += r["lines"]
    fights += r["fights"]
    if r["layout"] != "retail-v22" or not r["layout_verified"] or r["layout_inferred"]:
        bad_layout.append((r["path"], r["layout"], r["layout_verified"]))

print(f"files        {len(rows)}")
print(f"lines        {lines}")
print(f"fights       {fights}")
print(f"parse errors {errors}")
print(f"unknown      {len(unknown)} distinct")
for name, n in sorted(unknown.items()):
    print(f"  {name:<34} {n}")

fail = False
if bad_layout:
    fail = True
    print(f"\nFAIL: {len(bad_layout)} files did not select a verified retail-v22 row")
    for path, lay, ver in bad_layout[:10]:
        print(f"  {path} layout={lay} verified={ver}")
if errors:
    fail = True
    print(f"\nFAIL: {errors} parse errors")
if unknown:
    fail = True
    print(f"\nFAIL: {len(unknown)} distinct unknown events")
if fail:
    sys.exit(1)
print("\nPASS: every file parsed clean under retail-v22")
PY
