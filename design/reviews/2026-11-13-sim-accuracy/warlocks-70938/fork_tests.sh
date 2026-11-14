#!/usr/bin/env bash
# fork_tests.sh <branch> <sha> [go test -run pattern] [packages...]: on the box, in a throwaway worktree of
# /home/ci/wowsims-forever at <sha> (fetched from origin <branch>), run the fork's warlock, core and
# conformance tests (or the named packages), one job at a time.
set -uo pipefail
branch=$1; sha=$2; run=${3:-.}; shift 3 2>/dev/null; pk=${*:-"./sim/warlock/... ./sim/core/ ./sim/conformance/"}
export PATH=/usr/local/go/bin:/home/ci/go/bin:$HOME/go/bin:$PATH GOMAXPROCS=5 GOFLAGS=-mod=mod
repo=/home/ci/wowsims-forever; wt=/home/ci/wf-sawarlock
git -C $repo fetch -q origin "+refs/heads/$branch:refs/remotes/origin/$branch" master
git -C $repo worktree remove --force $wt 2>/dev/null; git -C $repo worktree add -q --detach $wt "$sha" || exit 1
cd $wt || exit 1
# shellcheck disable=SC2086
nice -n 15 go test -tags=with_db -p 1 -count=1 -run "$run" $pk 2>&1 | grep -v "no test files" | grep -E "^(ok|FAIL|---|\s+--- |\s+\S+_test.go|panic)|golden|mismatch" | head -80 | cut -c1-240
git -C $repo worktree remove --force $wt
