#!/usr/bin/env bash
# fork_tests.sh <pin sha> <fix sha>: on the box, in a throwaway worktree of /home/ci/wowsims-forever, run the
# Deep Wounds test red at the pin (the new test file dropped in) and green at the fix, then the fork suites
# the warrior touches, with the two suite goldens rewritten and printed back as a patch.
set -uo pipefail
pin=$1; fix=$2
export PATH=/usr/local/go/bin:/home/ci/go/bin:$HOME/go/bin:$PATH
repo=/home/ci/wowsims-forever; wt=/home/ci/wf-dpswar
git -C $repo fetch -q origin "+refs/heads/lane/sim-dpswar:refs/remotes/origin/lane/sim-dpswar" master
git -C $repo worktree remove --force $wt 2>/dev/null; git -C $repo worktree add -q --detach $wt "$fix" || exit 1
cd $wt || exit 1
export GOMAXPROCS=5 GOFLAGS=-mod=mod
echo "### red at $pin"
git show "$fix":sim/warrior/dps_warrior/deep_wounds_crit_test.go > /tmp/dw_test.go
git checkout -q "$pin" && cp /tmp/dw_test.go sim/warrior/dps_warrior/deep_wounds_crit_test.go
nice -n 15 go test -tags=with_db -count=1 -run 'TestDeepWoundsTicksNeverCrit' ./sim/warrior/dps_warrior/ 2>&1 | tail -5
git checkout -q -- . && git checkout -q "$fix"
echo "### green at $fix"
nice -n 15 go test -tags=with_db -count=1 -run 'TestDeepWoundsTicksNeverCrit' ./sim/warrior/dps_warrior/ 2>&1 | tail -3
echo "### suites"
nice -n 15 go test -tags=with_db -p 1 -count=1 ./sim/core/ ./sim/warrior/... ./sim/common/... ./sim/conformance/ 2>&1 | grep -E "^(ok|FAIL|---|\s+--- FAIL|panic)" | grep -v "no test files" | cut -c1-160
echo "### goldens"
for f in $(find sim/warrior -name "*.results.tmp"); do cp "$f" "${f%.results.tmp}.results"; echo "updated ${f%.results.tmp}.results"; done
nice -n 15 go test -tags=with_db -count=1 ./sim/warrior/dps_warrior/ ./sim/warrior/tank_warrior/ 2>&1 | tail -3
echo "=====PATCH====="
git diff -- sim/warrior/dps_warrior/TestP1DPSWarrior.results sim/warrior/tank_warrior/TestP1TankWarrior.results
echo "=====END====="
git -C $repo worktree remove --force $wt
