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

    // MARK: Round 4 (i2): the officer's Raid tab, the drawer in flight, the push landings

    /// The officer's Raid tab on the fixture clock (the raid ahead, four raiders not synced for eight
    /// days): Needs you, the not-synced line and Covered, then the not-synced tap's filtered roster.
    func testOfficerRaidTab() {
        let app = CaptureLaunch.signedIn(tab: "raid", extra: Self.officer)
        sleep(Self.settles)
        save("r4-officer-raid-01-top")
        for step in 2 ... 4 {
            app.swipeUp(velocity: .slow)
            sleep(1)
            save("r4-officer-raid-0\(step)-scrolled")
        }
        app.swipeDown(velocity: .fast)
        app.swipeDown(velocity: .fast)
        sleep(1)
        let line = app.descendants(matching: .any).matching(NSPredicate(format: "label BEGINSWITH %@", "Not synced this week")).firstMatch
        for _ in 0 ..< 4 where !line.isHittable { app.swipeUp(velocity: .slow) }
        guard line.waitForExistence(timeout: Self.appears), line.isHittable else { return XCTFail("no not-synced line") }
        save("r4-officer-raid-05-not-synced-line")
        line.tap()
        save("r4-officer-raid-06-roster-t0")
        usleep(400_000)
        save("r4-officer-raid-06-roster-t400")
        sleep(2)
        save("r4-officer-raid-06-roster-settled")
        // Round 5, row 2: the filter chip's clear control brings the whole roster back.
        let clear = app.buttons.matching(NSPredicate(format: "label ==[c] %@", "Show the whole roster")).firstMatch
        guard clear.waitForExistence(timeout: Self.appears) else { return XCTFail("no clear control on the filter chip") }
        clear.tap()
        sleep(1)
        save("r5-officer-raid-07-roster-cleared")
    }

    /// The drawer's Refresh: the `Refreshing…` frame inside 200 ms, then the new stamp.
    func testMenuRefreshInFlight() {
        // The fixture answers at once; the debug hold keeps the call out long enough to see.
        let app = CaptureLaunch.signedIn(tab: "home", extra: ["-FSRefreshHold", "2"])
        sleep(Self.settles)
        CaptureLaunch.openMenu(app)
        let refresh = app.buttons.matching(NSPredicate(format: "label ==[c] %@", "Refresh")).firstMatch
        guard refresh.waitForExistence(timeout: Self.appears) else { return XCTFail("no Refresh link") }
        save("r4-menu-refresh-01-before")
        refresh.tap()
        save("r4-menu-refresh-02-t0")
        usleep(150_000)
        save("r4-menu-refresh-03-t150")
        sleep(2)
        save("r4-menu-refresh-04-settled")
    }

    /// The raid push's tap-through: the banner over the app (on Activity) lands on the raid.
    func testRaidPushLanding() {
        let app = CaptureLaunch.signedIn(tab: "activity", extra: ["-FSPostRaidPush"])
        allowNotifications()
        guard let note = banner("Onyxia's Lair") else { return XCTFail("no raid notification") }
        save("r4-push-raid-01-delivered")
        note.coordinate(withNormalizedOffset: CGVector(dx: 0.5, dy: 0.5)).tap()
        sleep(Self.settles)
        save("r4-push-raid-02-landing")
        _ = app
    }

    /// The officer pushes: raid.short with Mention and the shade after it ran in place; then
    /// guild.applicant with Accept and Decline, and the shade after Accept.
    /// The pushes' actions answered from the banner, each with its result: the raid push's Late (the
    /// server's confirmation replaces it), Mention on raid.short, Accept and Decline on guild.applicant.
    func testOfficerPushes() {
        pushAction(["-FSPostRaidPush"], title: "Onyxia's Lair", action: "Late", name: "r4-push-raid-late")
        let officer = Self.officer + ["-FSPostPush"]
        // Round 5, row 3: each result replaces the push and shows as its banner.
        pushAction(
            officer + ["push.raid-short"], title: "Onyxia's Lair", action: "Mention in Discord", name: "r4-push-short",
            result: "Mentioned 3 healers in Discord."
        )
        pushAction(
            officer + ["push.guild-applicant"], title: "New applicant", action: "Accept", name: "r4-push-applicant-accept",
            result: "Application accepted"
        )
        pushAction(
            officer + ["push.guild-applicant"], title: "New applicant", action: "Decline", name: "r4-push-applicant-decline",
            result: "Application declined"
        )
    }

    private static let officer = ["-FSRaidFixture", "officer", "-FSNow", "2026-12-17T22:30:00Z"]

    // MARK: Round 6: the drawer's Notifications value, Activity, the filtered roster, the failure lines, stale Gear

    /// Eleven days past the fixture sync: past the seven-day stale age.
    private static let staleNow = ["-FSNow", "2026-11-20T18:00:00Z"]

    func testRound6Drawer() {
        let app = CaptureLaunch.signedIn(tab: "home")
        sleep(Self.settles)
        CaptureLaunch.openMenu(app)
        sleep(2)
        save("r6-menu-01-open")
        app.swipeUp()
        sleep(1)
        save("r6-menu-02-account")
    }

    func testRound6Activity() {
        let app = CaptureLaunch.signedIn(tab: "activity")
        sleep(Self.settles)
        save("r6-activity-01-top")
        _ = app
    }

    /// The not-synced tap's roster: the filter chip, and the plate's one figure.
    func testRound6FilteredRoster() {
        let app = CaptureLaunch.signedIn(tab: "raid", extra: Self.officer)
        sleep(Self.settles)
        let line = app.descendants(matching: .any).matching(NSPredicate(format: "label BEGINSWITH %@", "Not synced this week")).firstMatch
        for _ in 0 ..< 4 where !line.isHittable { app.swipeUp(velocity: .slow) }
        guard line.waitForExistence(timeout: Self.appears), line.isHittable else { return XCTFail("no not-synced line") }
        line.tap()
        sleep(2)
        save("r6-roster-01-filtered")
    }

    // MARK: Round 7: Gear in the house format and its sheet's Where line, the push failures, Notifications

    func testRound7Home() {
        let app = CaptureLaunch.signedIn(tab: "home")
        sleep(Self.settles)
        save("r7-home-01-top")
        _ = app
    }

    // MARK: Round 8: the sheet's source row, the Decline and Mention failures, the drawer from its top

    /// The Gear sheet's source row: the head pick crafted at a skill with its runner-up a boss drop
    /// (Ancient Treasure, Uldaman), the off-hand crafted upgrade, and the back's world-drop upgrade.
    func testRound8GearSheets() { gearSheets(round: "r8") }

    /// Round 9: the runner-up row names its source only when it differs, `Same DPS`, PvP without a faction.
    func testRound9GearSheets() { gearSheets(round: "r9") }

    private func gearSheets(round: String) {
        let app = CaptureLaunch.signedIn(tab: "gear")
        sleep(Self.settles)
        for (slot, name) in [("Head", "crafted-pvp-runner-up"), ("Off hand", "crafted-upgrade"), ("Back", "upgrade"), ("Legs", "vendor-upgrade")] {
            if tapSteps(app, slot, "\(round)-gear-sheet-\(name)") {
                app.swipeDown(velocity: .fast)
                sleep(1)
            }
        }
    }

    func testRound8PushFailures() {
        let failing = Self.officer + ["-FSFailOfficerPush", "-FSPostPush"]
        pushAction(
            failing + ["push.raid-short"], title: "Onyxia's Lair", action: "Mention in Discord", name: "r8-push-short-fail",
            result: "Could not mention"
        )
        pushAction(
            failing + ["push.guild-applicant"], title: "New applicant", action: "Decline", name: "r8-push-applicant-decline-fail",
            result: "Could not decline"
        )
    }

    /// The drawer's Notifications value from the top of the drawer: the header and the row in one frame.
    func testRound8DrawerNotifications() {
        for (state, extra) in [("some-off", [String]()), ("all-on", ["-FSNotifications", "on"]), ("all-off", ["-FSNotifications", "off"])] {
            let app = CaptureLaunch.signedIn(tab: "home", extra: extra)
            sleep(Self.settles)
            CaptureLaunch.openMenu(app)
            sleep(2)
            save("r8-menu-notifications-\(state)")
            app.terminate()
        }
    }

    /// Gear on a clock past the stale age: the synced line in the warning tone.
    func testRound6StaleGear() {
        let app = CaptureLaunch.signedIn(tab: "gear", extra: Self.staleNow)
        sleep(Self.settles)
        save("r6-gear-01-stale")
        _ = app
    }

    private func pushAction(_ extra: [String], title: String, action: String, name: String, result: String? = nil) {
        let app = CaptureLaunch.signedIn(tab: "home", extra: extra)
        allowNotifications()
        guard let note = banner(title) else { return XCTFail("\(name): no notification") }
        save("\(name)-01-delivered")
        note.coordinate(withNormalizedOffset: CGVector(dx: 0.5, dy: 0.5)).press(forDuration: 1.5)
        let button = springboard.buttons[action]
        guard button.waitForExistence(timeout: 5) else { return XCTFail("\(name): no \(action)") }
        save("\(name)-02-actions")
        button.tap()
        if let result {
            // The banner that replaced the push, once it is up.
            if banner(result) == nil { XCTFail("\(name): no \(result) banner") }
        } else {
            usleep(1_500_000)
        }
        save("\(name)-03-result")
        sleep(3)
        save("\(name)-04-settled")
        XCTAssertEqual(app.state, .runningForeground, "\(name): the app kept running after the action")
        app.terminate()
    }

    private func allowNotifications() {
        let alert = springboard.alerts.firstMatch
        if alert.waitForExistence(timeout: 4), alert.buttons["Allow"].exists { alert.buttons["Allow"].tap() }
        sleep(1)
    }

    /// The banner the app presents in the foreground (`willPresent`), by the start of its title.
    private func banner(_ title: String) -> XCUIElement? {
        let note = springboard.staticTexts.matching(NSPredicate(format: "label BEGINSWITH %@", title)).firstMatch
        return note.waitForExistence(timeout: 10) ? note : nil
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
