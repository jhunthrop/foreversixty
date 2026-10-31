#!/usr/bin/env bash
# Tests for tools/release-apps.sh, the one release command of both native apps. Plain bash,
# the shape of test_activation.sh: each case runs the real script in a scratch repository
# holding only apps/VERSION and apps/CHANGELOG.md, with gradlew, xcodegen, xcodebuild,
# fastlane, node (the Play upload) and gcloud (the Play token) stubbed (GRADLEW, XCODEGEN,
# XCODEBUILD, FASTLANE, NODE, GCLOUD) to record their calls.
#
#   bash tools/tests/test_release_apps.sh
#
# Prints "ok N tests" and exits 0 when every assertion holds; names each failure and exits
# 1 otherwise.
set -uo pipefail

HERE=$(cd "$(dirname "$0")" && pwd)
readonly HERE
readonly RELEASE=$HERE/../release-apps.sh

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

# stub <dir> <name>: a command that appends its arguments, and the Play token it saw in its
# environment, to <dir>/calls; the Play upload also keeps the notes file it was given, and
# gcloud prints a minted token (or fails, when <dir>/gcloud-fails exists).
stub() {
	local dir=$1 name=$2
	mkdir -p "$dir"
	cat >"$dir/$name" <<EOF
#!/usr/bin/env bash
echo "$name \$* [token:\${PLAY_ACCESS_TOKEN:-none}]" >>"$dir/calls"
prev=
for arg in "\$@"; do
  [ "\$prev" = --notes ] && cat "\$arg" >"$dir/play-notes"
  prev=\$arg
done
if [ "$name" = gcloud ]; then
  [ -e "$dir/gcloud-fails" ] && exit 1
  printf minted-token
fi
exit 0
EOF
	chmod +x "$dir/$name"
}

# repo <name> <notes>: a scratch git repository with apps/VERSION at 1.4.2 (17) and a
# changelog whose Unreleased section holds <notes>; prints its path.
repo() {
	local dir=$SCRATCH/$1 notes=$2
	mkdir -p "$dir/apps"
	printf '%s\n' '// comment line' 'APP_VERSION = 1.4.2' 'APP_BUILD = 17' >"$dir/apps/VERSION"
	printf '%s\n' '# Notes' '' '## Unreleased' '' "$notes" '' '## 1.4.2 (16), 2026-10-01' '' '- Older.' >"$dir/apps/CHANGELOG.md"
	git -C "$dir" init -q
	git -C "$dir" add -A
	git -C "$dir" -c user.name=t -c user.email=t@t commit -q -m init
	echo "$dir"
}

# release <repo> <args...>: runs the script against <repo> with every tool stubbed in
# <repo>/stubs; stdout and stderr go to <repo>/out, the exit status is returned.
release() {
	local dir=$1
	shift
	local stubs=$dir/stubs
	for name in gradlew xcodegen xcodebuild fastlane node gcloud; do stub "$stubs" "$name"; done
	(
		cd "$dir" || exit 1
		RELEASE_ROOT=$dir GRADLEW=$stubs/gradlew XCODEGEN=$stubs/xcodegen XCODEBUILD=$stubs/xcodebuild \
			FASTLANE=$stubs/fastlane NODE=$stubs/node GCLOUD=$stubs/gcloud GIT_AUTHOR_NAME=t GIT_AUTHOR_EMAIL=t@t GIT_COMMITTER_NAME=t \
			GIT_COMMITTER_EMAIL=t@t "$RELEASE" "$@"
	) >"$dir/out" 2>&1
}

# The CI shape: the signing values and a federated token.
with_android_secrets() {
	ANDROID_KEYSTORE_PATH=/k.jks ANDROID_KEYSTORE_PASSWORD=p ANDROID_KEY_ALIAS=a \
		ANDROID_KEY_PASSWORD=p PLAY_ACCESS_TOKEN=ci-token "$@"
}

# The local shape: the signing values and the Play service account to impersonate.
with_android_impersonation() {
	ANDROID_KEYSTORE_PATH=/k.jks ANDROID_KEYSTORE_PASSWORD=p ANDROID_KEY_ALIAS=a \
		ANDROID_KEY_PASSWORD=p PLAY_SA=play-publisher@example.iam.gserviceaccount.com "$@"
}

calls() { cat "$1/stubs/calls" 2>/dev/null; }
version_file() { cat "$1/apps/VERSION"; }

# A dry run changes nothing, runs nothing and says what it would cut.
dir=$(repo dry '- A fix.')
release "$dir"
assert "dry run exits 0" equal $? 0
assert "dry run names the next build" contains "$(cat "$dir/out")" "1.4.2 (17) -> 1.4.2 (18)"
assert "dry run names the missing Android secrets" contains "$(cat "$dir/out")" "blocked until these are set: ANDROID_KEYSTORE_PATH"
assert "dry run explains the iOS skip" contains "$(cat "$dir/out")" "Apple enrollment is not finished"
assert "dry run leaves VERSION alone" contains "$(version_file "$dir")" "APP_BUILD = 17"
assert "dry run calls no tool" equal "$(calls "$dir")" ""
assert "dry run makes no commit" equal "$(git -C "$dir" rev-list --count HEAD)" 1

# The bump arithmetic.
release "$dir" --bump patch
assert "patch" contains "$(cat "$dir/out")" "-> 1.4.3 (18)"
release "$dir" --bump minor
assert "minor" contains "$(cat "$dir/out")" "-> 1.5.0 (18)"
release "$dir" --bump major
assert "major" contains "$(cat "$dir/out")" "-> 2.0.0 (18)"
release "$dir" --bump sideways
assert "an unknown bump is refused" equal $? 1

# A version bump needs notes; a build bump does not.
dir=$(repo empty '')
release "$dir" --bump minor
assert "a minor release without notes fails" equal $? 1
assert "and says where the notes go" contains "$(cat "$dir/out")" "write them under '## Unreleased'"
release "$dir"
assert "a build release without notes is fine" equal $? 0

# Notes over Play's 500 characters are refused.
dir=$(repo long "- $(printf 'x%.0s' {1..520})")
release "$dir"
assert "notes over 500 characters fail" equal $? 1

# An Android release: bump, roll, commit, tag, build, upload.
dir=$(repo android '- Raid sign-up from the app.')
with_android_secrets release "$dir" --bump minor --platform android --execute
assert "android release exits 0" equal $? 0
assert "VERSION moves to 1.5.0" contains "$(version_file "$dir")" "APP_VERSION = 1.5.0"
assert "VERSION build moves to 18" contains "$(version_file "$dir")" "APP_BUILD = 18"
assert "VERSION keeps its comment" contains "$(version_file "$dir")" "// comment line"
changelog=$(cat "$dir/apps/CHANGELOG.md")
assert "the changelog opens a new Unreleased" contains "$changelog" "## Unreleased

## 1.5.0 (18), "
assert "the notes sit under the new version" contains "$changelog" "- Raid sign-up from the app."
assert "the bump is committed" equal "$(git -C "$dir" log -1 --format=%s)" "chore(apps): release 1.5.0 (18)"
assert "and tagged" equal "$(git -C "$dir" tag)" "apps-v1.5.0-18"
assert "with an annotated tag, which --follow-tags pushes" equal "$(git -C "$dir" cat-file -t apps-v1.5.0-18)" tag
assert "gradle builds the release bundle" contains "$(calls "$dir")" "gradlew -p $dir/apps/android bundleRelease"
assert "the Play upload goes to the internal track as draft" contains "$(calls "$dir")" "play-upload.ts --package gg.foreversixty.app --track internal --status draft --bundle $dir/apps/android/app/build/outputs/bundle/release/app-release.aab --mapping $dir/apps/android/app/build/outputs/mapping/release/mapping.txt"
assert "the upload sees the federated token in its environment" contains "$(calls "$dir")" "[token:ci-token]"
assert "the token is never an argument" lacks "$(calls "$dir" | sed 's/\[token:[^]]*\]//')" "ci-token"
assert "with a token given, gcloud is not asked" lacks "$(calls "$dir")" "gcloud"
assert "fastlane is not used for Play" lacks "$(calls "$dir")" "fastlane"
assert "Play gets the notes" equal "$(cat "$dir/stubs/play-notes")" "- Raid sign-up from the app."
assert "iOS is not touched" lacks "$(calls "$dir")" "xcodebuild"

# Both platforms before Apple enrollment: Android ships, iOS is skipped with the reason.
dir=$(repo both '- A fix.')
with_android_secrets release "$dir" --execute
assert "both, without Apple, still releases" equal $? 0
assert "iOS skipped with the reason" contains "$(cat "$dir/out")" "iOS: skipped. Apple enrollment is not finished"
assert "Android uploaded" contains "$(calls "$dir")" "play-upload.ts"

# Asking for iOS alone without the Apple account fails before anything changes.
dir=$(repo ios '- A fix.')
release "$dir" --platform ios --execute
assert "iOS alone without Apple fails" equal $? 1
assert "and leaves VERSION alone" contains "$(version_file "$dir")" "APP_BUILD = 17"

# With the Apple account, iOS archives, exports and goes to TestFlight.
dir=$(repo apple '- A fix.')
APPLE_TEAM_ID=T ASC_KEY_ID=K ASC_ISSUER_ID=I ASC_KEY_P8='-----BEGIN PRIVATE KEY-----' \
	release "$dir" --platform ios --execute
assert "iOS release exits 0" equal $? 0
assert "xcodebuild archives" contains "$(calls "$dir")" "xcodebuild archive"
assert "xcodebuild exports" contains "$(calls "$dir")" "xcodebuild -exportArchive"
assert "fastlane uploads to TestFlight with the notes" contains "$(calls "$dir")" "upload_to_testflight"
assert "the .p8 key is never an argument" lacks "$(calls "$dir")" "BEGIN PRIVATE KEY"

# Android without its signing secrets fails before anything changes.
dir=$(repo nokeys '- A fix.')
release "$dir" --platform android --execute
assert "android without secrets fails" equal $? 1
assert "and names them" contains "$(cat "$dir/out")" "set ANDROID_KEYSTORE_PATH"
assert "and the Play identity" contains "$(cat "$dir/out")" "PLAY_SA (or PLAY_ACCESS_TOKEN)"
assert "and leaves VERSION alone" contains "$(version_file "$dir")" "APP_BUILD = 17"

# A local release: the token is minted by impersonating PLAY_SA, and reaches only the upload.
dir=$(repo local '- A fix.')
with_android_impersonation release "$dir" --platform android --execute
assert "a local release exits 0" equal $? 0
assert "gcloud mints the token as the Play service account" contains "$(calls "$dir")" "gcloud auth print-access-token --impersonate-service-account=play-publisher@example.iam.gserviceaccount.com --scopes=https://www.googleapis.com/auth/androidpublisher"
assert "the upload sees the minted token" contains "$(grep play-upload "$dir/stubs/calls")" "[token:minted-token]"
assert "and the build never sees it" lacks "$(grep gradlew "$dir/stubs/calls")" "minted-token"

# A dry run mints no token.
dir=$(repo localdry '- A fix.')
with_android_impersonation release "$dir" --platform android
assert "a dry run asks gcloud nothing" lacks "$(calls "$dir")" "gcloud"

# gcloud refusing (a lapsed login, no Token Creator) stops the release before the upload.
dir=$(repo nologin '- A fix.')
mkdir -p "$dir/stubs" && touch "$dir/stubs/gcloud-fails"
with_android_impersonation release "$dir" --platform android --execute
assert "no token fails the release" equal $? 1
assert "and says what to do" contains "$(cat "$dir/out")" "no Play token"
assert "and uploads nothing" lacks "$(calls "$dir")" "play-upload.ts"

# A dirty VERSION is refused on execute.
dir=$(repo dirty '- A fix.')
echo "// local edit" >>"$dir/apps/VERSION"
with_android_secrets release "$dir" --platform android --execute
assert "a dirty VERSION is refused" equal $? 1

if [ "$FAILED" -gt 0 ]; then
	echo "$FAILED failed, $PASSED passed" >&2
	exit 1
fi
echo "ok $PASSED tests"
