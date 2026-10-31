import XCTest

/// Lane F4 review captures that need a real tap: the More menu, the calendar sheet, the officer
/// move menu and the raid notification answered from the lock screen's long-press. Runs on the
/// shared simulator only through `make -C apps/ios captures`; the PNGs land in
/// design/reviews/2026-10-31-mobile-raids. Not part of `make test`.
@MainActor
final class RaidCaptureTests: XCTestCase {
    private static let out = URL(fileURLWithPath: #filePath)
        .deletingLastPathComponent() // UITests
        .deletingLastPathComponent() // apps/ios
        .deletingLastPathComponent() // apps
        .deletingLastPathComponent() // repo root
        .appending(path: "design/reviews/2026-10-31-mobile-raids")
    private static let raidLink = "https://foreversixty.gg/app/raid/2/31"
    private let springboard = XCUIApplication(bundleIdentifier: "com.apple.springboard")

    override func setUp() async throws {
        continueAfterFailure = false
    }

    private func save(_ name: String) throws {
        try XCUIScreen.main.screenshot().pngRepresentation.write(to: Self.out.appending(path: name))
    }

    private func launch(_ arguments: [String]) -> XCUIApplication {
        let app = XCUIApplication()
        app.launchArguments = arguments
        app.launch()
        return app
    }

    /// Taps a system alert's allowing button if one shows (calendar, notifications).
    private func allowSystemAlert() {
        let alert = springboard.alerts.firstMatch
        guard alert.waitForExistence(timeout: 4) else { return }
        for label in ["Allow Full Access", "Allow", "OK", "Continue"] where alert.buttons[label].exists {
            alert.buttons[label].tap()
            return
        }
    }

    func testMoreMenuAndCalendarSheet() throws {
        let app = launch(["-FSOpenLink", Self.raidLink])
        XCTAssertTrue(app.staticTexts["ONYXIA'S LAIR"].waitForExistence(timeout: 15))
        app.buttons["More"].firstMatch.tap()
        XCTAssertTrue(app.buttons["Share the raid"].waitForExistence(timeout: 5))
        try save("raid-page-more-menu-ios.png")
        app.buttons["Add to calendar"].firstMatch.tap()
        allowSystemAlert()
        sleep(3)
        try save("raid-page-calendar-sheet-ios.png")
    }

    /// The roster with real spec icons and the composition block with real buff icons (and the
    /// labels for the two the icon export does not publish yet), loaded from the CDN.
    func testRosterAndCompositionWithIcons() throws {
        let app = launch(["-FSOpenLink", Self.raidLink])
        XCTAssertTrue(app.staticTexts["ONYXIA'S LAIR"].waitForExistence(timeout: 15))
        let composition = app.staticTexts["RAID COMPOSITION"]
        for _ in 0..<8 where !composition.isHittable { app.swipeUp() }
        sleep(4)
        try save("raid-composition-ios.png")
        app.swipeUp()
        sleep(3)
        try save("raid-page-roster-ios.png")
    }

    func testOfficerMoveMenu() throws {
        let app = launch(["-FSOpenLink", Self.raidLink, "-FSRaidFixture", "officer"])
        XCTAssertTrue(app.staticTexts["ONYXIA'S LAIR"].waitForExistence(timeout: 15))
        let place = app.buttons["Kraggor, In. Change the place"]
        for _ in 0..<12 where !place.isHittable { app.swipeUp() }
        place.tap()
        XCTAssertTrue(app.buttons["Move to Bench"].waitForExistence(timeout: 5))
        try save("officer-move-menu-ios.png")
    }

    func testNotificationInFromTheLongPress() throws {
        let app = launch(["-FSPostRaidPush"])
        allowSystemAlert()
        sleep(1)
        XCUIDevice.shared.press(.home)
        sleep(4)
        // The notification centre: drag down from the top edge.
        let top = springboard.coordinate(withNormalizedOffset: CGVector(dx: 0.3, dy: 0.01))
        top.press(forDuration: 0.1, thenDragTo: springboard.coordinate(withNormalizedOffset: CGVector(dx: 0.3, dy: 0.7)))
        let note = springboard.otherElements["Notification"].firstMatch.exists
            ? springboard.otherElements["Notification"].firstMatch
            : springboard.staticTexts["Onyxia's Lair"].firstMatch
        XCTAssertTrue(note.waitForExistence(timeout: 10))
        try save("notification-delivered-ios.png")
        note.press(forDuration: 1.5)
        let signIn = springboard.buttons["In"]
        XCTAssertTrue(signIn.waitForExistence(timeout: 5))
        try save("notification-actions-ios.png")
        signIn.tap()
        sleep(6)
        try save("notification-answer-in-ios.png")
        _ = app
    }

    /// Integration C: signed in on the fixtures, the Raids tab under the one character selector
    /// (the same toolbar slot as Home and BiS), then Home's for the side-by-side.
    func testRaidsTabUnderTheSelector() throws {
        let app = launch(["-FSTab", "you"])
        // A session kept from an earlier run is already signed in (You shows Sign out).
        let signIn = app.buttons.matching(NSPredicate(format: "label CONTAINS[c] 'Battle.net'")).firstMatch
        if !app.buttons["Sign out"].waitForExistence(timeout: 5) {
            XCTAssertTrue(signIn.waitForExistence(timeout: 15))
            signIn.tap()
        }
        app.tabBars.buttons["Raids"].tap()
        // The guild head reads as one element, so wait on the next raid card's link.
        XCTAssertTrue(app.buttons.matching(NSPredicate(format: "label BEGINSWITH 'Who is coming'")).firstMatch.waitForExistence(timeout: 15))
        sleep(3)
        try save("selector-raids-ios.png")
        // The switch: Raids follows the picked character's guild (Yellfury has none).
        app.navigationBars.buttons.matching(NSPredicate(format: "label CONTAINS 'Obnoxious Yell'")).firstMatch.tap()
        let alt = app.buttons.matching(NSPredicate(format: "label CONTAINS 'Yellfury'")).firstMatch
        XCTAssertTrue(alt.waitForExistence(timeout: 5))
        sleep(1)
        try save("selector-open-raids-ios.png")
        alt.tap()
        sleep(4)
        try save("selector-raids-switched-ios.png")
        app.navigationBars.buttons.matching(NSPredicate(format: "label CONTAINS 'Yellfury'")).firstMatch.tap()
        let main = app.buttons.matching(NSPredicate(format: "label CONTAINS 'Obnoxious Yell'")).firstMatch
        XCTAssertTrue(main.waitForExistence(timeout: 5))
        main.tap()
        sleep(3)
        app.tabBars.buttons["Home"].tap()
        sleep(4)
        try save("selector-home-ios.png")
        app.tabBars.buttons["BiS"].tap()
        sleep(4)
        try save("selector-bis-ios.png")
    }
}
