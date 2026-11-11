import XCTest

/// Review captures for the 2026-10-31 mobile follow-ups: the profile menu's character lines on the
/// Raid root at the default size and at AX5, Home and Gear on the same band for a level-58
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

    /// Home and Gear for Gorvash (level 58): both on `50 to 59`.
    func testHomeAndBisShareTheBand() throws {
        let app = CaptureLaunch.signedIn(tab: "home")
        sleep(5)
        try save("home-band-58")
        app.swipeUp()
        sleep(2)
        try save("home-band-58-scrolled")
        app.tabBars.buttons["Gear"].tap()
        sleep(5)
        try save("bis-band-58")
    }

    /// The switch from the Raid tab's profile menu: the menu with Yellfury's row (`Hunter · 31`),
    /// then Raid and Home re-answered for Yellfury, at the default size and at AX5. The toolbar
    /// selector these captures first showed is gone (flows spec 2026-11-08 section 3).
    func testSelectorDetailOnRaids() throws {
        for (size, arguments) in [("default", [String]()), ("ax5", Self.largestText)] {
            let app = CaptureLaunch.signedIn(tab: "guild", extra: arguments)
            sleep(4)
            CaptureLaunch.openMenu(app)
            sleep(1)
            try save("selector-menu-raids-\(size)")
            CaptureLaunch.pick(app, "Yellfury")
            try save("selector-raids-\(size)")
            app.tabBars.buttons["Home"].tap()
            sleep(4)
            try save("selector-home-\(size)")
            CaptureLaunch.switchCharacter(app, to: "Gorvash")
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
