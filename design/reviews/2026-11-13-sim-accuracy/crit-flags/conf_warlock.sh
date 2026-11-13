#!/usr/bin/env bash
# conf_warlock.sh <pin> <fix>: TestConformance/warlock at both shas, with its failure text.
set -uo pipefail
export PATH=/usr/local/go/bin:/home/ci/go/bin:$HOME/go/bin:$PATH GOMAXPROCS=5 GOFLAGS=-mod=mod
repo=/home/ci/wowsims-forever; wt=/home/ci/wf-critflags
git -C $repo fetch -q origin "+refs/heads/lane/crit-flags:refs/remotes/origin/lane/crit-flags" master
for sha in "$@"; do
	git -C $repo worktree remove --force $wt 2>/dev/null; git -C $repo worktree add -q --detach $wt "$sha" || exit 1
	echo "### $sha"
	(cd $wt && nice -n 15 go test -tags=with_db -count=1 -run 'TestConformance/warlock' ./sim/conformance/ 2>&1 | grep -v "^\s*$" | head -40 | cut -c1-240)
	(cd $wt && git status --short | head; git diff --stat | tail -3; git diff | grep "^[-+][^-+]" | head -20 | cut -c1-240)
done
git -C $repo worktree remove --force $wt
