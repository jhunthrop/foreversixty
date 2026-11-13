#!/usr/bin/env bash
# fork_tests.sh <pin sha> <fix sha>: on the box, in a throwaway worktree of /home/ci/wowsims-forever, run the
# crit-flag tests red at the pin (the fix's test files dropped in) and green at the fix, then the fork suites
# the five spells touch, with any suite golden rewritten and printed back as a patch.
set -uo pipefail
pin=$1; fix=$2
export PATH=/usr/local/go/bin:/home/ci/go/bin:$HOME/go/bin:$PATH
repo=/home/ci/wowsims-forever; wt=/home/ci/wf-critflags
git -C $repo fetch -q origin "+refs/heads/lane/crit-flags:refs/remotes/origin/lane/crit-flags" master
git -C $repo worktree remove --force $wt 2>/dev/null; git -C $repo worktree add -q --detach $wt "$fix" || exit 1
cd $wt || exit 1
export GOMAXPROCS=5 GOFLAGS=-mod=mod
tests='TestIgniteTicksNeverCrit|TestGroundChannelTicksCrit'
echo "### red at $pin"
git show "$fix":sim/mage/ignite_test.go > /tmp/cf_ignite_test.go
git show "$fix":sim/conformance/ground_channel_crit_test.go > /tmp/cf_ground_test.go
git checkout -q "$pin" && cp /tmp/cf_ignite_test.go sim/mage/ignite_test.go && cp /tmp/cf_ground_test.go sim/conformance/ground_channel_crit_test.go
nice -n 15 go test -tags=with_db -count=1 -run "$tests" ./sim/mage/ ./sim/conformance/ 2>&1 | grep -E "^(ok|FAIL|---|\s+--- FAIL|\s+\S+_test.go)" | cut -c1-200
git checkout -q -- . && git clean -qfd sim && git checkout -q "$fix"
echo "### green at $fix"
nice -n 15 go test -tags=with_db -count=1 -run "$tests" -v ./sim/mage/ ./sim/conformance/ 2>&1 | grep -E "^(ok|FAIL|--- |\s+--- )" | cut -c1-200
echo "### suites"
pk="./sim/core/ ./sim/mage/... ./sim/druid/... ./sim/warlock/... ./sim/hunter/... ./sim/common/... ./sim/conformance/"
nice -n 15 go test -tags=with_db -p 1 -count=1 $pk 2>&1 | grep -E "^(ok|FAIL|---|\s+--- FAIL|panic)" | grep -v "no test files" | cut -c1-200
echo "### goldens"
n=0; for f in $(find sim -name "*.results.tmp"); do cp "$f" "${f%.results.tmp}.results"; echo "updated ${f%.results.tmp}.results"; n=$((n+1)); done
if [ $n -gt 0 ]; then nice -n 15 go test -tags=with_db -p 1 -count=1 $pk 2>&1 | grep -E "^(ok|FAIL)" | grep -v "no test files" | cut -c1-200; fi
echo "=====PATCH====="
git diff -- '*.results'
echo "=====END====="
git -C $repo worktree remove --force $wt
