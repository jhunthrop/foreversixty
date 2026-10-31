import XCTest

/// Lane F16 review captures that need the running app: the link reset alert, and the synced
/// event as Apple Calendar shows it (F16.6). Runs on the shared simulator only through
/// `make -C apps/ios captures`; the PNGs land in design/reviews/2026-10-31-mobile-wave3-guild/build/ios.
@MainActor
final class GuildCaptureTests: XCTestCase {
    private static let out = URL(fileURLWithPath: #filePath)
        .deletingLastPathComponent() // UITests
        .deletingLastPathComponent() // apps/ios
        .deletingLastPathComponent() // apps
        .deletingLastPathComponent() // repo root
        .appending(path: "design/reviews/2026-10-31-mobile-wave3-guild/build/ios")
    private let springboard = XCUIApplication(bundleIdentifier: "com.apple.springboard")
    /// The snapshot goldens' frozen instant (AppTests RaidTime.now), so countdowns match them.
    private static let fixtureNow = ["-FSNow", "2026-12-17T22:30:00Z"]

    override func setUp() async throws {
        continueAfterFailure = false
        try FileManager.default.createDirectory(at: Self.out, withIntermediateDirectories: true)
    }

    private func save(_ name: String) throws {
        try XCUIScreen.main.screenshot().pngRepresentation.write(to: Self.out.appending(path: name))
    }

    /// The app on You › Calendar.
    private func openCalendarScreen() -> XCUIApplication {
        let app = XCUIApplication()
        app.launchArguments = ["-FSTab", "you"] + Self.fixtureNow
        app.launch()
        let row = app.buttons.containing(NSPredicate(format: "label BEGINSWITH %@", "Calendar")).firstMatch
        if !row.waitForExistence(timeout: 20) {
            try? save("debug-you-tab.png")
            XCTFail("no Calendar row")
        }
        row.tap()
        XCTAssertTrue(app.staticTexts.containing(NSPredicate(format: "label BEGINSWITH %@", "Calendar link")).firstMatch.waitForExistence(timeout: 10))
        return app
    }

    func testResetAlert() throws {
        let app = openCalendarScreen()
        let reset = app.buttons["Reset link"].firstMatch
        XCTAssertTrue(reset.waitForExistence(timeout: 10))
        for _ in 0..<4 where !reset.isHittable { app.swipeUp() }
        reset.tap()
        XCTAssertTrue(app.alerts.firstMatch.waitForExistence(timeout: 5))
        try save("f16.4-reset-alert.png")
        app.alerts.firstMatch.buttons["Cancel"].tap()
    }

    /// Turns the switch on (allowing full access), then finds the synced raid in Apple Calendar.
    func testEventInAppleCalendar() throws {
        let app = openCalendarScreen()
        let toggle = app.switches.firstMatch
        XCTAssertTrue(toggle.waitForExistence(timeout: 5))
        if (toggle.value as? String) != "1" {
            toggle.switches.firstMatch.exists ? toggle.switches.firstMatch.tap() : toggle.tap()
            let alert = springboard.alerts.firstMatch
            if alert.waitForExistence(timeout: 6) {
                for label in ["Allow Full Access", "Allow", "OK"] where alert.buttons[label].exists {
                    alert.buttons[label].tap()
                    break
                }
            }
        }
        sleep(4)
        try save("f16.4-calendar-on-app.png")

        let calendar = XCUIApplication(bundleIdentifier: "com.apple.mobilecal")
        calendar.launch()
        for label in ["Continue", "Not Now", "Done"] where calendar.buttons[label].waitForExistence(timeout: 2) {
            calendar.buttons[label].tap()
        }
        let search = calendar.buttons["Search"].firstMatch
        XCTAssertTrue(search.waitForExistence(timeout: 10))
        search.tap()
        calendar.typeText("Onyxia")
        let hit = calendar.staticTexts.containing(NSPredicate(format: "label BEGINSWITH %@", "Onyxia's Lair ·")).firstMatch
        XCTAssertTrue(hit.waitForExistence(timeout: 10))
        hit.tap()
        sleep(2)
        try save("f16.6-event.png")
    }

    private func begins(_ text: String) -> NSPredicate { NSPredicate(format: "label BEGINSWITH %@", text) }

    /// The app on the Raids tab, pushed to one of the guild rows.
    private func openRaidsRow(_ title: String) -> XCUIApplication {
        let app = XCUIApplication()
        app.launchArguments = ["-FSTab", "raids"] + Self.fixtureNow
        app.launch()
        let row = app.buttons.containing(begins(title)).firstMatch
        XCTAssertTrue(row.waitForExistence(timeout: 20))
        sleep(3)
        return app
    }

    /// F16.7 with the tab bar, then the roster with the tab bar, its sort menu, and its foot
    /// after scrolling all 38 rows.
    func testRaidsTabAndRosterInApp() throws {
        let app = openRaidsRow("Roster")
        try save("f16.7-raids-tab-app.png")
        app.buttons.containing(begins("Roster")).firstMatch.tap()
        XCTAssertTrue(app.staticTexts["ROSTER"].waitForExistence(timeout: 10))
        sleep(4)
        try save("f16.1-roster-app.png")
        let sort = app.buttons.containing(begins("Sort:")).firstMatch
        XCTAssertTrue(sort.waitForExistence(timeout: 5))
        sort.tap()
        sleep(1)
        try save("f16.1-sort-menu.png")
        app.buttons["Attendance"].firstMatch.tap()
        sleep(1)
        let footer = app.staticTexts.containing(begins("Attendance is the last")).firstMatch
        for _ in 0..<20 where !(footer.exists && footer.isHittable) { app.swipeUp(velocity: .fast) }
        sleep(2)
        try save("f16.1-roster-scrolled-foot.png")
    }

    func testHistoryInApp() throws {
        let app = openRaidsRow("Raid history")
        app.buttons.containing(begins("Raid history")).firstMatch.tap()
        XCTAssertTrue(app.staticTexts["RAID HISTORY"].waitForExistence(timeout: 10))
        sleep(2)
        try save("f16.3-history-app.png")
    }
}
