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

.PHONY: project build test clean device captures

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

# Prints the shared simulator's UDID, for screenshots: wrap simctl calls in
# `../../tools/mobile-device.sh ios-run xcrun simctl ...`.
device:
	@$(DEVICE) ios-udid

clean:
	rm -rf build ForeverSixty.xcodeproj Generated
