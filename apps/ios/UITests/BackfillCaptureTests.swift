import XCTest

/// The mobile review backfill's step captures (design/reviews/2026-11-10-mobile-backfill/ios): every
/// interaction of the four tabs, the profile menu and the raid notification, before, during and after,
/// on the fixture app. Runs on a simulator only (`make -C apps/ios captures` or `-only-testing`); not
/// part of `make test`. A step whose element is missing is recorded as a failure and the walk goes on,
/// so one missing control never hides the rest.
@MainActor
final class BackfillCaptureTests: XCTestCase {
    private static let out = URL(fileURLWithPath: #filePath)
        .deletingLastPathComponent() // UITests
        .deletingLastPathComponent() // apps/ios
        .deletingLastPathComponent() // apps
        .deletingLastPathComponent() // repo root
        .appending(path: "design/reviews/2026-11-10-mobile-backfill/ios")
    /// How long a tab takes to draw its fixture answers.
    private static let settles: UInt32 = 4
    private static let appears: TimeInterval = 8
    private let springboard = XCUIApplication(bundleIdentifier: "com.apple.springboard")

    override func setUp() async throws {
        continueAfterFailure = true
        try FileManager.default.createDirectory(at: Self.out, withIntermediateDirectories: true)
    }

    private func save(_ name: String) {
        try? XCUIScreen.main.screenshot().pngRepresentation.write(to: Self.out.appending(path: "\(name).png"))
    }

    /// The button labelled exactly `text`, else the first element of any kind whose label is or
    /// begins with it (`Late` must not find `Plate Leggings`).
    private func element(_ app: XCUIApplication, _ text: String) -> XCUIElement {
        let button = app.buttons.matching(NSPredicate(format: "label ==[c] %@", text)).firstMatch
        if button.exists { return button }
        return app.descendants(matching: .any).matching(NSPredicate(format: "label ==[c] %@ OR label BEGINSWITH[c] %@", text, text)).firstMatch
    }

    /// Taps `text` and captures at once, after 400 ms and settled. False when it is not on screen.
    @discardableResult
    private func tapSteps(_ app: XCUIApplication, _ text: String, _ name: String) -> Bool {
        let target = element(app, text)
        guard target.waitForExistence(timeout: Self.appears) else {
            XCTFail("\(name): no \(text) on screen")
            return false
        }
        target.tap()
        save("\(name)-t0")
        usleep(400_000)
        save("\(name)-t400")
        sleep(2)
        save("\(name)-settled")
        return true
    }

    /// A long touch on `text`, captured mid-press: the pressed state.
    private func pressed(_ app: XCUIApplication, _ text: String, _ name: String) {
        let target = element(app, text)
        guard target.waitForExistence(timeout: Self.appears) else { return XCTFail("\(name): no \(text)") }
        let start = target.coordinate(withNormalizedOffset: CGVector(dx: 0.5, dy: 0.5))
        // Dragging to a point far off cancels the press, so nothing is chosen.
        let away = start.withOffset(CGVector(dx: 0, dy: 600))
        DispatchQueue.main.asyncAfter(deadline: .now() + 0.6) { self.save("\(name)-pressed") }
        start.press(forDuration: 1.0, thenDragTo: away)
    }

    /// The four tab roots and the menu, named by `TEST_RUNNER_FS_CAPTURE_PREFIX` (`light`, `ax5`): the
    /// host sets the simulator's appearance or text size before the run.
    func testTabRootsUnderTheSystemSetting() {
        guard let prefix = ProcessInfo.processInfo.environment["FS_CAPTURE_PREFIX"] else { return }
        let app = CaptureLaunch.signedIn(tab: "home")
        sleep(Self.settles)
        save("\(prefix)-home")
        for tab in ["Raid", "Gear", "Activity"] {
            app.tabBars.buttons[tab].tap()
            sleep(Self.settles)
            save("\(prefix)-\(tab.lowercased())")
        }
        app.tabBars.buttons["Home"].tap()
        sleep(1)
        CaptureLaunch.openMenu(app)
        save("\(prefix)-menu")
    }

    func testHomeAndSignup() {
        let app = CaptureLaunch.signedIn(tab: "home")
        sleep(Self.settles)
        save("home-01-top")
        app.swipeUp()
        sleep(1)
        save("home-02-scrolled")
        app.swipeDown()
        app.swipeDown()
        sleep(2)
        save("home-03-pull-refresh")
        pressed(app, "Late", "home-signup-01-late")
        tapSteps(app, "Late", "home-signup-02-late")
        tapSteps(app, "30 min", "home-signup-03-late-30")
        tapSteps(app, "Bench", "home-signup-04-bench")
        tapSteps(app, "In", "home-signup-05-in")
        tapSteps(app, "Gear", "home-06-gear-link")
    }

    func testRaidTab() {
        let app = CaptureLaunch.signedIn(tab: "raid")
        sleep(Self.settles)
        save("raid-01-top")
        tapSteps(app, "Late", "raid-signup-02-late")
        tapSteps(app, "30 min", "raid-signup-03-late-30")
        tapSteps(app, "Out", "raid-signup-04-out")
        tapSteps(app, "In", "raid-signup-05-in")
        app.swipeUp()
        sleep(1)
        save("raid-06-scrolled")
        if tapSteps(app, "Roster", "raid-07-roster") {
            app.navigationBars.buttons.firstMatch.tap()
            sleep(2)
            save("raid-08-back-to-root")
        }
        app.swipeDown()
        app.swipeDown()
        sleep(2)
        save("raid-09-pull-refresh")
    }

    func testGearTab() {
        let app = CaptureLaunch.signedIn(tab: "gear")
        sleep(Self.settles)
        save("gear-01-top")
        app.swipeUp()
        sleep(1)
        save("gear-02-scrolled")
        app.swipeDown()
        sleep(1)
        if tapSteps(app, "Head", "gear-03-item-sheet") {
            app.swipeDown(velocity: .fast)
            sleep(1)
            save("gear-04-sheet-dismissed")
        }
        if tapSteps(app, "Change", "gear-05-change") {
            app.swipeDown(velocity: .fast)
            sleep(1)
        }
    }

    func testActivityTab() {
        let app = CaptureLaunch.signedIn(tab: "activity")
        sleep(Self.settles)
        save("activity-01-top")
        for (index, filter) in ["Raids", "Gear and sims", "Logs", "All"].enumerated() {
            tapSteps(app, filter, "activity-0\(index + 2)-filter-\(filter.lowercased().replacingOccurrences(of: " ", with: "-"))")
        }
        tapSteps(app, "Mark all read", "activity-06-mark-read")
        if !tapSteps(app, "Show older", "activity-07-older") { tapSteps(app, "More", "activity-07-more") }
        app.swipeDown()
        app.swipeDown()
        sleep(2)
        save("activity-08-pull-refresh")
    }

    func testMenu() {
        let app = CaptureLaunch.signedIn(tab: "home")
        sleep(Self.settles)
        CaptureLaunch.openMenu(app)
        save("menu-01-open")
        tapSteps(app, "Refresh", "menu-02-refresh")
        app.swipeUp()
        sleep(1)
        save("menu-03-scrolled")
        app.swipeDown()
        CaptureLaunch.pick(app, "Yellfury")
        save("menu-04-switched-home")
        for tab in ["Raid", "Gear", "Activity"] {
            app.tabBars.buttons[tab].tap()
            sleep(Self.settles)
            save("menu-05-switched-\(tab.lowercased())")
        }
        app.tabBars.buttons["Home"].tap()
        CaptureLaunch.switchCharacter(app, to: "Gorvash")
        CaptureLaunch.openMenu(app)
        tapSteps(app, "Planner", "menu-06-planner")
        app.tabBars.buttons["Raid"].tap()
        sleep(2)
        save("menu-07-raid-after-planner")
        app.tabBars.buttons["Home"].tap()
        sleep(2)
        save("menu-08-home-after-planner")
    }

    func testRaidNotificationAndLinks() {
        let app = XCUIApplication()
        app.launchArguments = ["-FSPostRaidPush"]
        app.launch()
        let alert = springboard.alerts.firstMatch
        if alert.waitForExistence(timeout: 4), alert.buttons["Allow"].exists { alert.buttons["Allow"].tap() }
        sleep(1)
        XCUIDevice.shared.press(.home)
        sleep(4)
        let top = springboard.coordinate(withNormalizedOffset: CGVector(dx: 0.3, dy: 0.01))
        top.press(forDuration: 0.1, thenDragTo: springboard.coordinate(withNormalizedOffset: CGVector(dx: 0.3, dy: 0.7)))
        let note = springboard.staticTexts["Onyxia's Lair"].firstMatch
        guard note.waitForExistence(timeout: 10) else { return XCTFail("no raid notification") }
        save("notif-01-delivered")
        note.press(forDuration: 1.5)
        let late = springboard.buttons["Late"]
        if late.waitForExistence(timeout: 5) {
            save("notif-02-actions")
            late.tap()
            sleep(6)
            save("notif-03-late-answered")
        }
        let answered = springboard.staticTexts.matching(NSPredicate(format: "label CONTAINS[c] %@", "Onyxia")).firstMatch
        if answered.waitForExistence(timeout: 5) {
            answered.tap()
            sleep(Self.settles)
            save("notif-04-tap-through")
        }
        app.terminate()
        let linked = CaptureLaunch.signedIn(tab: "home", extra: ["-FSOpenLink", "https://foreversixty.gg/app/raid/2/31"])
        sleep(Self.settles)
        save("link-01-app-raid")
        _ = linked
    }
}
