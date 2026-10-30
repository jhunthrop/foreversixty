#!/usr/bin/env bash
# tools/activate-build.sh <build> [--no-commit | --no-publish]
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
#      none, so the site has picks the same morning; the nightly ranking (bis.yml,
#      dispatched with force at the end) replaces bis/ with the new build's own within
#      hours. addon-builds.json is NOT carried: it is rebuilt from the moved guides.
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
#
# --no-commit leaves the working tree changed and stops (inspect with git diff).
# --no-publish commits and stops (no push, no tag, no ranking dispatch).
set -euo pipefail

usage() { echo "usage: $0 <build> [--no-commit | --no-publish]" >&2; exit 2; }
[ $# -ge 1 ] || usage
NEW="$1"; shift
MODE=publish
for arg in "$@"; do
  case "$arg" in
    --no-commit) MODE=inspect ;;
    --no-publish) MODE=commit ;;
    *) usage ;;
  esac
done
[[ "$NEW" =~ ^[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+$ ]] || { echo "$NEW is not a four-part client build" >&2; exit 2; }

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"
ACTIVE_JSON=web/src/data/active-build.json
ACTIVE="$(sed -n 's/.*"build"[[:space:]]*:[[:space:]]*"\([^"]*\)".*/\1/p' "$ACTIVE_JSON")"
[ -n "$ACTIVE" ] || { echo "$ACTIVE_JSON names no build" >&2; exit 1; }
if [ "$NEW" = "$ACTIVE" ]; then echo "$NEW is already the active build; nothing to do"; exit 0; fi

NEW_DIR="data/builds/$NEW"; OLD_DIR="data/builds/$ACTIVE"
for f in items.json simdb.bin talents.json levels.json spellranks.json loot.json manifest.json; do
  [ -e "$NEW_DIR/$f" ] || { echo "$NEW_DIR/$f is missing: run data.yml's fetch for $NEW first" >&2; exit 1; }
done

# The Interface number Blizzard's TOC convention gives this version (1.60.1 -> 16001).
IFS=. read -r MAJOR MINOR PATCH _ <<<"$NEW"
NEW_IFACE=$((10#$MAJOR * 10000 + 10#$MINOR * 100 + 10#$PATCH))
OLD_IFACE="$(sed -n 's/^## Interface: *\([0-9]*\).*/\1/p' addon/ForeverSixty/ForeverSixty.toc)"

echo "activating $NEW (was $ACTIVE); Interface $OLD_IFACE -> $NEW_IFACE"

# 2. Carry the ranking's outputs forward where the new build has none.
carried=()
if [ ! -d "$NEW_DIR/bis" ] && [ -d "$OLD_DIR/bis" ]; then
  # The carried files now stand for the new build (web/src/lib/bis reads them from its
  # directory and expects their build field to agree); generated_at and engine_version
  # keep saying which ranking they came from. A pick or alternative the new client's
  # items.json lacks is dropped with a report line (pipeline/bis_carry.py).
  ( cd data && uv run python -m pipeline carry-bis --from "$ACTIVE" --to "$NEW" ); carried+=("bis/")
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
node web/scripts/reencode-guide-builds.mjs
grep -rl "foreversixty.gg/data/$ACTIVE/" web/src/content/guides | while read -r f; do
  sed -i.bak "s#foreversixty.gg/data/$ACTIVE/#foreversixty.gg/data/$NEW/#g" "$f" && rm -f "$f.bak"
done

# 4. The addon's default builds, from the guides just moved, then the build's
#    weights copy, addon-data.json and the addon's Data.lua.
go run ./sim/cmd/addon-builds -repo-root .
( cd data && uv run python -m pipeline weights --build "$NEW" && uv run python -m pipeline addon-data --build "$NEW" )

# 4a. The log coach's rule files, restamped against the new build's tables.
( cd data && uv run python -m pipeline coach-rules --build "$NEW" )

# 4b. The API's Battle.net import fixture is stamped with the active build.
sed -i.bak "s/^FS1:$ACTIVE:/FS1:$NEW:/" api/internal/bnetbuild/testdata/era-kiloz.fs1 && rm -f api/internal/bnetbuild/testdata/era-kiloz.fs1.bak
grep -q "^FS1:$NEW:" api/internal/bnetbuild/testdata/era-kiloz.fs1 || { echo "the API fixture was not re-stamped"; exit 1; }

# 5. The addon: Interface, version, changelog.
for toc in addon/ForeverSixty/ForeverSixty.toc addon/ForeverSixtyData/ForeverSixtyData.toc addon/ForeverSixtyInbox/ForeverSixtyInbox.toc; do
  sed -i.bak "s/^## Interface: .*/## Interface: $NEW_IFACE/" "$toc" && rm -f "$toc.bak"
done
OLD_VER="$(sed -n 's/^## Version: *//p' addon/ForeverSixty/ForeverSixty.toc)"
IFS=. read -r VMAJ VMIN VPAT <<<"$OLD_VER"
NEW_VER="$VMAJ.$VMIN.$((10#$VPAT + 1))"
sed -i.bak "s/^## Version: .*/## Version: $NEW_VER/" addon/ForeverSixty/ForeverSixty.toc && rm -f addon/ForeverSixty/ForeverSixty.toc.bak
iface_note=""
[ "$NEW_IFACE" = "$OLD_IFACE" ] || iface_note=" and Interface $NEW_IFACE"
python3 - "$NEW_VER" "$NEW" "$iface_note" <<'PY'
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
( cd data && uv run python -m pipeline weights --build "$NEW" --check && uv run python -m pipeline addon-data --build "$NEW" --check && uv run python -m pipeline coach-rules --build "$NEW" --check )

git add "$ACTIVE_JSON" "$NEW_DIR" addon/ForeverSixty addon/ForeverSixtyData addon/ForeverSixtyInbox data/curated/coach web/src/content/guides api/internal/bnetbuild/testdata/era-kiloz.fs1
git status --short | head -40
[ "$MODE" != inspect ] || { echo "--no-commit: stopping with the tree changed and staged"; exit 0; }

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
git -c user.name=forever-client-bot -c user.email=bot@foreversixty.gg commit --quiet --file="$MSG"
rm -f "$MSG"
echo "committed $(git rev-parse --short HEAD)"
[ "$MODE" != commit ] || { echo "--no-publish: stopping after the commit"; exit 0; }

for attempt in 1 2 3 4 5; do
  if git pull --rebase --quiet origin main && git push --quiet origin HEAD:main; then break; fi
  [ "$attempt" -lt 5 ] || { echo "could not push the activation after 5 attempts" >&2; exit 1; }
  echo "push attempt $attempt failed; retrying"; sleep 20
done
echo "pushed to main"

TAG="addon-v$NEW_VER"
git tag -a "$TAG" -m "Forever Sixty $NEW_VER: data for client build $NEW" && git push --quiet origin "$TAG"
echo "tagged $TAG (addon-release.yml publishes it to CurseForge and Wago)"

# A push made with GITHUB_TOKEN triggers no workflow (the 2026-10-16 move landed
# with no deploy, no addon publish and no ranking until each was dispatched by
# hand), so every consumer of the new build is dispatched here: the site, the
# API image (its simdb), the sim artifacts, the addon release for the tag, the
# engine follow, and the ranking that replaces the carried-forward picks.
if command -v gh >/dev/null 2>&1; then
  gh workflow run web.yml --ref main && echo "dispatched the site deploy"
  gh workflow run api.yml --ref main && echo "dispatched the API build and deploy"
  gh workflow run sim.yml --ref main && echo "dispatched the sim artifacts"
  gh workflow run addon-release.yml --ref "$TAG" -f dry_run=false -f publish=true && echo "dispatched the addon release for $TAG"
  gh workflow run bis.yml --ref main -f force=true && echo "dispatched the ranking for $NEW"
  gh workflow run engine.yml --ref main -f force=true && echo "dispatched the engine follow for $NEW"
fi
