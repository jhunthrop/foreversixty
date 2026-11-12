#!/usr/bin/env bash
# rankjob-2.sh: the second pass of rankjob.sh: the cells at 30 and above
# (rankcells-2.txt, the first pass ran them on band 20 and the tool refused),
# then the one seed-1 flip (rankcells-seed2.txt) at seed 50001.
set -uo pipefail
here=$(cd "$(dirname "$0")" && pwd)
CELLS="$here/rankcells-2.txt" bash "$here/rankjob.sh"
echo "=== seed 50001"
SEED=50001 CELLS="$here/rankcells-seed2.txt" bash "$here/rankjob.sh"
