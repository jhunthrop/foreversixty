import XCTest

/// Lane F12 review capture that needs the running app: the item page's strip chip scrolls to its
/// section (a snapshot cannot scroll). Runs on the shared simulator only through
/// `make -C apps/ios captures`; the PNG lands in design/reviews/2026-10-31-mobile-wave3-search/build/ios.
@MainActor
final class SearchCaptureTests: XCTestCase {
    private static let out = URL(fileURLWithPath: #filePath)
        .deletingLastPathComponent() // UITests
        .deletingLastPathComponent() // apps/ios
        .deletingLastPathComponent() // apps
        .deletingLastPathComponent() // repo root
        .appending(path: "design/reviews/2026-10-31-mobile-wave3-search/build/ios")

    override func setUp() async throws {
        continueAfterFailure = false
        try FileManager.default.createDirectory(at: Self.out, withIntermediateDirectories: true)
    }

    /// Search, open Spritecaster Cape, tap `Best in slot for`: the section scrolls into view.
    func testBestInSlotChipScrollsToItsSection() throws {
        let app = XCUIApplication()
        app.launchArguments = ["-FSTab", "home"]
        app.launch()
        let search = app.buttons["Search"].firstMatch
        XCTAssertTrue(search.waitForExistence(timeout: 20))
        search.tap()
        let field = app.textFields.firstMatch.exists ? app.textFields.firstMatch : app.textViews.firstMatch
        XCTAssertTrue(field.waitForExistence(timeout: 5))
        field.typeText("spritec")
        let cape = app.buttons.matching(NSPredicate(format: "label CONTAINS 'Spritecaster Cape'")).firstMatch
        XCTAssertTrue(cape.waitForExistence(timeout: 10))
        cape.tap()
        let chip = app.buttons["Best in slot for"].firstMatch
        XCTAssertTrue(chip.waitForExistence(timeout: 10))
        sleep(2)
        chip.tap()
        sleep(2)
        let druid = app.buttons.matching(NSPredicate(format: "label BEGINSWITH 'Balance Druid'")).firstMatch
        XCTAssertTrue(druid.waitForExistence(timeout: 5))
        XCTAssertTrue(druid.isHittable, "the chip did not scroll the Best in slot for rows on screen")
        try XCUIScreen.main.screenshot().pngRepresentation.write(to: Self.out.appending(path: "f12-5-bis-for-scrolled-ios.png"))
    }
}
