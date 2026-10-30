# Local entry points for the iOS app; CI (.github/workflows/ios.yml) runs the same steps.
SIMULATOR ?= platform=iOS Simulator,name=iPhone 17
XCODEBUILD = xcodebuild -project ForeverSixty.xcodeproj -scheme ForeverSixty \
	-destination '$(SIMULATOR)' -derivedDataPath build/DerivedData

.PHONY: project build test clean

project:
	xcodegen generate

build: project
	$(XCODEBUILD) build

test: project
	$(XCODEBUILD) test

clean:
	rm -rf build ForeverSixty.xcodeproj Generated
