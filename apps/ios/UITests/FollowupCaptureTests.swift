import XCTest

/// Review captures for the 2026-10-31 mobile follow-ups: the selector's full detail line on the
/// Raids root at the default size and at AX5, Home and BiS on the same band for a level-58
/// character, and the composition's named buff tiles with live icons. Runs on the shared
/// simulator only (`make -C apps/ios captures`); the PNGs land in
/// design/reviews/2026-10-31-mobile-followups. Not part of `make test`.
@MainActor
final class FollowupCaptureTests: XCTestCase {
    private static let out = URL(fileURLWithPath: #filePath)
        .deletingLastPathComponent() // UITests
        .deletingLastPathComponent() // apps/ios
        .deletingLastPathComponent() // apps
        .deletingLastPathComponent() // repo root
        .appending(path: "design/reviews/2026-10-31-mobile-followups")
    private static let raidLink = "https://foreversixty.gg/app/raid/2/31"
    private static let largestText = ["-UIPreferredContentSizeCategoryName", "UICTContentSizeCategoryAccessibilityXXXL"]

    override func setUp() async throws {
        continueAfterFailure = false
        try FileManager.default.createDirectory(at: Self.out, withIntermediateDirectories: true)
    }

    private func save(_ name: String) throws {
        try XCUIScreen.main.screenshot().pngRepresentation.write(to: Self.out.appending(path: "\(name)-ios.png"))
    }

    /// Launches signed in on the fixtures (a session kept from an earlier run is reused).
    private func signedIn(_ extra: [String] = []) -> XCUIApplication {
        let app = XCUIApplication()
        app.launchArguments = ["-FSTab", "you"] + extra
        app.launch()
        let signIn = app.buttons.matching(NSPredicate(format: "label CONTAINS[c] 'Battle.net'")).firstMatch
        if !app.buttons["Sign out"].waitForExistence(timeout: 5), signIn.waitForExistence(timeout: 15) {
            signIn.tap()
        }
        return app
    }

    private func pick(_ app: XCUIApplication, from current: String, to name: String) {
        app.navigationBars.buttons.matching(NSPredicate(format: "label CONTAINS %@", current)).firstMatch.tap()
        let row = app.buttons.matching(NSPredicate(format: "label CONTAINS %@", name)).firstMatch
        XCTAssertTrue(row.waitForExistence(timeout: 5))
        row.tap()
        sleep(3)
    }

    /// Home and BiS for Obnoxious Yell (level 58): both on `50 to 59`.
    func testHomeAndBisShareTheBand() throws {
        let app = signedIn()
        app.tabBars.buttons["Home"].tap()
        sleep(5)
        try save("home-band-58")
        app.swipeUp()
        sleep(2)
        try save("home-band-58-scrolled")
        app.tabBars.buttons["BiS"].tap()
        sleep(5)
        try save("bis-band-58")
    }

    /// The selector on Raids for Yellfury (`Hunter · 31`), at the default size and at AX5.
    func testSelectorDetailOnRaids() throws {
        for (size, arguments) in [("default", [String]()), ("ax5", Self.largestText)] {
            let app = signedIn(arguments)
            app.tabBars.buttons["Raids"].tap()
            sleep(4)
            pick(app, from: "Obnoxious Yell", to: "Yellfury")
            try save("selector-raids-\(size)")
            app.tabBars.buttons["Home"].tap()
            sleep(4)
            try save("selector-home-\(size)")
            pick(app, from: "Yellfury", to: "Obnoxious Yell")
            app.terminate()
        }
    }

    /// The composition block with live buff icons, each named under its tile.
    func testNamedBuffTiles() throws {
        let app = XCUIApplication()
        app.launchArguments = ["-FSOpenLink", Self.raidLink]
        app.launch()
        XCTAssertTrue(app.staticTexts["ONYXIA'S LAIR"].waitForExistence(timeout: 15))
        let buffs = app.staticTexts["BUFFS AND UTILITY"]
        for _ in 0..<8 where !buffs.isHittable { app.swipeUp() }
        sleep(5)
        try save("raid-buffs")
        // VoiceOver's element for the tile carries the full name, live art or not.
        let fort = app.descendants(matching: .any).matching(NSPredicate(format: "label BEGINSWITH 'Power Word: Fortitude'"))
        XCTAssertTrue(fort.firstMatch.exists)
    }
}
