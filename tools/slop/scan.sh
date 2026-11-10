#!/usr/bin/env bash
# Greps every player-facing copy source (tools/slop/sources.txt) for the AI-slop patterns in
# tools/slop/patterns.txt, minus comments and the allowlist.
# Usage: tools/slop/scan.sh [--md | --counts] [--check] [--skip <path prefix>]... | --probe <text>
set -euo pipefail
exec node "$(dirname "$0")/scan.mjs" "$@"
