#!/usr/bin/env bash
# bis.yml's re-measure of a carried dps-per-point.json. tools/activate-build.sh carries the
# previous build's dps-per-point.json onto a new build and writes dps-per-point.carried-from
# beside it; the first ranking on the new build re-measures the figures here, deletes the
# marker, and keeps the build's manifest.json in step: the marker's entry dropped and the
# re-measured file re-hashed. 2026-10-31: the marker was deleted but left listed, and every
# web build failed sync-data's "manifest.json lists files that are missing" check.
#
#   tools/remeasure-dps-per-point.sh <build>
#
# Run from the repository root. A build with no marker is its own measure and is left
# untouched. On a re-measure, appends measured=true to $GITHUB_OUTPUT when it is set.
# Environment:
#   DPP_MEASURE   the measuring command (tests stub it); default the real measure
set -euo pipefail

readonly MARKER=dps-per-point.carried-from
readonly FIGURES=dps-per-point.json
readonly DEFAULT_MEASURE="go run ./sim/cmd/dps-per-point -repo-root ."

[ $# -eq 1 ] || { echo "usage: $0 <build>" >&2; exit 2; }
readonly DIR="data/builds/$1"
readonly MANIFEST="$DIR/manifest.json"
[ -d "$DIR" ] || { echo "$DIR does not exist" >&2; exit 2; }

if [ ! -f "$DIR/$MARKER" ]; then
  echo "$FIGURES is $1's own (no $MARKER); kept"
  exit 0
fi

echo "$FIGURES is carried from $(cat "$DIR/$MARKER"); re-measuring on tonight's picks"
# shellcheck disable=SC2086 # the command is a word list by design
${DPP_MEASURE:-$DEFAULT_MEASURE}
rm "$DIR/$MARKER"

# The manifest records every published file's digest: the deleted marker leaves it, the
# re-measured figures get their new digest. A build without a manifest has nothing to keep.
if [ -f "$MANIFEST" ]; then
  digest=$(sha256sum "$DIR/$FIGURES" | cut -d' ' -f1)
  tmp=$(mktemp)
  jq --indent 2 --arg marker "$MARKER" --arg figures "$FIGURES" --arg digest "$digest" \
    'del(.files[$marker]) | .files[$figures] = $digest' \
    "$MANIFEST" >"$tmp"
  mv "$tmp" "$MANIFEST"
fi

if [ -n "${GITHUB_OUTPUT:-}" ]; then echo "measured=true" >>"$GITHUB_OUTPUT"; fi
