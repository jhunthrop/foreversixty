# Local entry points for the iOS app; CI (.github/workflows/ios.yml) runs the same steps.
#
# Locally every lane shares ONE simulator, "Forever Sixty Test (iPhone 17)", and takes a lock
# before touching it (tools/mobile-device.sh), so parallel lanes queue instead of colliding.
# CI sets SIMULATOR itself and runs without the lock: its runners are single-tenant.
DEVICE := ../../tools/mobile-device.sh
ifdef CI
SIMULATOR ?= platform=iOS Simulator,name=iPhone 17
LOCKED :=
else
SIMULATOR ?= platform=iOS Simulator,id=$(shell $(DEVICE) ios-udid)
LOCKED := $(DEVICE) ios-run
endif
XCODEBUILD = xcodebuild -project ForeverSixty.xcodeproj -scheme ForeverSixty \
	-destination '$(SIMULATOR)' -derivedDataPath build/DerivedData

.PHONY: project build test clean device captures snapshots

project:
	xcodegen generate

build: project
	$(XCODEBUILD) build

test: project
	$(LOCKED) $(XCODEBUILD) test

# Lane F4's tap-driven review captures (UITests/RaidCaptureTests) on the shared simulator.
captures: project
	$(LOCKED) xcodebuild -project ForeverSixty.xcodeproj -scheme ForeverSixtyCaptures \
		-destination '$(SIMULATOR)' -derivedDataPath build/DerivedData test

# Snapshot goldens (AppTests/**/__Snapshots__) are rendered on CI's pinned runtime (Xcode 26.6,
# iOS 26.5; ios.yml). The shared local simulator runs iOS 26.5 too (tools/mobile-device.sh), so
# local runs compare the same pixels; iOS 27's navigation bars and line metrics differ. To re-record,
# push the branch, then run this: it dispatches ios.yml in record mode on BRANCH, waits, and
# copies the goldens CI rewrote into every AppTests/**/__Snapshots__ for review and commit.
BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD)
snapshots:
	gh workflow run ios.yml --ref '$(BRANCH)' -f record=true
	sleep 10
	run=$$(gh run list --workflow ios.yml --branch '$(BRANCH)' --event workflow_dispatch --limit 1 --json databaseId -q '.[0].databaseId') && \
	gh run watch "$$run" --exit-status >/dev/null; \
	rm -rf build/ios-snapshots && gh run download "$$run" -n ios-snapshots -D build/ios-snapshots && \
	cp -R build/ios-snapshots/. AppTests/ && git status --short AppTests

# Prints the shared simulator's UDID, for screenshots: wrap simctl calls in
# `../../tools/mobile-device.sh ios-run xcrun simctl ...`.
device:
	@$(DEVICE) ios-udid

clean:
	rm -rf build ForeverSixty.xcodeproj Generated
