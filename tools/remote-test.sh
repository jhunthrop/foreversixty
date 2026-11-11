#!/usr/bin/env bash
# tools/remote-test.sh [options] <git-ref | worktree-path> -- <command...>
#
# Runs a test command for one commit on the remote test box (the owner's Windows PC, its
# dedicated WSL2 distro `forever-ci` on the D: NVMe) instead of on the laptop. It pushes the
# commit to a bare repo inside the distro, checks it out into a clean per-run worktree, links
# in node_modules from a cache keyed on web/package-lock.json and the Node version, gives the
# run its own Postgres database and Playwright port, runs the command, streams its output and
# exits with its status. Only committed work travels: commit a lane's changes first (the
# script warns when the worktree it was given is dirty). No secret ever goes to the box; the
# database is local and throwaway.
#
#   tools/remote-test.sh HEAD -- npm test                                   # vitest, from web/
#   tools/remote-test.sh ../forever-lane-x -- npm run check
#   tools/remote-test.sh origin/main -- npx playwright test tests/e2e/guild-raids.spec.ts
#   tools/remote-test.sh --workers 10 HEAD -- npx playwright test           # the full suite
#   tools/remote-test.sh --dir api HEAD -- go test -race -p 1 ./internal/guilds/
#   tools/remote-test.sh --dir companion HEAD -- go test -tags nogui ./...
#   tools/remote-test.sh --dir addon HEAD -- busted
#
# Options:
#   --dir <path>     where the command runs, relative to the repo root (default: web; `.` = root)
#   --workers <n>    Playwright workers, appended as --workers=<n> when the command is
#                    `npx playwright test ...` without its own --workers (default: 6)
#   --pin-engine     rewrite sim/go.mod's engine replace to the pinned fork sha, as CI's
#                    .github/actions/pin-engine does; on by default when the command is `go ...`
#   --engine-sha <sha>  pin the engine to this pushed fork sha instead of sim/enginever's
#                    (a lane's unpinned fork branch); implies --pin-engine
#   --no-pin-engine  never pin (a Go command that does not load the sim module)
#   --keep           keep the run's worktree, ref and database for inspection
#
# In the run: CI=1, TZ=UTC, TEST_DATABASE_URL and DATABASE_URL point at the run's own database
# on the box's Postgres 16 (localhost only, trust auth), E2E_PORT is a free port so runs never
# collide, and FOREVER_DATA is left to the command (Playwright and `npm test` default to the
# fixture). The npm hooks (pretest, prebuild) run the data syncs exactly as they do locally.
#
# Environment (defaults): REMOTE_TEST_HOST (jh@192.168.1.182), REMOTE_TEST_DISTRO (forever-ci).
# Exit status: the command's own; 64 usage; 65 the ref does not resolve; 66 the push failed.
# Interrupting the script ends the ssh session; a run left behind is pruned after two days.
set -euo pipefail

readonly HOST=${REMOTE_TEST_HOST:-jh@192.168.1.182}
readonly DISTRO=${REMOTE_TEST_DISTRO:-forever-ci}
readonly REMOTE_REPO=/home/ci/forever-ci/repo.git
readonly DEFAULT_WORKERS=6

usage() {
	sed -n '2,/^set -euo/p' "$0" | sed '$d; s/^# \{0,1\}//' >&2
	exit 64
}

dir=web workers=$DEFAULT_WORKERS pin=auto keep=0 engine_sha=
while (($#)); do
	case $1 in
	--dir) dir=${2:?--dir needs a path}; shift 2 ;;
	--workers) workers=${2:?--workers needs a number}; shift 2 ;;
	--pin-engine) pin=1; shift ;;
	--no-pin-engine) pin=0; shift ;;
	--engine-sha) engine_sha=${2:?--engine-sha needs a sha}; pin=1; shift 2 ;;
	--keep) keep=1; shift ;;
	-h | --help) usage ;;
	--) usage ;;
	-*) echo "remote-test: unknown option $1" >&2; usage ;;
	*) break ;;
	esac
done
(($# >= 3)) && [[ $2 == -- ]] || usage
readonly target=$1
shift 2
[[ $workers =~ ^[1-9][0-9]*$ ]] || { echo "remote-test: --workers must be a positive number" >&2; exit 64; }
[[ $dir =~ ^[A-Za-z0-9._/-]+$ && $dir != *..* ]] || { echo "remote-test: --dir must be a plain repo-relative path" >&2; exit 64; }
[[ -z $engine_sha || $engine_sha =~ ^[0-9a-f]{7,40}$ ]] || { echo "remote-test: --engine-sha must be a hex commit sha" >&2; exit 64; }

# The commit to run, from a worktree path or a ref.
if [[ -d $target ]]; then
	repo=$target
	sha=$(git -C "$target" rev-parse --verify HEAD) || exit 65
	[[ -z $(git -C "$target" status --porcelain --untracked-files=no) ]] ||
		echo "remote-test: $target has uncommitted changes; only HEAD ($sha) runs" >&2
else
	repo=.
	sha=$(git rev-parse --verify --quiet "$target^{commit}") || {
		echo "remote-test: $target is not a commit" >&2
		exit 65
	}
fi

cmd=("$@")
if [[ ${cmd[0]:-} == npx && ${cmd[1]:-} == playwright && ${cmd[2]:-} == test ]] &&
	! printf '%s\n' "${cmd[@]}" | grep -qE '^(--workers|-j)'; then
	cmd+=("--workers=$workers")
fi
if [[ $pin == auto ]]; then
	[[ ${cmd[0]} == go ]] && pin=1 || pin=0
fi

run_id="$(date -u +%Y%m%d%H%M%S)-${sha:0:9}-$$"
readonly run_id ref="refs/heads/run-$run_id"

echo "remote-test: $sha -> $DISTRO on ${HOST#*@} (run $run_id, in $dir): ${cmd[*]}" >&2
git -C "$repo" push --quiet --receive-pack="wsl -d $DISTRO -- git-receive-pack" \
	"$HOST:$REMOTE_REPO" "$sha:$ref" || {
	echo "remote-test: pushing $sha to the box failed" >&2
	exit 66
}

# Every value crosses cmd.exe and wsl.exe, so each travels as base64 (no quoting to survive).
b64() { printf '%s' "$1" | base64 | tr -d '\n'; }
cmd_line=$(printf '%q ' "${cmd[@]}")

# shellcheck disable=SC2016 # the body below runs on the box, not here
remote_body='
main() {
set -euo pipefail
dec() { printf "%s" "$1" | base64 -d; }
run_id=$(dec "$1") sha=$(dec "$2") dir=$(dec "$3") pin=$(dec "$4") keep=$(dec "$5") cmd_line=$(dec "$6") engine_sha=$(dec "${7:-}")
root=$HOME/forever-ci repo=$root/repo.git wt=$root/runs/$run_id db=ft_${run_id//-/_}
export TZ=UTC CI=1

prune_old() {
	find "$root/runs" -mindepth 1 -maxdepth 1 -type d -mtime +2 -exec rm -rf {} + 2>/dev/null || true
	git -C "$repo" worktree prune
	git -C "$repo" for-each-ref --format="%(refname)" refs/heads/run-* | while read -r r; do
		id=${r#refs/heads/run-}
		[[ -d $root/runs/$id ]] || git -C "$repo" update-ref -d "$r"
	done
}
cleanup() {
	((keep)) && { echo "remote-test: kept $wt, ref run-$run_id, database $db" >&2; return; }
	git -C "$repo" worktree remove --force "$wt" 2>/dev/null || rm -rf "$wt"
	git -C "$repo" update-ref -d "refs/heads/run-$run_id" || true
	dropdb -h localhost -U forever --if-exists "$db" || true
}

prune_old
git -C "$repo" worktree add --quiet --detach "$wt" "$sha"
trap cleanup EXIT

# node_modules: one cache entry per lockfile and Node version, hard-linked into the run.
if [[ -f $wt/web/package-lock.json ]]; then
	key=$( { node --version; cat "$wt/web/package-lock.json"; } | sha256sum | cut -c1-16)
	cache=$root/cache/node_modules/$key
	(
		flock 9
		if [[ ! -d $cache/node_modules ]]; then
			echo "remote-test: building the node_modules cache $key" >&2
			rm -rf "$cache.tmp" && mkdir -p "$cache.tmp"
			cp "$wt/web/package.json" "$wt/web/package-lock.json" "$cache.tmp/"
			(cd "$cache.tmp" && npm ci --no-audit --no-fund --loglevel=error)
			pw=$(node -p "require(\"$cache.tmp/package-lock.json\").packages[\"node_modules/@playwright/test\"].version")
			(cd "$cache.tmp" && npx playwright install chromium >/dev/null) ||
				echo "remote-test: playwright $pw chromium install failed (system deps: sudo npx playwright install-deps)" >&2
			rm -rf "$cache" && mv "$cache.tmp" "$cache"
		fi
	) 9>"$root/cache/node_modules.lock"
	cp -al "$cache/node_modules" "$wt/web/node_modules"
fi

if ((pin)); then
	(
		cd "$wt/sim"
		engine=${engine_sha:-$(sed -n "s/.*Version = \"\(.*\)\"/\1/p" enginever/version.go)}
		echo "remote-test: engine pinned to $engine" >&2
		[[ -n $engine ]] || { echo "remote-test: sim/enginever/version.go has no Version" >&2; exit 1; }
		go mod edit -replace "github.com/wowsims/classic=github.com/jhunthrop/wowsims-forever@$engine"
		go mod tidy || { rm -rf "$(go env GOMODCACHE)/cache/vcs"; go mod tidy; }
	)
fi

createdb -h localhost -U forever "$db"
export TEST_DATABASE_URL="postgres://forever:forever@localhost:5432/$db?sslmode=disable"
export DATABASE_URL=$TEST_DATABASE_URL
E2E_PORT=$(python3 -c "import socket; s=socket.socket(); s.bind((\"127.0.0.1\", 0)); print(s.getsockname()[1])")
export E2E_PORT

cd "$wt/$dir"
set +e
start=$SECONDS
eval "$cmd_line"
status=$?
echo "remote-test: exit $status in $((SECONDS - start))s ($sha)" >&2
exit $status
}
# Parsed whole before it runs, so a command that reads stdin cannot eat the script.
main "$@" </dev/null
'

# shellcheck disable=SC2029 # the arguments are meant to expand here, on the Mac
printf '%s' "$remote_body" | ssh -o BatchMode=yes "$HOST" \
	"wsl -d $DISTRO -- bash -l -s -- $(b64 "$run_id") $(b64 "$sha") $(b64 "$dir") $(b64 "$pin") $(b64 "$keep") $(b64 "$cmd_line") $(b64 "$engine_sha")"
