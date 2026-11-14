#!/usr/bin/env bash
# fork_tests.sh <red sha> <fix sha> <test regex> <red pkg> <suite pkgs...>: on the box, in a throwaway
# worktree of /home/ci/wowsims-forever, run the regex red at <red sha> and green at <fix sha> in <red pkg>,
# then the suite packages whole; any suite golden that moved is rewritten and printed back as a patch.
# Run: ssh -o BatchMode=yes jh@192.168.1.182 "wsl -d forever-ci -- bash -l -s -- <args>" < fork_tests.sh
set -uo pipefail
red=$1; fix=$2; tests=${3//,/|}; redpk=$4; shift 4; pk="$*"
export PATH=/usr/local/go/bin:/home/ci/go/bin:$HOME/go/bin:$PATH
repo=/home/ci/wowsims-forever; wt=/home/ci/wf-sawar
git -C $repo fetch -q origin "+refs/heads/lane/sa-warrior-70938:refs/remotes/origin/lane/sa-warrior-70938" master
git -C $repo worktree remove --force $wt 2>/dev/null; git -C $repo worktree add -q --detach $wt "$red" || exit 1
cd $wt || exit 1
export GOMAXPROCS=5 GOFLAGS=-mod=mod
echo "### red at $red"
nice -n 15 go test -tags=with_db -count=1 -run "$tests" -v $redpk 2>&1 | grep -E "^(ok|FAIL|--- |\s+--- |\s+\S+_test.go)" | cut -c1-260
git checkout -q "$fix"
echo "### green at $fix"
nice -n 15 go test -tags=with_db -count=1 -run "$tests" -v $redpk 2>&1 | grep -E "^(ok|FAIL|--- |\s+--- )" | cut -c1-260
echo "### suites"
nice -n 15 go test -tags=with_db -p 1 -count=1 $pk 2>&1 | grep -E "^(ok|FAIL|---|\s+--- FAIL|\s+\S+_test.go|panic)" | grep -v "no test files" | cut -c1-260
echo "### goldens"
n=0; for f in $(find sim -name "*.results.tmp"); do cp "$f" "${f%.results.tmp}.results"; echo "updated ${f%.results.tmp}.results"; n=$((n+1)); done
if [ $n -gt 0 ]; then nice -n 15 go test -tags=with_db -p 1 -count=1 $pk 2>&1 | grep -E "^(ok|FAIL)" | grep -v "no test files" | cut -c1-220; fi
echo "=====PATCH====="
git diff -- '*.results'
echo "=====END====="
cd / && git -C $repo worktree remove --force $wt
