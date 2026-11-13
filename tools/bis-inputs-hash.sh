#!/usr/bin/env bash
# bis-inputs-hash.sh: the hash of everything bis.yml's ranking reads, at HEAD.
#
#   tools/bis-inputs-hash.sh <build>
#
# Run from the repository root. Prints 16 hex characters: the ranker and the engine pin
# (sim/), the curated specs and rotations (data/curated/), the active build, the Makefile,
# bis.yml, the guides, and the build's own data directory minus the ranking's OUTPUTS
# (bis/, addon-data.json, dps-per-point.*, manifest.json). Each landed spec carries it as
# data/builds/<build>/bis/<spec>.inputs, which is how a ranking skips the specs whose
# picks on the branch were ranked on the same inputs (the resume after a lost leg) and
# how a night with nothing new skips the whole ranking.
set -euo pipefail
[ $# -eq 1 ] || { echo "usage: $0 <build>" >&2; exit 2; }
{
  git rev-parse HEAD:sim HEAD:data/curated HEAD:web/src/data/active-build.json HEAD:Makefile HEAD:.github/workflows/bis.yml HEAD:web/src/content/guides
  git ls-tree -r HEAD "data/builds/$1" | grep -vE "/bis/|addon-data\.json|dps-per-point\.|/manifest\.json$"
} | sha256sum | cut -c1-16
