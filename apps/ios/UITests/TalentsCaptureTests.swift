import XCTest

/// Review captures for F6 Talents in the running app: Home's Talents row, the pushed screen, the
/// pinned compare bar, a rank sheet with the tapped cell left visible, and the same at AX5. Runs
/// on the shared simulator only; the PNGs land in design/reviews/2026-10-31-mobile-talents. Not
/// part of `make test`.
@MainActor
final class TalentsCaptureTests: XCTestCase {
    private static let out = URL(fileURLWithPath: #filePath)
        .deletingLastPathComponent() // UITests
        .deletingLastPathComponent() // apps/ios
        .deletingLastPathComponent() // apps
        .deletingLastPathComponent() // repo root
        .appending(path: "design/reviews/2026-10-31-mobile-talents")
    private static let largestText = ["-UIPreferredContentSizeCategoryName", "UICTContentSizeCategoryAccessibilityXXXL"]

    override func setUp() async throws {
        continueAfterFailure = false
        try FileManager.default.createDirectory(at: Self.out, withIntermediateDirectories: true)
    }

    private func save(_ name: String) throws {
        try XCUIScreen.main.screenshot().pngRepresentation.write(to: Self.out.appending(path: "\(name)-ios.png"))
    }

    /// Launches signed in on the fixtures (a session kept from an earlier run is reused), on Home.
    private func home(_ extra: [String] = []) -> XCUIApplication {
        let app = XCUIApplication()
        app.launchArguments = ["-FSTab", "you"] + extra
        app.launch()
        let signIn = app.buttons.matching(NSPredicate(format: "label CONTAINS[c] 'Battle.net'")).firstMatch
        if !app.buttons["Sign out"].waitForExistence(timeout: 5), signIn.waitForExistence(timeout: 15) {
            signIn.tap()
        }
        app.tabBars.buttons["Home"].tap()
        sleep(4)
        return app
    }

    private func openTalents(_ app: XCUIApplication) {
        let row = app.buttons.matching(NSPredicate(format: "label BEGINSWITH 'Talents'")).firstMatch
        var swipes = 0
        while !row.isHittable, swipes < 6 {
            app.swipeUp()
            swipes += 1
        }
        XCTAssertTrue(row.waitForExistence(timeout: 5))
        row.tap()
        sleep(4)
    }

    func testTalentsScreen() throws {
        let app = home()
        app.swipeUp()
        sleep(1)
        try save("home-talents-row")
        openTalents(app)
        try save("app-talents-differs")
        app.swipeUp()
        sleep(2)
        try save("app-talents-differs-scrolled")
        let goreDrinker = app.buttons.matching(NSPredicate(format: "label BEGINSWITH 'Gore Drinker'")).firstMatch
        if goreDrinker.waitForExistence(timeout: 3) {
            goreDrinker.tap()
            sleep(2)
            try save("app-talents-rank")
        }
    }

    func testTalentsScreenAtLargestText() throws {
        let app = home(Self.largestText)
        openTalents(app)
        try save("app-talents-differs-large")
        app.swipeUp()
        sleep(2)
        try save("app-talents-differs-large-tree")
    }
}
