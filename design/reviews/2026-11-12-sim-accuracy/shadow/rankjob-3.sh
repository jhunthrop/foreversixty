#!/usr/bin/env bash
# rankjob-3.sh: the flip and tie cells of the second pass at seed 50001
# (rankcells-seed2.txt), then the lane's gates and the ladder goldens.
set -uo pipefail
here=$(cd "$(dirname "$0")" && pwd)
echo "=== seed 50001"
SEED=50001 CELLS="$here/rankcells-seed2.txt" bash "$here/rankjob.sh"
echo "=== gates"
cd "$here/../../../.." && bash design/reviews/2026-11-12-sim-accuracy/shadow/gates.sh
