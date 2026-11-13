#!/usr/bin/env bash
# fork_tests.sh <pin sha> <fix sha> <test regex> <packages...>: on the box, in a throwaway worktree of
# /home/ci/wowsims-forever, run the lane's tests red at the pin (the fix's test files dropped in) and green
# at the fix, then the named packages whole, with any suite golden rewritten and printed back as a patch.
# Run: ssh -o BatchMode=yes jh@192.168.1.182 "wsl -d forever-ci -- bash -l -s -- <pin> <fix> <regex> <pkgs>" < fork_tests.sh
set -uo pipefail
pin=$1; fix=$2; tests=$3; shift 3; pk="$*"
export PATH=/usr/local/go/bin:/home/ci/go/bin:$HOME/go/bin:$PATH
repo=/home/ci/wowsims-forever; wt=/home/ci/wf-arcshadow
git -C $repo fetch -q origin "+refs/heads/lane/sim-arcshadow:refs/remotes/origin/lane/sim-arcshadow" master
git -C $repo worktree remove --force $wt 2>/dev/null; git -C $repo worktree add -q --detach $wt "$fix" || exit 1
cd $wt || exit 1
export GOMAXPROCS=5 GOFLAGS=-mod=mod
files=$(git diff --name-only "$pin" "$fix" -- '*_test.go')
echo "### red at $pin (test files from $fix: $files)"
mkdir -p /tmp/arcshadow-tests && for f in $files; do mkdir -p /tmp/arcshadow-tests/$(dirname $f); git show "$fix:$f" > /tmp/arcshadow-tests/$f; done
git checkout -q "$pin" && for f in $files; do cp /tmp/arcshadow-tests/$f $f; done
nice -n 15 go test -tags=with_db -count=1 -run "$tests" $pk 2>&1 | grep -E "^(ok|FAIL|---|\s+--- FAIL|\s+\S+_test.go)" | cut -c1-220
git checkout -q -- . && git clean -qfd sim && git checkout -q "$fix"
echo "### green at $fix"
nice -n 15 go test -tags=with_db -count=1 -run "$tests" -v $pk 2>&1 | grep -E "^(ok|FAIL|--- |\s+--- )" | cut -c1-220
echo "### suites"
nice -n 15 go test -tags=with_db -p 1 -count=1 $pk 2>&1 | grep -E "^(ok|FAIL|---|\s+--- FAIL|panic)" | grep -v "no test files" | cut -c1-220
echo "### goldens"
n=0; for f in $(find sim -name "*.results.tmp"); do cp "$f" "${f%.results.tmp}.results"; echo "updated ${f%.results.tmp}.results"; n=$((n+1)); done
if [ $n -gt 0 ]; then nice -n 15 go test -tags=with_db -p 1 -count=1 $pk 2>&1 | grep -E "^(ok|FAIL)" | grep -v "no test files" | cut -c1-220; fi
echo "=====PATCH====="
git diff -- '*.results'
echo "=====END====="
cd / && git -C $repo worktree remove --force $wt
