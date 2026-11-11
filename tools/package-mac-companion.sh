#!/usr/bin/env bash
# tools/package-mac-companion.sh <version> <universal-binary> <out-dir>
#
# Turns the companion's universal macOS binary into what the release publishes: the app
# bundle "Forever Sixty.app", signed with the hardened runtime and a timestamp, notarized
# and stapled, then shipped twice:
#
#   <out-dir>/foreversixty-companion_darwin_universal.dmg      the download /setup offers:
#                                                              open it, drag the app to Applications
#   <out-dir>/foreversixty-companion_darwin_universal.app.zip  what the companion updates itself
#                                                              from (companion/internal/updater/bundle.go)
#
# companion.yml's build-mac job runs it on a macOS runner; it runs the same on a Mac by hand.
#
# Environment:
#   MAC_SIGNING_IDENTITY  the codesign identity (a Developer ID Application certificate's
#                         name or SHA-1). Empty: the app is signed ad hoc, nothing is
#                         notarized, and the output is for testing only, never for the feed.
#   MAC_KEYCHAIN          the keychain holding that identity (optional).
#   NOTARY_KEY_FILE, NOTARY_KEY_ID, NOTARY_ISSUER
#                         the App Store Connect API key notarytool signs in with (the .p8
#                         file, its key id, its issuer id). Required with an identity.
#
# Exit status: 0 packaged; 1 a tool failed (signing, notarization, the disk image); 64 usage.
set -euo pipefail

readonly APP_NAME="Forever Sixty"
readonly BUNDLE_ID=gg.foreversixty.companion
readonly EXECUTABLE=foreversixty-companion
readonly ASSET_STEM=foreversixty-companion_darwin_universal
# Go 1.25, the companion's toolchain, runs on macOS 12 and later.
readonly MIN_MACOS=12.0
readonly ICON=companion/internal/icon/assets/seal.icns

die() {
	echo "package-mac-companion: $*" >&2
	exit 64
}

(($# == 3)) || die "usage: tools/package-mac-companion.sh <version> <universal-binary> <out-dir>"
readonly VERSION=$1 BINARY=$2 OUT=$3
[[ $VERSION =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]] || die "'$VERSION' is not a version like 0.3.0"
[[ -x $BINARY ]] || die "$BINARY is not an executable"
ROOT=$(cd "$(dirname "$0")/.." && pwd)
readonly ROOT
[[ -s $ROOT/$ICON ]] || die "$ICON is missing"

readonly IDENTITY=${MAC_SIGNING_IDENTITY:-}
if [[ -n $IDENTITY ]]; then
	[[ -s ${NOTARY_KEY_FILE:-} && -n ${NOTARY_KEY_ID:-} && -n ${NOTARY_ISSUER:-} ]] ||
		die "a signed build is notarized: set NOTARY_KEY_FILE, NOTARY_KEY_ID and NOTARY_ISSUER"
fi

work=$(mktemp -d)
trap 'rm -rf "$work"' EXIT
readonly APP="$work/$APP_NAME.app"

bundle() {
	mkdir -p "$APP/Contents/MacOS" "$APP/Contents/Resources"
	cp "$BINARY" "$APP/Contents/MacOS/$EXECUTABLE"
	cp "$ROOT/$ICON" "$APP/Contents/Resources/seal.icns"
	# LSUIElement is false: the app has a Dock icon and a menu bar while it runs, and the
	# seal in the menu bar's status area besides (companion/internal/shell).
	cat >"$APP/Contents/Info.plist" <<PLIST
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
  <key>CFBundleIdentifier</key><string>$BUNDLE_ID</string>
  <key>CFBundleName</key><string>$APP_NAME</string>
  <key>CFBundleDisplayName</key><string>$APP_NAME</string>
  <key>CFBundleExecutable</key><string>$EXECUTABLE</string>
  <key>CFBundleIconFile</key><string>seal</string>
  <key>CFBundlePackageType</key><string>APPL</string>
  <key>CFBundleShortVersionString</key><string>$VERSION</string>
  <key>CFBundleVersion</key><string>$VERSION</string>
  <key>LSMinimumSystemVersion</key><string>$MIN_MACOS</string>
  <key>LSApplicationCategoryType</key><string>public.app-category.games</string>
  <key>LSUIElement</key><false/>
  <key>NSHighResolutionCapable</key><true/>
  <key>NSHumanReadableCopyright</key><string>Commish LLC</string>
</dict>
</plist>
PLIST
	plutil -lint "$APP/Contents/Info.plist"
}

codesign_item() { # codesign_item <path>
	if [[ -z $IDENTITY ]]; then
		codesign --force --sign - "$1"
		return
	fi
	local keychain=()
	[[ -n ${MAC_KEYCHAIN:-} ]] && keychain=(--keychain "$MAC_KEYCHAIN")
	# The ${a[@]+...} form: macOS's own bash 3.2 calls an empty array unbound under set -u.
	codesign --force --timestamp --options runtime ${keychain[@]+"${keychain[@]}"} --sign "$IDENTITY" "$1"
}

notarize() { # notarize <file>: submits, waits, and prints Apple's log when it is refused
	local result json id status
	# A refusal or a timeout is printed on stderr, as JSON like the answer; both are read.
	result=$(xcrun notarytool submit "$1" --key "$NOTARY_KEY_FILE" --key-id "$NOTARY_KEY_ID" \
		--issuer "$NOTARY_ISSUER" --wait --timeout 30m --output-format json 2>&1) || true
	json=$(grep '^{' <<<"$result" | tail -1 || true)
	id=$(jq -r '.id // empty' <<<"$json" 2>/dev/null || true)
	status=$(jq -r '.status // empty' <<<"$json" 2>/dev/null || true)
	echo "notarization of $(basename "$1"): ${status:-no status} (${id:-no id})"
	if [[ $status != Accepted ]]; then
		echo "$result" >&2
		if [[ -z $status && -n $id ]]; then
			# The 2026-11-07 run: Apple took longer than the wait on a team's first submissions.
			echo "package-mac-companion: Apple was still processing $id when the wait ran out; rerun the job" >&2
		elif [[ -n $id ]]; then
			xcrun notarytool log "$id" --key "$NOTARY_KEY_FILE" --key-id "$NOTARY_KEY_ID" \
				--issuer "$NOTARY_ISSUER" >&2 || true
		fi
		echo "package-mac-companion: Apple did not accept $(basename "$1")" >&2
		exit 1
	fi
}

# The disk image's window: the background tools/art/dmg_background.py draws, and the window size and
# icon places it writes beside it (dmg-layout.sh), the app at the left, Applications at the right;
# 128 px icons, 13 pt labels.
readonly DMG_BACKGROUND=companion/packaging/macos/dmg-background
# shellcheck source=SCRIPTDIR/../companion/packaging/macos/dmg-layout.sh
source "$ROOT/companion/packaging/macos/dmg-layout.sh"

disk_image() { # disk_image <dmg>: the designed drag window, under the seal as the volume icon
	local stage=$work/dmg background=$work/background.tiff
	command -v create-dmg >/dev/null || die "create-dmg is not installed (brew install create-dmg)"
	mkdir -p "$stage"
	ditto "$APP" "$stage/$APP_NAME.app"
	# One TIFF with both scales, so a Retina screen draws the 2x background.
	tiffutil -cathidpicheck "$ROOT/$DMG_BACKGROUND.png" "$ROOT/$DMG_BACKGROUND@2x.png" -out "$background"
	# create-dmg lays the window out through Finder (background, bounds, icon places, no toolbar or
	# status bar), writes the .DS_Store into the image, sets the volume icon with its custom-icon
	# flag, and compresses the result read-only. It retries hdiutil's "Resource busy" by itself.
	rm -f "$1"
	create-dmg \
		--volname "$APP_NAME" \
		--volicon "$ROOT/$ICON" \
		--background "$background" \
		--window-pos 200 120 \
		--window-size "${DMG_WINDOW[@]}" \
		--icon-size 128 \
		--text-size 13 \
		--icon "$APP_NAME.app" "${DMG_APP[@]}" \
		--hide-extension "$APP_NAME.app" \
		--app-drop-link "${DMG_DROP[@]}" \
		--no-internet-enable \
		--format UDZO \
		--hdiutil-retries 8 \
		"$1" "$stage"
}

mkdir -p "$OUT"
readonly ZIP="$OUT/$ASSET_STEM.app.zip" DMG="$OUT/$ASSET_STEM.dmg"

bundle
codesign_item "$APP"
codesign --verify --strict --verbose=2 "$APP"
if [[ -n $IDENTITY ]]; then
	# The app is notarized and stapled first, so the copy inside the disk image and the
	# copy in the update zip both open offline.
	ditto -c -k --keepParent "$APP" "$work/notarize.zip"
	notarize "$work/notarize.zip"
	xcrun stapler staple "$APP"
	xcrun stapler validate "$APP"
fi
# --norsrc: no extended attributes or AppleDouble files; the bundle needs none, and the
# updater unpacks plain files, folders and links.
ditto -c -k --norsrc --keepParent "$APP" "$ZIP"
disk_image "$DMG"
codesign_item "$DMG"
if [[ -n $IDENTITY ]]; then
	notarize "$DMG"
	xcrun stapler staple "$DMG"
	xcrun stapler validate "$DMG"
	# What Gatekeeper will say on a player's Mac: both must be accepted as notarized
	# Developer ID software.
	spctl --assess --type execute --verbose=2 "$APP"
	spctl --assess --type open --context context:primary-signature --verbose=2 "$DMG"
else
	echo "package-mac-companion: signed ad hoc and not notarized; for testing only"
fi
ls -l "$OUT"
