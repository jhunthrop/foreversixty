#!/usr/bin/env bash
# Tests for tools/land.sh, the integrator's landing command (docs/runbooks/landing.md). Plain
# bash, the shape of test_release_apps.sh: each case pushes a land branch from a scratch clone
# to a scratch bare "origin", runs the real script there with real git and gh stubbed (GH) to
# serve canned run lists, polls and job lists, and checks what reached origin.
#
#   bash tools/tests/test_land.sh
#
# Prints "ok N tests" and exits 0 when every assertion holds; names each failure and exits
# 1 otherwise.
set -uo pipefail

HERE=$(cd "$(dirname "$0")" && pwd)
readonly HERE
readonly LAND=$HERE/../land.sh

SCRATCH=$(mktemp -d)
trap 'rm -rf "$SCRATCH"' EXIT
PASSED=0
FAILED=0

assert() {
	local what=$1
	shift
	if "$@"; then
		PASSED=$((PASSED + 1))
	else
		FAILED=$((FAILED + 1))
		echo "FAIL: $what" >&2
	fi
}

equal() { [[ $1 == "$2" ]]; }
contains() { [[ $1 == *"$2"* ]]; }
lacks() { [[ $1 != *"$2"* ]]; }

commit() { git -C "$1" -c user.name=t -c user.email=t@t commit -q --allow-empty -m "$2"; }

# setup <name>: a bare origin whose main has one commit and a clone with land/1 one commit
# ahead of it, pushed; prints the clone's path. The clone holds the gh stub's fixtures in
# <clone>/gh: list.<n>.json is the run list the n-th poll sees (the highest n at or below
# the poll count), run.<id>.json what `gh run view <id>` prints.
setup() {
	local dir=$SCRATCH/$1
	git init -q --bare -b main "$dir/origin.git"
	git clone -q "$dir/origin.git" "$dir/clone" 2>/dev/null
	commit "$dir/clone" base
	git -C "$dir/clone" push -q origin HEAD:main
	git -C "$dir/clone" checkout -q -b land/1
	commit "$dir/clone" batch
	git -C "$dir/clone" push -q origin land/1
	mkdir -p "$dir/clone/gh"
	cat >"$dir/clone/gh/gh" <<EOF
#!/usr/bin/env bash
echo "gh \$*" >>"$dir/clone/gh/calls"
if [ "\$1 \$2" = "run list" ]; then
  n=\$(( \$(cat "$dir/clone/gh/polls" 2>/dev/null || echo 0) + 1 ))
  echo "\$n" >"$dir/clone/gh/polls"
  while [ "\$n" -gt 0 ] && [ ! -e "$dir/clone/gh/list.\$n.json" ]; do n=\$((n - 1)); done
  cat "$dir/clone/gh/list.\$n.json"
elif [ "\$1 \$2" = "run view" ]; then
  cat "$dir/clone/gh/run.\$3.json"
elif [ "\$1 \$2" = "workflow run" ]; then
  [ ! -e "$dir/clone/gh/dispatch-fails" ]
else
  exit 1
fi
EOF
	chmod +x "$dir/clone/gh/gh"
	echo "$dir/clone"
}

# run_json <id> <workflow> <status> <conclusion> <job:status:conclusion...>; the run's event
# is EVENT (push by default).
run_json() {
	local id=$1 name=$2 status=$3 conclusion=$4 jobs='' job event=${EVENT:-push}
	shift 4
	for job in "$@"; do
		IFS=: read -r jname jstatus jconclusion <<<"$job"
		jobs+="${jobs:+,}{\"name\":\"$jname\",\"status\":\"$jstatus\",\"conclusion\":\"$jconclusion\",\"url\":\"https://ci/$id/$jname\"}"
	done
	printf '{"databaseId":%s,"name":"%s","status":"%s","conclusion":"%s","event":"%s","jobs":[%s]}' \
		"$id" "$name" "$status" "$conclusion" "$event" "$jobs"
}

# fixture <clone> <poll> <run json...>: the run list of that poll and each run's view.
fixture() {
	local clone=$1 poll=$2 list='' run id
	shift 2
	for run in "$@"; do
		id=$(jq -r .databaseId <<<"$run")
		jq -c 'del(.jobs)' <<<"$run" >"$clone/gh/tmp"
		list+="${list:+,}$(cat "$clone/gh/tmp")"
		echo "$run" >"$clone/gh/run.$id.json"
	done
	echo "[$list]" >"$clone/gh/list.$poll.json"
}

# land <clone> <args...>: runs the script in <clone>; output in <clone>/out, status returned.
land() {
	local clone=$1
	shift
	(cd "$clone" && GH="$clone/gh/gh" LAND_POLL_SECONDS=0 LAND_SETTLE_SECONDS=0 \
		LAND_TIMEOUT_SECONDS=${TIMEOUT:-60} bash "$LAND" "$@") >"$clone/out" 2>&1
}

# with_workflows <clone>: land/1's batch now changes api/ only, and carries an api workflow
# (push on land/**, paths api/**) and a docs workflow (paths docs/**) the batch does not touch.
with_workflows() {
	local clone=$1
	mkdir -p "$clone/.github/workflows" "$clone/api"
	printf 'name: api\non:\n  push:\n    branches: [main, "land/**"]\n    paths: ["api/**"]\n  workflow_dispatch:\n' \
		>"$clone/.github/workflows/api.yml"
	printf 'name: docs\non:\n  push:\n    branches: [main, "land/**"]\n    paths: ["docs/**"]\n' \
		>"$clone/.github/workflows/docs.yml"
	echo change >"$clone/api/x.go"
	git -C "$clone" add -A
	git -C "$clone" -c user.name=t -c user.email=t@t commit -q --amend -m batch
	git -C "$clone" push -q -f origin land/1
}

origin_ref() { git -C "$1/../origin.git" rev-parse -q --verify "$2" 2>/dev/null; }

# --- usage: only a land/<n> branch ---
c=$(setup usage)
land "$c" main
assert "a branch outside land/ is refused with usage (64)" equal $? 64
land "$c"
assert "no branch is refused with usage (64)" equal $? 64

# --- green: waits for the running run, then fast-forwards main and deletes the branch ---
c=$(setup green)
sha=$(git -C "$c" rev-parse HEAD)
fixture "$c" 1 "$(run_json 11 web in_progress '' build:completed:success e2e:in_progress:)" \
	"$(run_json 12 api completed success test:completed:success deploy:completed:skipped)"
fixture "$c" 2 "$(run_json 11 web completed success build:completed:success e2e:completed:success)" \
	"$(run_json 12 api completed success test:completed:success deploy:completed:skipped)"
land "$c" land/1
assert "a green landing exits 0" equal $? 0
assert "main is the land branch's sha" equal "$(origin_ref "$c" refs/heads/main)" "$sha"
assert "the land branch is deleted" equal "$(origin_ref "$c" refs/heads/land/1)" ''
assert "it polled until the running run completed" equal "$(cat "$c/gh/polls")" 2
assert "it asks for the push runs of the branch's head sha" \
	contains "$(cat "$c/gh/calls")" "run list --branch land/1 --commit $sha"
assert "a skipped job is green" lacks "$(cat "$c/out")" 'failed'

# --- launch mode: skipped test jobs, and whole runs whose every job skipped, are green ---
c=$(setup launch)
sha=$(git -C "$c" rev-parse HEAD)
fixture "$c" 1 "$(run_json 51 web completed success build:completed:success \
	'e2e (1):completed:skipped' real-engine:completed:skipped lhci:completed:skipped deploy:completed:skipped)" \
	"$(run_json 52 api completed success check:completed:success test:completed:skipped image:completed:skipped)" \
	"$(run_json 53 ios completed skipped build-test:completed:skipped)" \
	"$(run_json 54 coach completed skipped)"
land "$c" land/1
assert "a launch-mode landing with skipped jobs exits 0" equal $? 0
assert "main is the launch-mode sha" equal "$(origin_ref "$c" refs/heads/main)" "$sha"
assert "no skipped job or run is listed as failed" lacks "$(cat "$c/out")" 'failed'

# --- red: a failed job stops the wait at once and is listed, main and the branch stay ---
c=$(setup red)
main_before=$(origin_ref "$c" refs/heads/main)
fixture "$c" 1 "$(run_json 21 web in_progress '' 'e2e (2):completed:failure' 'e2e (1):in_progress:')" \
	"$(run_json 22 sim completed failure test:completed:success rotations:completed:timed_out)" \
	"$(run_json 23 tools completed startup_failure)"
land "$c" land/1
assert "a red landing exits 1" equal $? 1
out=$(cat "$c/out")
assert "the failed job is listed while its run still runs" contains "$out" 'web / e2e (2) (failure) https://ci/21/e2e (2)'
assert "a timed-out job is listed" contains "$out" 'sim / rotations (timed_out)'
assert "a run that never started a job is listed" contains "$out" 'tools (startup_failure)'
assert "a green job is not listed" lacks "$out" 'sim / test'
assert "it stopped on the first poll" equal "$(cat "$c/gh/polls")" 1
assert "main did not move" equal "$(origin_ref "$c" refs/heads/main)" "$main_before"
assert "the land branch is kept" equal "$(origin_ref "$c" refs/heads/land/1)" "$(git -C "$c" rev-parse HEAD)"

# --- a cancelled run is red ---
c=$(setup cancelled)
fixture "$c" 1 "$(run_json 31 ios completed cancelled build-test:completed:cancelled)"
land "$c" land/1
assert "a cancelled job fails the landing" equal $? 1
assert "the cancelled job is listed" contains "$(cat "$c/out")" 'ios / build-test (cancelled)'

# --- main moved: the push is refused, the landing says a merge is needed (2) ---
c=$(setup moved)
fixture "$c" 1 "$(run_json 41 web completed success build:completed:success)"
git clone -q "$c/../origin.git" "$c/../other" 2>/dev/null
commit "$c/../other" "someone else's landing"
git -C "$c/../other" push -q origin HEAD:main
moved=$(origin_ref "$c" refs/heads/main)
land "$c" land/1
assert "a moved main exits 2" equal $? 2
assert "it says a merge is needed" contains "$(cat "$c/out")" 'main moved'
assert "main keeps the other landing" equal "$(origin_ref "$c" refs/heads/main)" "$moved"
assert "the land branch is kept for the merge" equal "$(origin_ref "$c" refs/heads/land/1)" "$(git -C "$c" rev-parse HEAD)"

# --- no runs ever start: the wait times out rather than landing unchecked ---
c=$(setup none)
echo '[]' >"$c/gh/list.1.json"
TIMEOUT=0 land "$c" land/1
assert "no runs at all is a timeout, not a landing" equal $? 1
assert "the timeout is reported" contains "$(cat "$c/out")" 'timed out'
assert "main did not move with no checks" lacks "$(origin_ref "$c" refs/heads/main)" "$(git -C "$c" rev-parse HEAD)"

# --- a run whose jobs cannot be read is not green: the landing stops rather than guessing ---
c=$(setup unreadable)
fixture "$c" 1 "$(run_json 51 web completed failure build:completed:failure)"
rm "$c/gh/run.51.json"
land "$c" land/1
assert "a run whose jobs cannot be read fails the landing" equal $? 1
assert "it names the run it could not read" contains "$(cat "$c/out")" 'could not read run 51'
assert "main did not move on an unreadable run" lacks "$(origin_ref "$c" refs/heads/main)" "$(git -C "$c" rev-parse HEAD)"

# --- an expected workflow that never started is dispatched, then waited for ---
c=$(setup dispatch)
with_workflows "$c"
sha=$(git -C "$c" rev-parse HEAD)
fixture "$c" 1 "$(run_json 61 companion completed success build:completed:success)"
fixture "$c" 2 "$(run_json 61 companion completed success build:completed:success)" \
	"$(EVENT=workflow_dispatch run_json 62 api completed success test:completed:success)"
land "$c" land/1
assert "a dispatched expected workflow lands once green" equal $? 0
assert "it dispatched the api workflow on the land branch" contains "$(cat "$c/gh/calls")" "workflow run api.yml --ref land/1"
assert "it did not dispatch a workflow whose paths the batch misses" lacks "$(cat "$c/gh/calls")" "docs.yml"
assert "it says why it dispatched" contains "$(cat "$c/out")" 'api never started on land/1'
assert "main is the dispatched landing's sha" equal "$(origin_ref "$c" refs/heads/main)" "$sha"

# --- an expected workflow that never starts, even dispatched, is a timeout, not a landing ---
c=$(setup never)
with_workflows "$c"
fixture "$c" 1 "$(run_json 71 companion completed success build:completed:success)"
TIMEOUT=0 land "$c" land/1
assert "a missing expected workflow fails the landing" equal $? 1
assert "it names the workflow that never started" contains "$(cat "$c/out")" 'never started: api'
assert "main did not move without the api run" lacks "$(origin_ref "$c" refs/heads/main)" "$(git -C "$c" rev-parse HEAD)"

# --- a dispatch that fails stops the landing ---
c=$(setup undispatchable)
with_workflows "$c"
fixture "$c" 1 "$(run_json 81 companion completed success build:completed:success)"
touch "$c/gh/dispatch-fails"
land "$c" land/1
assert "a failed dispatch fails the landing" equal $? 1
assert "it names the workflow it could not dispatch" contains "$(cat "$c/out")" 'could not dispatch api.yml'

# --- the expected-workflow script runs through python3, whatever its file mode ---
c=$(setup nomode)
with_workflows "$c"
cp "$HERE/../land_workflows.py" "$c/land_workflows.py"
chmod -x "$c/land_workflows.py"
fixture "$c" 1 "$(run_json 91 api completed success test:completed:success)"
(cd "$c" && GH="$c/gh/gh" LAND_POLL_SECONDS=0 LAND_SETTLE_SECONDS=0 LAND_TIMEOUT_SECONDS=60 \
	LAND_EXPECTED="$c/land_workflows.py" bash "$LAND" land/1) >"$c/out" 2>&1
assert "a non-executable land_workflows.py still lands" equal $? 0

# --- a branch origin does not have ---
c=$(setup missing)
land "$c" land/404
assert "a missing land branch fails" equal $? 1
assert "it names the missing branch" contains "$(cat "$c/out")" 'no branch land/404'

if ((FAILED > 0)); then
	echo "$FAILED of $((PASSED + FAILED)) tests failed" >&2
	exit 1
fi
echo "ok $PASSED tests"
