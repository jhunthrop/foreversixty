#!/usr/bin/env bash
# tools/activate-build.sh <build> [--no-commit [--no-gate] | --no-publish]
#
# Moves the site, the addon and the guides to a client build that data.yml's fetch job
# has already landed under data/builds/<build>/. This is the scripted form of the
# 2026-10-14 move to 1.60.1.70374 (design/reviews/2026-10-14-build-70374-move.md) and of
# step 10 of data/README.md's "New build checklist"; .github/workflows/client.yml runs it
# the morning a new World of Warcraft: Forever client appears on wago.tools.
#
#   1. web/src/data/active-build.json names the new build.
#   2. The ranking's outputs that no fetch produces are carried forward from the build
#      that was active (bis/, dps-per-point.json), each only where the new build has
#      none, so the site has picks the same morning. The ranking (bis.yml) replaces
#      bis/ and re-measures dps-per-point.json once the pinned engine is stamped for the
#      new build: dispatched with force at the end when it already is, else by
#      engine.yml's green pin of the follow (bis.yml itself refuses a mismatched
#      engine). addon-builds.json is NOT carried: it is rebuilt from the moved guides.
#   3. The build's stat-weights copy, addon-data.json and addon/ForeverSixty/Data.lua
#      are rendered for the new build (the documented exception to "nightly owns
#      generated data": Data.lua tracks the ACTIVE build).
#   4. Every guide's FS1 build code is re-stamped (web/scripts/reencode-guide-builds.mjs,
#      which proves each code legal against the new trees first) and the guides' data
#      links follow.
#   5. The three addon TOCs take the Interface number the client version implies, the
#      addon's patch version is bumped with a changelog line, and the tag the push
#      makes releases it to CurseForge and Wago, so players get a Data.lua and an
#      Interface that match the client they are about to run.
#      The log coach's 28 spec rule files (data/curated/coach) are restamped for the new
#      build (pipeline coach-rules); a spell the new tables lost stops the move.
#   6. The --check gates data.yml runs are run here first; a drift stops the move
#      before anything is committed.
#   7. The test gate (tools/activation-gate.sh): the web vitest suite web.yml's deploy
#      waits on and the API packages that read the active build, against the moved tree.
#      A red test stops the move before the commit, so the addon is never tagged for a
#      client the site then refuses to deploy; the failing tests go to the run summary.
#
# --no-commit leaves the working tree changed and stops (inspect with git diff).
# --no-publish commits and stops (no push, no tag, no ranking dispatch).
# --no-gate skips step 7; only with --no-commit, for a quick local look at a move.
#
# On a failure the script names the step, the reason (the step's output, or the failing
# tests) and the exact hand command that finishes or redoes the move: on stderr,
# and in the run summary and an error annotation under GitHub Actions. client.yml passes
# PRODUCT (the wago.tools key) so the hand commands name it.
set -euo pipefail

usage() { echo "usage: $0 <build> [--no-commit [--no-gate] | --no-publish]" >&2; exit 2; }
[ $# -ge 1 ] || usage
NEW="$1"; shift
MODE=publish
GATE=on
for arg in "$@"; do
  case "$arg" in
    --no-commit) MODE=inspect ;;
    --no-publish) MODE=commit ;;
    --no-gate) GATE=off ;;
    *) usage ;;
  esac
done
[ "$GATE" = on ] || [ "$MODE" = inspect ] || { echo "--no-gate is only for --no-commit: a move that commits is always gated" >&2; exit 2; }
[[ "$NEW" =~ ^[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+$ ]] || { echo "$NEW is not a four-part client build" >&2; exit 2; }

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

# --- Failure reporting ---------------------------------------------------------------
# step <name> <hand command>: what runs next and how a person finishes the move if it
# fails. run <command...>: runs it with its output kept, so a failure can quote it.
# die <reason>: stops at the current step with that reason.
REDO="gh workflow run client.yml -f build=$NEW${PRODUCT:+ -f product=$PRODUCT}"
STEP="reading the active build"; HAND="$REDO"; REASON=""; UNDO=""
STEP_LOG="$(mktemp)"
step() { STEP="$1"; HAND="$2"; REASON=""; : > "$STEP_LOG"; echo "== $STEP"; }
die() { REASON="$1"; echo "$1" >&2; exit 1; }
run() {
  local status=0
  "$@" 2>&1 | tee "$STEP_LOG" || status=$?
  [ "$status" -eq 0 ] || REASON="$(quote_log "$STEP_LOG")"
  return "$status"
}
# quote_log <file>: all of a short log; the head (which usually says what failed) and the
# tail (where the error usually is) of a long one.
quote_log() {
  local lines
  lines=$(wc -l < "$1" | tr -d ' ')
  if [ "$lines" -le 100 ]; then cat "$1"; return; fi
  head -n 20 "$1"; echo "... ($((lines - 70)) lines omitted; the job log has them) ..."; tail -n 50 "$1"
}
in_dir() { local dir="$1"; shift; (cd "$dir" && "$@"); }
report_failure() {
  local status=$? summary
  rm -f "$STEP_LOG"
  [ "$status" -ne 0 ] && [ "$status" -ne 2 ] || return 0
  summary="$(mktemp)"
  {
    echo "## Client move to $NEW stopped"
    echo
    echo "**Step:** $STEP"
    echo
    echo "**Reason:**"
    echo
    echo '```'
    echo "${REASON:-exit status $status; the log above has the detail}"
    echo '```'
    echo
    echo "**To finish or redo the move:** $HAND"
    if [ -n "$UNDO" ] && [ "${GITHUB_ACTIONS:-}" != true ]; then
      echo
      echo "**To undo the half-made move in this checkout:** $UNDO"
    fi
  } > "$summary"
  echo >&2; cat "$summary" >&2
  if [ -n "${GITHUB_STEP_SUMMARY:-}" ]; then cat "$summary" >> "$GITHUB_STEP_SUMMARY"; fi
  if [ "${GITHUB_ACTIONS:-}" = true ]; then
    echo "::error title=client move to $NEW stopped at: $STEP::$(head -n 1 <<<"${REASON:-exit status $status}") -- to finish: $HAND"
  fi
  rm -f "$summary"
}
trap report_failure EXIT

ACTIVE_JSON=web/src/data/active-build.json
ACTIVE="$(sed -n 's/.*"build"[[:space:]]*:[[:space:]]*"\([^"]*\)".*/\1/p' "$ACTIVE_JSON")"
[ -n "$ACTIVE" ] || die "$ACTIVE_JSON names no build"
if [ "$NEW" = "$ACTIVE" ]; then echo "$NEW is already the active build; nothing to do"; exit 0; fi

NEW_DIR="data/builds/$NEW"; OLD_DIR="data/builds/$ACTIVE"
step "the fetched tables for $NEW" \
  "gh workflow run data.yml -f product=${PRODUCT:-<wago product key>} -f build=$NEW, then $REDO"
for f in items.json simdb.bin talents.json levels.json spellranks.json loot.json manifest.json; do
  [ -e "$NEW_DIR/$f" ] || die "$NEW_DIR/$f is missing: data.yml's fetch has not landed $NEW"
done

# The Interface number Blizzard's TOC convention gives this version (1.60.1 -> 16001).
IFS=. read -r MAJOR MINOR PATCH _ <<<"$NEW"
NEW_IFACE=$((10#$MAJOR * 10000 + 10#$MINOR * 100 + 10#$PATCH))
OLD_IFACE="$(sed -n 's/^## Interface: *\([0-9]*\).*/\1/p' addon/ForeverSixty/ForeverSixty.toc)"

echo "activating $NEW (was $ACTIVE); Interface $OLD_IFACE -> $NEW_IFACE"
MOVED_PATHS=("$ACTIVE_JSON" "$NEW_DIR" addon/ForeverSixty addon/ForeverSixtyData addon/ForeverSixtyInbox data/curated/coach web/src/content/guides api/internal/bnetbuild/testdata/era-kiloz.fs1)
UNDO="git restore --staged --worktree -- ${MOVED_PATHS[*]} && git clean -fd -- $NEW_DIR"
FIX_THEN_REDO="fix what the reason names on main, then $REDO"

# 2. Carry the ranking's outputs forward where the new build has none.
step "carrying the picks forward from $ACTIVE (pipeline carry-bis)" "$FIX_THEN_REDO"
carried=()
if [ ! -d "$NEW_DIR/bis" ] && [ -d "$OLD_DIR/bis" ]; then
  # The carried files now stand for the new build (web/src/lib/bis reads them from its
  # directory and expects their build field to agree); generated_at and engine_version
  # keep saying which ranking they came from. A pick or alternative the new client's
  # items.json lacks is dropped with a report line (pipeline/bis_carry.py).
  run in_dir data uv run python -m pipeline carry-bis --from "$ACTIVE" --to "$NEW"; carried+=("bis/")
fi
for f in dps-per-point.json; do
  if [ ! -e "$NEW_DIR/$f" ] && [ -e "$OLD_DIR/$f" ]; then cp "$OLD_DIR/$f" "$NEW_DIR/$f"; carried+=("$f"); fi
done
[ ${#carried[@]} -eq 0 ] || echo "carried forward from $ACTIVE: ${carried[*]}"

# 1. The active build.
printf '{\n  "build": "%s"\n}\n' "$NEW" > "$ACTIVE_JSON"

# 3. The guides: FS1 codes carried onto the new trees by talent name (the script
#    exits non-zero and writes nothing when a tree changed shape in a way a person
#    has to resolve, which stops the move before anything is committed), data
#    links moved.
step "re-encoding the guide builds onto $NEW's talent trees" \
  "re-spend the builds named above against $NEW's trees (web/src/content/guides, codes stamped FS1:$NEW), then from main with those edits in the tree run: tools/activate-build.sh $NEW (it commits them with the move, gates, pushes, tags and dispatches); or land them on main and $REDO"
run node web/scripts/reencode-guide-builds.mjs
grep -rl "foreversixty.gg/data/$ACTIVE/" web/src/content/guides | while read -r f; do
  sed -i.bak "s#foreversixty.gg/data/$ACTIVE/#foreversixty.gg/data/$NEW/#g" "$f" && rm -f "$f.bak"
done

# 4. The addon's default builds, from the guides just moved, then the build's
#    weights copy, addon-data.json and the addon's Data.lua.
step "the addon's default builds (sim/cmd/addon-builds)" "$FIX_THEN_REDO"
run go run ./sim/cmd/addon-builds -repo-root .
step "the weights copy, addon-data.json and Data.lua for $NEW" "$FIX_THEN_REDO"
run in_dir data uv run python -m pipeline weights --build "$NEW"
run in_dir data uv run python -m pipeline addon-data --build "$NEW"

# 4a. The log coach's rule files, restamped against the new build's tables.
step "the log coach's rule files (pipeline coach-rules)" \
  "replace the spells the reason names in data/curated/coach for $NEW's tables, land that on main, then $REDO"
run in_dir data uv run python -m pipeline coach-rules --build "$NEW"

# 4b. The API's Battle.net import fixture is stamped with the active build.
step "the API's import fixture stamp" "$FIX_THEN_REDO"
sed -i.bak "s/^FS1:$ACTIVE:/FS1:$NEW:/" api/internal/bnetbuild/testdata/era-kiloz.fs1 && rm -f api/internal/bnetbuild/testdata/era-kiloz.fs1.bak
grep -q "^FS1:$NEW:" api/internal/bnetbuild/testdata/era-kiloz.fs1 || die "api/internal/bnetbuild/testdata/era-kiloz.fs1 does not start FS1:$ACTIVE:, so it was not re-stamped"

# 5. The addon: Interface, version, changelog.
step "the addon's TOCs, version and changelog" "$FIX_THEN_REDO"
for toc in addon/ForeverSixty/ForeverSixty.toc addon/ForeverSixtyData/ForeverSixtyData.toc addon/ForeverSixtyInbox/ForeverSixtyInbox.toc; do
  sed -i.bak "s/^## Interface: .*/## Interface: $NEW_IFACE/" "$toc" && rm -f "$toc.bak"
done
OLD_VER="$(sed -n 's/^## Version: *//p' addon/ForeverSixty/ForeverSixty.toc)"
IFS=. read -r VMAJ VMIN VPAT <<<"$OLD_VER"
NEW_VER="$VMAJ.$VMIN.$((10#$VPAT + 1))"
TAG="addon-v$NEW_VER"
sed -i.bak "s/^## Version: .*/## Version: $NEW_VER/" addon/ForeverSixty/ForeverSixty.toc && rm -f addon/ForeverSixty/ForeverSixty.toc.bak
iface_note=""
[ "$NEW_IFACE" = "$OLD_IFACE" ] || iface_note=" and Interface $NEW_IFACE"
run python3 - "$NEW_VER" "$NEW" "$iface_note" <<'PY'
import sys
ver, build, iface = sys.argv[1:4]
path = "addon/ForeverSixty/CHANGELOG.md"
text = open(path, encoding="utf-8").read()
head, sep, rest = text.partition("\n## ")
entry = (f"\n## {ver}\n\nData for client build {build}{iface}: the item, talent and spell tables the addon "
         "ships were regenerated for the new client by the client watcher. Nothing else changed.\n")
open(path, "w", encoding="utf-8").write(head + entry + sep + rest)
PY

# 6. The gates data.yml runs, before anything is committed.
step "data.yml's --check gates for $NEW" "$FIX_THEN_REDO"
run in_dir data uv run python -m pipeline weights --build "$NEW" --check
run in_dir data uv run python -m pipeline addon-data --build "$NEW" --check
run in_dir data uv run python -m pipeline coach-rules --build "$NEW" --check

# 7. The test gate, before anything is committed, pushed or tagged.
if [ "$GATE" = on ]; then
  step "the test gate: web vitest and the API's active-build tests (tools/activation-gate.sh)" \
    "make each failing test read the build's data instead of pinning a name the client renamed (as web/src/pages/guides/_spec.test.ts's hit line does), or fix what it caught; land that on main, then $REDO"
  GATE_FAILURES="$(mktemp)"
  if ! tools/activation-gate.sh "$GATE_FAILURES"; then
    die "$(printf 'the test gate is red for %s; nothing was committed, pushed or tagged. Failing:\n%s' "$NEW" "$(cat "$GATE_FAILURES")")"
  fi
fi

git add "${MOVED_PATHS[@]}"
git status --short | head -40
[ "$MODE" != inspect ] || { UNDO=""; echo "--no-commit: stopping with the tree changed and staged"; exit 0; }

step "committing the move" "$REDO"
MSG="$(mktemp)"
cat > "$MSG" <<EOF
chore(client): activate build $NEW

The client watcher (client.yml) found $NEW on wago.tools, newer than the
active $ACTIVE, and data.yml's fetch landed its tables. This commit moves
the site, the addon and the guides to it: active-build.json, the build's
weights copy, every guide's FS1 code carried onto the new trees by talent
name and its data links, the addon's default builds from those guides,
addon-data.json and Data.lua, the addon TOCs' Interface ($OLD_IFACE ->
$NEW_IFACE), the addon version ($OLD_VER -> $NEW_VER) and the API's import
fixture stamp. Carried forward from $ACTIVE until the next ranking
replaces them: ${carried[*]:-nothing}.
EOF
run git -c user.name=forever-client-bot -c user.email=bot@foreversixty.gg commit --quiet --file="$MSG"
rm -f "$MSG"
UNDO="git reset --hard HEAD~1 (the move is one local commit; nothing was pushed)"
echo "committed $(git rev-parse --short HEAD)"
[ "$MODE" != commit ] || { echo "--no-publish: stopping after the commit"; exit 0; }

# --autostash: in CI the gate's Go run needs .github/actions/pin-engine, which leaves
# sim/go.mod and sim/go.sum dirty, and a plain rebase refuses a dirty tree.
step "pushing the move to main" "nothing was pushed or tagged; $REDO"
pushed=false
for attempt in 1 2 3 4 5; do
  if git pull --rebase --autostash --quiet origin main && git push --quiet origin HEAD:main; then pushed=true; break; fi
  [ "$attempt" -lt 5 ] || break
  echo "push attempt $attempt failed; retrying"; sleep 20
done
[ "$pushed" = true ] || die "could not rebase onto and push to origin/main after 5 attempts"
UNDO=""
MOVE_SHA="$(git rev-parse HEAD)"
echo "pushed to main"

# Main has moved from here on: a redo of the client run would say "already the active
# build", so every hand command below names the one step left.
RELEASE="gh workflow run addon-release.yml --ref $TAG -f dry_run=false -f publish=true"
step "tagging the addon release $TAG" \
  "git fetch origin && git tag -a $TAG $MOVE_SHA -m 'Forever Sixty $NEW_VER: data for client build $NEW' && git push origin $TAG, then $RELEASE and the dispatches the log did not reach (web.yml, api.yml, sim.yml, engine.yml with -f force=true)"
run git tag -a "$TAG" -m "Forever Sixty $NEW_VER: data for client build $NEW"
run git push --quiet origin "$TAG"
echo "tagged $TAG (addon-release.yml publishes it to CurseForge and Wago)"

# The ranking (bis.yml) and the dps-per-point re-measure it makes run only on an engine
# stamped for the new build (tools/engine/pinned-build.sh): an engine on the old trees
# reads the new client's talent strings by the old indexes, so its picks would be
# quietly wrong. Until engine.yml pins an engine that follows $NEW, the carried picks
# and figures stay live, and that pin's green push dispatches bis.yml with force.
ENGINE_PIN="$(sed -n 's/.*Version = "\(.*\)"/\1/p' sim/enginever/version.go)"
ENGINE_BUILD="$(tools/engine/pinned-build.sh . 2>/dev/null)" || ENGINE_BUILD="an unreadable build"
RANK_NOW=false
if [ "$ENGINE_BUILD" = "$NEW" ]; then
  RANK_NOW=true
else
  WAIT_LINE="ranking waits for engine follow: the pinned engine $ENGINE_PIN is stamped for $ENGINE_BUILD, not $NEW. The picks and dps-per-point carried from $ACTIVE stay live; engine.yml's green pin of an engine that follows $NEW dispatches bis.yml with force, which ranks and re-measures."
  echo "$WAIT_LINE"
  [ -z "${GITHUB_STEP_SUMMARY:-}" ] || printf '%s\n\n' "- $WAIT_LINE" >> "$GITHUB_STEP_SUMMARY"
fi

# A push made with GITHUB_TOKEN triggers no workflow (the 2026-10-16 move landed
# with no deploy, no addon publish and no ranking until each was dispatched by
# hand), so every consumer of the new build is dispatched here: the site, the
# API image (its simdb), the sim artifacts, the addon release for the tag, the
# engine follow, and (on an engine already stamped for it) the ranking. A failed
# dispatch does not stop the others; the run then fails naming each one to re-run.
if command -v gh >/dev/null 2>&1; then
  step "dispatching the consumers of $NEW" "run each command the reason lists"
  missed=()
  dispatch() { if gh workflow run "$@"; then echo "dispatched $1"; else missed+=("gh workflow run $*"); fi; }
  dispatch web.yml --ref main
  dispatch api.yml --ref main
  dispatch sim.yml --ref main
  dispatch addon-release.yml --ref "$TAG" -f dry_run=false -f publish=true
  [ "$RANK_NOW" != true ] || dispatch bis.yml --ref main -f force=true
  dispatch engine.yml --ref main -f force=true
  [ ${#missed[@]} -eq 0 ] || die "$(printf 'the move is on main and tagged, but these dispatches failed:\n%s' "$(printf '%s\n' "${missed[@]}")")"
fi
