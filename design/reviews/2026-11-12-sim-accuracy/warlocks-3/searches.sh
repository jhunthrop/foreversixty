#!/usr/bin/env bash
# searches.sh: the warlocks pass 3 talent and rotation searches, one at a
# time, niced to 15, GOMAXPROCS=5 (other lanes share the box). Raid 60 on
# the raid builds and bare 60 on the leveling builds for Affliction and
# Destruction; the keep lists hold what each card casts or leans on
# (Suppression for hit and threat). Run from the site's sim directory
# (tools/remote-test.sh --dir sim --engine-sha <sha>); prints every report.
set -euo pipefail
make -C .. simdb >/dev/null
out=$(mktemp -d)
bin=$(mktemp -d)
mkdir -p "$out/talents" "$out/rotation"
go build -o "$bin/talent-search" ./cmd/talent-search
go build -o "$bin/rotation-search" ./cmd/rotation-search
run() { echo "## $*"; nice --adjustment=15 env GOMAXPROCS=5 "$@" 2>&1 | grep -v '^$' | tail -2; }
aff_raid_keep='Wrack,Siphon Life,Amplify Curse,Nightfall,Demonic Sacrifice,Suppression'
aff_bare_keep='Wrack,Siphon Life,Amplify Curse,Suppression'
dest_keep='Conflagrate,Shadowburn,Incinerate,Bane of Havoc,Suppression'
run "$bin/rotation-search" -repo-root .. -out "$out/rotation" -spec warlock-affliction -preset raid
run "$bin/rotation-search" -repo-root .. -out "$out/rotation" -spec warlock-destruction -preset raid
run "$bin/talent-search" -repo-root .. -out "$out/talents" -spec warlock-affliction -preset raid -keep "$aff_raid_keep"
run "$bin/talent-search" -repo-root .. -out "$out/talents" -spec warlock-destruction -preset raid -keep "$dest_keep"
run "$bin/rotation-search" -repo-root .. -out "$out/rotation" -spec warlock-destruction
run "$bin/rotation-search" -repo-root .. -out "$out/rotation" -spec warlock-affliction
run "$bin/talent-search" -repo-root .. -out "$out/talents" -spec warlock-destruction -keep "$dest_keep"
run "$bin/talent-search" -repo-root .. -out "$out/talents" -spec warlock-affliction -keep "$aff_bare_keep"
for f in "$out"/*/*.md; do echo "=== $(basename "$(dirname "$f")")/$(basename "$f")"; cat "$f"; done
