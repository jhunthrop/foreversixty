#!/usr/bin/env bash
# tools/publish-feed.sh <family> <tag> <file>...
#
# Publishes one signed release to the site's own release feed (web/src/lib/downloads): each
# file goes to the bucket at downloads/<family>/<tag>/<name>, served immutable at
# foreversixty.gg/dl/<family>/<tag>/<name>, and then downloads/<family>/latest.json is
# rewritten to name them. The manifest goes last, so it never names a file that is not
# there yet. The addon-release, addon-data-release and companion workflows run it beside
# their GitHub release; docs/runbooks/repo-private.md says why the feed exists.
#
#   tools/publish-feed.sh addon addon-v2.1.0 ForeverSixty-2.1.0.zip ForeverSixty-2.1.0.zip.minisig
#
# Every file but a .minisig or .txt must have its <file>.minisig beside it in the list: the
# companion installs nothing unsigned, so the feed never names an unsigned file.
#
# Environment (defaults): CLOUDFLARE_API_TOKEN and CLOUDFLARE_ACCOUNT_ID (wrangler's own),
# FEED_BUCKET (foreversixty-logs), FEED_ORIGIN (https://foreversixty.gg), FEED_PUBLISHED_AT
# (now, UTC), WRANGLER (npx --yes wrangler@4.131.1, the web/ lockfile's version).
#
# Exit status: 0 published; 1 an upload failed; 64 usage or an unsigned or misnamed file.
set -euo pipefail

readonly BUCKET=${FEED_BUCKET:-foreversixty-logs}
readonly ORIGIN=${FEED_ORIGIN:-https://foreversixty.gg}
read -r -a WRANGLER <<<"${WRANGLER:-npx --yes wrangler@4.131.1}"
readonly WRANGLER

usage() {
	echo "usage: tools/publish-feed.sh <companion|addon|data> <tag> <file>..." >&2
	exit 64
}

die() {
	echo "publish-feed: $*" >&2
	exit 64
}

(($# >= 3)) || usage
readonly FAMILY=$1 TAG=$2
shift 2

case $FAMILY in companion | addon | data) ;; *) die "unknown family '$FAMILY'" ;; esac
# The Worker serves only tags of the family's own prefix (web/src/lib/downloads/paths.ts).
[[ $TAG =~ ^${FAMILY}-v[0-9][0-9A-Za-z.-]{0,40}$ && $TAG != *..* ]] || die "'$TAG' is not a $FAMILY-v tag"

names=()
for file in "$@"; do
	[[ -s $file ]] || die "$file is missing or empty"
	name=$(basename "$file")
	[[ $name =~ ^[A-Za-z0-9][A-Za-z0-9._-]{0,127}$ && $name != *..* ]] || die "'$name' is not a name the feed serves"
	names+=("$name")
done
for name in "${names[@]}"; do
	case $name in *.minisig | *.txt) continue ;; esac
	printf '%s\n' "${names[@]}" | grep -qxF "$name.minisig" || die "$name has no $name.minisig beside it"
done

content_type() {
	case $1 in
	*.exe) echo application/vnd.microsoft.portable-executable ;;
	*.dmg) echo application/x-apple-diskimage ;;
	*.zip) echo application/zip ;;
	*.json) echo 'application/json; charset=utf-8' ;;
	*.minisig | *.txt) echo 'text/plain; charset=utf-8' ;;
	*) echo application/octet-stream ;;
	esac
}

put() { # put <key> <file>
	"${WRANGLER[@]}" r2 object put "$BUCKET/$1" --file "$2" --content-type "$(content_type "$1")" --remote
}

assets='[]'
for file in "$@"; do
	name=$(basename "$file")
	put "downloads/$FAMILY/$TAG/$name" "$file"
	assets=$(jq -c \
		--arg name "$name" \
		--arg url "$ORIGIN/dl/$FAMILY/$TAG/$name" \
		--argjson size "$(wc -c <"$file" | tr -d ' ')" \
		--arg sha "$(sha256sum "$file" | cut -d' ' -f1)" \
		'. + [{name: $name, browser_download_url: $url, size: $size, sha256: $sha}]' <<<"$assets")
done

manifest=$(mktemp)
trap 'rm -f "$manifest"' EXIT
jq -n \
	--arg family "$FAMILY" \
	--arg tag "$TAG" \
	--arg at "${FEED_PUBLISHED_AT:-$(date -u +%Y-%m-%dT%H:%M:%SZ)}" \
	--argjson assets "$assets" \
	'{schema: 1, family: $family, tag_name: $tag, published_at: $at, assets: $assets}' >"$manifest"
put "downloads/$FAMILY/latest.json" "$manifest"
echo "publish-feed: $FAMILY $TAG is live at $ORIGIN/dl/$FAMILY/latest.json"
