import XCTest

/// Lane F15 review captures from the real system surfaces: the Offline data screen with its
/// switches and its remove alert, the share sheet with a guide's and a character's message, and
/// Spotlight. Siri cannot be driven from a UI test.
/// Runs on the shared simulator only through `make -C apps/ios captures`; the PNGs land in
/// design/reviews/2026-11-01-mobile-integration/build/ios.
@MainActor
final class SystemCaptureTests: XCTestCase {
    private static let out = URL(fileURLWithPath: #filePath)
        .deletingLastPathComponent() // UITests
        .deletingLastPathComponent() // apps/ios
        .deletingLastPathComponent() // apps
        .deletingLastPathComponent() // repo root
        .appending(path: "design/reviews/2026-11-01-mobile-integration/build/ios")
    private static let largest = ["-UIPreferredContentSizeCategoryName", "UICTContentSizeCategoryAccessibilityXXXL"]
    private let springboard = XCUIApplication(bundleIdentifier: "com.apple.springboard")

    override func setUp() async throws {
        continueAfterFailure = false
        try FileManager.default.createDirectory(at: Self.out, withIntermediateDirectories: true)
    }

    private func save(_ name: String) throws {
        try XCUIScreen.main.screenshot().pngRepresentation.write(to: Self.out.appending(path: "\(name)-ios.png"))
    }

    /// On Home, with `link` opened through the debug `-FSOpenLink`.
    private func launch(link: String? = nil) -> XCUIApplication {
        let app = XCUIApplication()
        app.launchArguments = ["-FSTab", "home"] + (link.map { ["-FSOpenLink", $0] } ?? [])
        app.launch()
        return app
    }

    /// The share button in the bar, then the system sheet with the page's message.
    private func share(_ app: XCUIApplication, _ name: String) throws {
        let button = app.buttons.matching(NSPredicate(format: "label CONTAINS[c] %@", "Share")).firstMatch
        XCTAssertTrue(button.waitForExistence(timeout: 20))
        sleep(2)
        button.tap()
        sleep(3)
        try save(name)
    }

    func testGuideShareSheet() throws {
        try share(launch(link: "https://foreversixty.gg/app/guides/shaman/elemental"), "f15-share-guide-system")
    }

    func testCharacterShareSheet() throws {
        try share(launch(link: "https://foreversixty.gg/app/player/us/normal/voltaris"), "f15-share-character-system")
    }

    /// The Guides root bar (its character selector and search magnifier) went with the Guides tab:
    /// Guides is now a menu tool pushed under a tab, its bar a pushed screen's.
    func testGuidesRootBar() throws {
        throw XCTSkip("the Guides tab and its root bar were cut by the flows spec (2026-11-08-app-flows.md appendix C, \"The BiS, Guides and Guild tabs\")")
    }

    /// Offline data, pushed from the profile menu's Account rows.
    func testOfflineDataAndRemoveAlert() throws {
        for large in [false, true] {
            let app = CaptureLaunch.signedIn(menu: "offlineData", extra: large ? Self.largest : [])
            XCTAssertTrue(app.staticTexts["Offline data"].firstMatch.waitForExistence(timeout: 20), "Offline data did not open")
            sleep(3)
            try save(large ? "f15-offlinedata-app-large" : "f15-offlinedata-app")
            let remove = app.buttons["Remove downloaded data"].firstMatch
            for _ in 0..<8 where !remove.isHittable { app.swipeUp() }
            if remove.isEnabled {
                remove.tap()
                XCTAssertTrue(app.alerts.firstMatch.waitForExistence(timeout: 5))
                try save(large ? "f15-offlinedata-remove-system-large" : "f15-offlinedata-remove-system")
                app.alerts.firstMatch.buttons["Cancel"].tap()
            }
            app.terminate()
        }
    }

    /// The app indexes on launch; Spotlight then finds the raids and characters. The simulator's
    /// home screen may not offer the search field to a swipe; then this records nothing.
    func testSpotlight() throws {
        let app = launch()
        sleep(8)
        XCUIDevice.shared.press(.home)
        sleep(1)
        springboard.swipeDown(velocity: .slow)
        let field = springboard.searchFields.firstMatch
        guard field.waitForExistence(timeout: 6) else {
            throw XCTSkip("the simulator's home screen offers no Spotlight field to a swipe")
        }
        field.typeText("Onyxia")
        sleep(3)
        try save("f15-spotlight-raid-system")
        app.terminate()
    }
}
