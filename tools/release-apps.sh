#!/usr/bin/env bash
# tools/release-apps.sh [--bump build|patch|minor|major] [--platform both|android|ios] [--execute]
#
# Cuts one release of both native apps from apps/VERSION and apps/CHANGELOG.md
# (docs/runbooks/mobile-apps.md "Cutting a release"). `make -C apps release` runs it, and so
# does .github/workflows/apps-release.yml.
#
# A dry run is the default: it checks everything, prints the version it would cut, the notes
# and every command, and changes nothing. With --execute it:
#   1. bumps apps/VERSION (APP_BUILD always rises by one; --bump patch|minor|major also moves
#      APP_VERSION) and turns the changelog's Unreleased section into the version's section,
#   2. commits that as "chore(apps): release <version> (<build>)" and tags apps-v<version>-<build>,
#   3. Android: builds the signed release bundle and uploads it, with its R8 mapping and the
#      notes, to the Play track (PLAY_TRACK, default internal) through tools/play-upload.ts
#      (the Play Developer API with a short-lived token; no service-account key exists),
#   4. iOS: archives, exports and uploads to TestFlight with fastlane's upload_to_testflight,
#      once the Apple account exists. Until then iOS is skipped with what is missing named
#      (or fails, when --platform ios asked for it).
# It never pushes: push the commit and the tag afterwards (git push --follow-tags), as the
# workflow does. A build number spent on a failed upload is simply skipped next time.
#
# Environment (secrets come from the environment, never from a file in the repo):
#   Android signing  ANDROID_KEYSTORE_PATH (the upload keystore's path), ANDROID_KEYSTORE_PASSWORD,
#                    ANDROID_KEY_ALIAS, ANDROID_KEY_PASSWORD (docs/runbooks/mobile-apps.md,
#                    "Android release infrastructure": GSM holds them, GitHub has copies)
#   Play upload      PLAY_ACCESS_TOKEN (CI: minted by GitHub federation as PLAY_SA), or
#                    PLAY_SA (a local release: the token is minted with gcloud impersonating it)
#                    PLAY_TRACK (internal), PLAY_RELEASE_STATUS (draft until the app's first
#                    review; then completed)
#   iOS              APPLE_TEAM_ID, ASC_KEY_ID, ASC_ISSUER_ID, ASC_KEY_P8 (the .p8 key's text)
#   Crash reports    SENTRY_DSN_ANDROID, SENTRY_DSN_IOS (optional: without one that build
#                    reports nothing)
#   Tests            RELEASE_ROOT (repo root), GRADLEW, XCODEBUILD, XCODEGEN, FASTLANE, NODE,
#                    GCLOUD
set -euo pipefail

ROOT=${RELEASE_ROOT:-$(git rev-parse --show-toplevel)}
readonly ROOT
readonly VERSION_FILE=$ROOT/apps/VERSION
readonly CHANGELOG=$ROOT/apps/CHANGELOG.md
readonly ANDROID_DIR=$ROOT/apps/android
readonly IOS_DIR=$ROOT/apps/ios
readonly PACKAGE=gg.foreversixty.app
readonly GRADLEW=${GRADLEW:-$ANDROID_DIR/gradlew}
readonly XCODEBUILD=${XCODEBUILD:-xcodebuild}
readonly XCODEGEN=${XCODEGEN:-xcodegen}
readonly NODE=${NODE:-node}
readonly GCLOUD=${GCLOUD:-gcloud}
readonly PLAY_UPLOAD=$ROOT/tools/play-upload.ts
readonly PLAY_SCOPE=https://www.googleapis.com/auth/androidpublisher
readonly PLAY_TOKEN_HELP="run gcloud auth login; your account needs Service Account Token Creator on PLAY_SA"
# shellcheck disable=SC2206 # FASTLANE is a command line, split on purpose.
readonly FASTLANE_CMD=(${FASTLANE:-bundle exec fastlane})
# fastlane comes from apps/Gemfile (pinned), whatever directory the script runs from.
export BUNDLE_GEMFILE=${BUNDLE_GEMFILE:-$ROOT/apps/Gemfile}
readonly PLAY_TRACK=${PLAY_TRACK:-internal}
readonly PLAY_RELEASE_STATUS=${PLAY_RELEASE_STATUS:-draft}
readonly ANDROID_SIGNING=(ANDROID_KEYSTORE_PATH ANDROID_KEYSTORE_PASSWORD ANDROID_KEY_ALIAS ANDROID_KEY_PASSWORD)
readonly APPLE_ACCOUNT=(APPLE_TEAM_ID ASC_KEY_ID ASC_ISSUER_ID ASC_KEY_P8)

BUMP=build
PLATFORM=both
EXECUTE=false

die() {
	echo "release-apps: $*" >&2
	exit 1
}

usage() {
	sed -n '2,3p' "$0" | sed 's/^# //'
	exit 2
}

parse_args() {
	while [ $# -gt 0 ]; do
		case $1 in
		--bump) BUMP=${2:-}; shift 2 ;;
		--platform) PLATFORM=${2:-}; shift 2 ;;
		--execute) EXECUTE=true; shift ;;
		-h | --help) usage ;;
		*) die "unknown argument: $1" ;;
		esac
	done
	case $BUMP in build | patch | minor | major) ;; *) die "--bump must be build, patch, minor or major" ;; esac
	case $PLATFORM in both | android | ios) ;; *) die "--platform must be both, android or ios" ;; esac
}

# setting <key>: the value of one KEY = value line of apps/VERSION.
setting() {
	sed -n "s/^$1[[:space:]]*=[[:space:]]*//p" "$VERSION_FILE" | tr -d '[:space:]'
}

read_version() {
	[ -f "$VERSION_FILE" ] || die "apps/VERSION is missing"
	VERSION=$(setting APP_VERSION)
	BUILD=$(setting APP_BUILD)
	[[ $VERSION =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]] || die "apps/VERSION: APP_VERSION '$VERSION' is not x.y.z"
	[[ $BUILD =~ ^[1-9][0-9]*$ ]] || die "apps/VERSION: APP_BUILD '$BUILD' is not a positive integer"
}

# next_version: NEXT_VERSION and NEXT_BUILD from the bump.
next_version() {
	local major minor patch
	IFS=. read -r major minor patch <<<"$VERSION"
	case $BUMP in
	build) NEXT_VERSION=$VERSION ;;
	patch) NEXT_VERSION=$major.$minor.$((patch + 1)) ;;
	minor) NEXT_VERSION=$major.$((minor + 1)).0 ;;
	major) NEXT_VERSION=$((major + 1)).0.0 ;;
	esac
	NEXT_BUILD=$((BUILD + 1))
}

# unreleased_notes: the lines under "## Unreleased", up to the next "## ", without the
# blank lines around them.
unreleased_notes() {
	awk '
		/^## Unreleased[[:space:]]*$/ { inside = 1; next }
		/^## / { inside = 0 }
		inside { lines[++n] = $0 }
		END {
			first = 1; while (first <= n && lines[first] ~ /^[[:space:]]*$/) first++
			last = n; while (last >= first && lines[last] ~ /^[[:space:]]*$/) last--
			for (i = first; i <= last; i++) print lines[i]
		}' "$CHANGELOG"
}

# missing <name>...: the names among its arguments whose variables are empty.
missing() {
	local name out=()
	for name in "$@"; do
		[ -n "${!name:-}" ] || out+=("$name")
	done
	echo "${out[*]:-}"
}

preflight() {
	[ -f "$CHANGELOG" ] || die "apps/CHANGELOG.md is missing"
	grep -q '^## Unreleased[[:space:]]*$' "$CHANGELOG" || die "apps/CHANGELOG.md has no '## Unreleased' section"
	NOTES=$(unreleased_notes)
	if [ "$BUMP" != build ] && [ -z "$NOTES" ]; then
		die "a $BUMP release needs notes: write them under '## Unreleased' in apps/CHANGELOG.md"
	fi
	if [ ${#NOTES} -gt 500 ]; then
		die "the Unreleased notes are ${#NOTES} characters; Play's limit is 500"
	fi
	if $EXECUTE && [ -n "$(git -C "$ROOT" status --porcelain -- apps/VERSION apps/CHANGELOG.md)" ]; then
		die "apps/VERSION or apps/CHANGELOG.md has uncommitted changes"
	fi
	ANDROID_MISSING=$(missing "${ANDROID_SIGNING[@]}")
	if [ -z "${PLAY_ACCESS_TOKEN:-}" ] && [ -z "${PLAY_SA:-}" ]; then
		ANDROID_MISSING="${ANDROID_MISSING:+$ANDROID_MISSING }PLAY_SA (or PLAY_ACCESS_TOKEN)"
	fi
	APPLE_MISSING=$(missing "${APPLE_ACCOUNT[@]}")
	DO_ANDROID=false
	DO_IOS=false
	case $PLATFORM in android | both) DO_ANDROID=true ;; esac
	case $PLATFORM in ios | both) DO_IOS=true ;; esac
	if $DO_IOS && [ -n "$APPLE_MISSING" ]; then
		IOS_SKIP="Apple enrollment is not finished: set $APPLE_MISSING (docs/runbooks/mobile-apps.md, 'TestFlight')"
		[ "$PLATFORM" = ios ] && $EXECUTE && die "iOS cannot be released: $IOS_SKIP"
		DO_IOS=false
	fi
	if $DO_ANDROID && $EXECUTE && [ -n "$ANDROID_MISSING" ]; then
		die "Android cannot be released: set $ANDROID_MISSING"
	fi
}

# bump_files: the new VERSION, and the changelog's Unreleased notes under the new heading.
bump_files() {
	local today heading tmp
	today=$(date -u +%Y-%m-%d)
	heading="## $NEXT_VERSION ($NEXT_BUILD), $today"
	sed -i.bak -e "s/^APP_VERSION[[:space:]]*=.*/APP_VERSION = $NEXT_VERSION/" \
		-e "s/^APP_BUILD[[:space:]]*=.*/APP_BUILD = $NEXT_BUILD/" "$VERSION_FILE"
	rm -f "$VERSION_FILE.bak"
	tmp=$(mktemp)
	awk -v heading="$heading" '/^## Unreleased[[:space:]]*$/ { print; print ""; print heading; next } { print }' \
		"$CHANGELOG" >"$tmp"
	mv "$tmp" "$CHANGELOG"
}

commit_and_tag() {
	git -C "$ROOT" add apps/VERSION apps/CHANGELOG.md
	git -C "$ROOT" commit -q -m "chore(apps): release $NEXT_VERSION ($NEXT_BUILD)"
	# Annotated, so `git push --follow-tags` carries it.
	git -C "$ROOT" tag -a -m "Forever Sixty apps $NEXT_VERSION ($NEXT_BUILD)" "apps-v$NEXT_VERSION-$NEXT_BUILD"
	echo "release-apps: committed and tagged apps-v$NEXT_VERSION-$NEXT_BUILD (push with: git push --follow-tags)"
}

# run <command...>: runs it, or in a dry run only prints it.
run() {
	if $EXECUTE; then
		"$@"
	else
		printf '  would run:'
		printf ' %q' "$@"
		printf '\n'
	fi
}

release_android() {
	local work=$1 aab=$ANDROID_DIR/app/build/outputs/bundle/release/app-release.aab
	local mapping=$ANDROID_DIR/app/build/outputs/mapping/release/mapping.txt
	local notes=$work/play-notes.txt
	echo "Android: signed bundle, then Play track '$PLAY_TRACK' as $PLAY_RELEASE_STATUS"
	[ -n "${SENTRY_DSN_ANDROID:-}" ] || echo "  note: SENTRY_DSN_ANDROID is empty, so this build reports no crashes"
	printf '%s\n' "${NOTES:-Bug fixes and improvements.}" >"$notes"
	run "$GRADLEW" -p "$ANDROID_DIR" bundleRelease
	local token=''
	if $EXECUTE; then token=$(play_token) || die "no Play token: $PLAY_TOKEN_HELP"; fi
	# The token reaches the upload through its environment, never a command line.
	PLAY_ACCESS_TOKEN=$token run "$NODE" "$PLAY_UPLOAD" --package "$PACKAGE" --track "$PLAY_TRACK" \
		--status "$PLAY_RELEASE_STATUS" --bundle "$aab" --mapping "$mapping" --notes "$notes"
}

# play_token: the caller's short-lived token, or one minted by impersonating PLAY_SA.
play_token() {
	if [ -n "${PLAY_ACCESS_TOKEN:-}" ]; then
		printf '%s' "$PLAY_ACCESS_TOKEN"
	else
		"$GCLOUD" auth print-access-token --impersonate-service-account="$PLAY_SA" --scopes="$PLAY_SCOPE"
	fi
}

release_ios() {
	local work=$1 archive=$1/ForeverSixty.xcarchive export=$1/export key=$1/asc-key.json
	echo "iOS: archive, export and TestFlight"
	[ -n "${SENTRY_DSN_IOS:-}" ] || echo "  note: SENTRY_DSN_IOS is empty, so this build reports no crashes"
	write_asc_key "$key"
	run "$XCODEGEN" generate --spec "$IOS_DIR/project.yml"
	run "$XCODEBUILD" archive -project "$IOS_DIR/ForeverSixty.xcodeproj" -scheme ForeverSixty \
		-configuration Release -destination 'generic/platform=iOS' -archivePath "$archive" \
		-allowProvisioningUpdates -authenticationKeyPath "$work/asc-key.p8" \
		-authenticationKeyID "${ASC_KEY_ID:-}" -authenticationKeyIssuerID "${ASC_ISSUER_ID:-}" \
		DEVELOPMENT_TEAM="${APPLE_TEAM_ID:-}" CODE_SIGN_STYLE=Automatic CODE_SIGN_IDENTITY="Apple Distribution" \
		SENTRY_DSN="${SENTRY_DSN_IOS:-}"
	run "$XCODEBUILD" -exportArchive -archivePath "$archive" -exportPath "$export" \
		-exportOptionsPlist "$IOS_DIR/ExportOptions.plist" -allowProvisioningUpdates \
		-authenticationKeyPath "$work/asc-key.p8" -authenticationKeyID "${ASC_KEY_ID:-}" \
		-authenticationKeyIssuerID "${ASC_ISSUER_ID:-}"
	run "${FASTLANE_CMD[@]}" run upload_to_testflight ipa:"$export/ForeverSixty.ipa" \
		api_key_path:"$key" changelog:"${NOTES:-Bug fixes and improvements.}" \
		skip_waiting_for_build_processing:true
}

# write_asc_key <json path>: the App Store Connect key, as the .p8 xcodebuild reads and the
# JSON fastlane reads, inside the run's private scratch directory.
write_asc_key() {
	$EXECUTE || return 0
	printf '%s\n' "$ASC_KEY_P8" >"$(dirname "$1")/asc-key.p8"
	ruby -rjson -e 'puts JSON.generate(key_id: ENV.fetch("ASC_KEY_ID"), issuer_id: ENV.fetch("ASC_ISSUER_ID"), key: ENV.fetch("ASC_KEY_P8"))' >"$1"
}

print_plan() {
	local mode=dry-run
	$EXECUTE && mode=execute
	echo "release-apps ($mode): $VERSION ($BUILD) -> $NEXT_VERSION ($NEXT_BUILD), bump $BUMP, platform $PLATFORM"
	echo "Notes:"
	if [ -n "$NOTES" ]; then printf '%s\n' "$NOTES" | sed 's/^/  /'; else echo "  (none)"; fi
	if $DO_ANDROID && [ -n "$ANDROID_MISSING" ]; then echo "Android: blocked until these are set: $ANDROID_MISSING"; fi
	if [ -n "${IOS_SKIP:-}" ]; then echo "iOS: skipped. $IOS_SKIP"; fi
}

main() {
	parse_args "$@"
	read_version
	next_version
	preflight
	print_plan
	$DO_ANDROID || $DO_IOS || die "nothing to release on platform '$PLATFORM'"
	local work
	work=$(mktemp -d)
	chmod 700 "$work"
	trap 'rm -rf "'"$work"'"' EXIT
	if $EXECUTE; then
		bump_files
		commit_and_tag
	else
		echo "  would write apps/VERSION ($NEXT_VERSION, $NEXT_BUILD), roll apps/CHANGELOG.md, commit and tag apps-v$NEXT_VERSION-$NEXT_BUILD"
	fi
	if $DO_ANDROID; then release_android "$work"; fi
	if $DO_IOS; then release_ios "$work"; fi
	$EXECUTE || echo "Dry run: nothing was changed. Add --execute to release."
}

main "$@"
