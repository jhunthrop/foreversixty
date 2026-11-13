#!/usr/bin/env bash
# bis-pending.sh: the specs of a bis.yml rank matrix whose picks were not ranked on these
# inputs.
#
#   tools/bis-pending.sh <build> <inputs hash> <matrix json>
#
# Run from the repository root. Prints, comma-separated in the matrix's order, every spec
# whose data/builds/<build>/bis/<spec>.inputs is missing or holds another hash; prints
# nothing when every spec is done. The matrix is tools/bis-matrix.sh's output.
set -euo pipefail
[ $# -eq 3 ] || { echo "usage: $0 <build> <inputs hash> <matrix json>" >&2; exit 2; }
build=$1 hash=$2 matrix=$3
pending=()
while IFS= read -r spec; do
  [ -n "$spec" ] || continue
  marker=$(cat "data/builds/$build/bis/$spec.inputs" 2>/dev/null || true)
  [ "$marker" = "$hash" ] || pending+=("$spec")
done < <(jq -r '.include[].specs | split(",")[]' <<<"$matrix")
(IFS=,; echo "${pending[*]:-}")
