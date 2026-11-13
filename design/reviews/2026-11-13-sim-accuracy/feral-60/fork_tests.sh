#!/usr/bin/env bash
# fork_tests.sh <pin sha> <fix sha>: on the box (ssh jh@<box> "wsl -d forever-ci -- bash -l -s -- <pin> <fix>" < fork_tests.sh),
# in a throwaway worktree of /home/ci/wowsims-forever: the Rake test red at the pin (the fix's test file dropped in) and
# green at the fix, then the druid suites with any golden rewritten and printed back as a patch.
set -uo pipefail
pin=$1; fix=$2
export PATH=/usr/local/go/bin:/home/ci/go/bin:$HOME/go/bin:$PATH
repo=/home/ci/wowsims-forever; wt=/home/ci/wf-feral60
git -C $repo fetch -q origin "+refs/heads/lane/sim-feral60:refs/remotes/origin/lane/sim-feral60" master
git -C $repo worktree remove --force $wt 2>/dev/null; git -C $repo worktree add -q --detach $wt "$fix" || exit 1
cd $wt || exit 1
export GOMAXPROCS=5 GOFLAGS=-mod=mod
tests='TestRakeHitIsTheRanksAmountPlusTheTicksAttackPowerShare|TestRakeTicksTakeEightPercentOfAttackPower'
echo "### red at $pin"
git show "$fix":sim/druid/rake_test.go > /tmp/f60_rake_test.go
git checkout -q "$pin" && cp /tmp/f60_rake_test.go sim/druid/rake_test.go
nice -n 15 go test -tags=with_db -count=1 -run "$tests" ./sim/druid/ 2>&1 | grep -E "^(ok|FAIL|---|\s+--- FAIL|\s+\S+_test.go)" | cut -c1-200
git checkout -q -- . && git checkout -q "$fix"
echo "### green at $fix"
nice -n 15 go test -tags=with_db -count=1 -run "$tests" -v ./sim/druid/ 2>&1 | grep -E "^(ok|FAIL|--- |\s+--- )" | cut -c1-200
echo "### suites"
pk="./sim/druid/... ./sim/core/ ./sim/conformance/"
nice -n 15 go test -tags=with_db -p 1 -count=1 $pk 2>&1 | grep -E "^(ok|FAIL|---|\s+--- FAIL|panic)" | grep -v "no test files" | cut -c1-200
echo "### goldens"
n=0; for f in $(find sim -name "*.results.tmp"); do cp "$f" "${f%.results.tmp}.results"; echo "updated ${f%.results.tmp}.results"; n=$((n+1)); done
if [ $n -gt 0 ]; then nice -n 15 go test -tags=with_db -p 1 -count=1 $pk 2>&1 | grep -E "^(ok|FAIL)" | grep -v "no test files" | cut -c1-200; fi
echo "=====PATCH====="
git diff -- '*.results'
echo "=====END====="
cd / && git -C $repo worktree remove --force $wt
