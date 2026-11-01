import XCTest

/// Lane F15 review captures from the real system surfaces: Spotlight's app entities, the Offline
/// data screen and its remove alert, and the share sheet with a guide's message. Siri cannot be
/// driven from a UI test. Runs on the shared simulator only through `make -C apps/ios captures`;
/// the PNGs land in design/reviews/2026-11-01-mobile-integration/build/ios.
@MainActor
final class SystemCaptureTests: XCTestCase {
    private static let out = URL(fileURLWithPath: #filePath)
        .deletingLastPathComponent() // UITests
        .deletingLastPathComponent() // apps/ios
        .deletingLastPathComponent() // apps
        .deletingLastPathComponent() // repo root
        .appending(path: "design/reviews/2026-11-01-mobile-integration/build/ios")
    private let springboard = XCUIApplication(bundleIdentifier: "com.apple.springboard")

    override func setUp() async throws {
        continueAfterFailure = false
        try FileManager.default.createDirectory(at: Self.out, withIntermediateDirectories: true)
    }

    private func save(_ name: String) throws {
        try XCUIScreen.main.screenshot().pngRepresentation.write(to: Self.out.appending(path: "\(name)-ios.png"))
    }

    private func begins(_ text: String) -> NSPredicate { NSPredicate(format: "label BEGINSWITH %@", text) }

    private func launch(tab: String) -> XCUIApplication {
        let app = XCUIApplication()
        app.launchArguments = ["-FSTab", tab]
        app.launch()
        return app
    }

    /// The app indexes on launch; Spotlight then finds the characters, raids and guides.
    func testSpotlight() throws {
        let app = launch(tab: "home")
        sleep(8)
        XCUIDevice.shared.press(.home)
        sleep(1)
        springboard.swipeDown(velocity: .slow)
        let field = springboard.searchFields.firstMatch
        if !field.waitForExistence(timeout: 6) {
            try save("f15-spotlight-debug")
            XCTFail("no Spotlight field")
        }
        field.typeText("Onyxia")
        sleep(3)
        try save("f15-spotlight-raid-system")
        field.typeText(String(repeating: XCUIKeyboardKey.delete.rawValue, count: "Onyxia".count))
        field.typeText("Obnoxious")
        sleep(3)
        try save("f15-spotlight-character-system")
        app.terminate()
    }

    func testOfflineDataAndRemoveAlert() throws {
        let app = launch(tab: "you")
        let row = app.buttons.containing(begins("Offline data")).firstMatch
        for _ in 0..<4 where !row.exists { app.swipeUp() }
        XCTAssertTrue(row.waitForExistence(timeout: 20))
        row.tap()
        sleep(3)
        try save("f15-offlinedata-app")
        let remove = app.buttons["Remove downloaded data"].firstMatch
        for _ in 0..<4 where !remove.isHittable { app.swipeUp() }
        if remove.isEnabled {
            remove.tap()
            XCTAssertTrue(app.alerts.firstMatch.waitForExistence(timeout: 5))
            try save("f15-offlinedata-remove-system")
            app.alerts.firstMatch.buttons["Cancel"].tap()
        }
    }

    func testGuideShareSheet() throws {
        let app = launch(tab: "guides")
        let guide = app.buttons.containing(begins("Elemental")).firstMatch
        XCTAssertTrue(guide.waitForExistence(timeout: 20))
        guide.tap()
        let share = app.buttons["Share"].firstMatch
        XCTAssertTrue(share.waitForExistence(timeout: 10))
        share.tap()
        sleep(3)
        try save("f15-share-guide-system")
    }
}
