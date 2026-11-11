import XCTest

/// The Guild tab (spec 2026-11-10-guild-tab.md) on the fixture app, every interaction stepped through and
/// captured into design/reviews/2026-11-10-guild-tab/ios/<appearance>: in a guild under 60 (Gorvash, 58),
/// the officer's tab (`-FSRaidFixture officer`), no guild under 60 (Yellfury, 31), and a raid link that
/// opens the raid inside the Guild tab. The host sets the simulator's appearance and names it with
/// `TEST_RUNNER_FS_APPEARANCE` (`dark`, `light`). Runs on a simulator only (the Captures scheme).
@MainActor
final class GuildTabCaptureTests: XCTestCase {
    private static var out: URL {
        let appearance = ProcessInfo.processInfo.environment["FS_APPEARANCE"] ?? "dark"
        return URL(fileURLWithPath: #filePath)
            .deletingLastPathComponent() // UITests
            .deletingLastPathComponent() // apps/ios
            .deletingLastPathComponent() // apps
            .deletingLastPathComponent() // repo root
            .appending(path: "design/reviews/2026-11-10-guild-tab/ios/\(appearance)")
    }

    private static let settles: UInt32 = 4
    private static let appears: TimeInterval = 8
    private static let officer = ["-FSRaidFixture", "officer", "-FSNow", "2026-12-17T22:30:00Z"]

    override func setUp() async throws {
        continueAfterFailure = true
        try FileManager.default.createDirectory(at: Self.out, withIntermediateDirectories: true)
    }

    private func save(_ name: String) {
        try? XCUIScreen.main.screenshot().pngRepresentation.write(to: Self.out.appending(path: "\(name).png"))
    }

    private func label(_ app: XCUIApplication, _ text: String) -> XCUIElement {
        app.descendants(matching: .any).matching(NSPredicate(format: "label ==[c] %@ OR label BEGINSWITH[c] %@", text, text)).firstMatch
    }

    /// Scrolls until `element` is on screen; false when it never shows.
    @discardableResult
    private func reveal(_ app: XCUIApplication, _ element: XCUIElement, swipes: Int = 8) -> Bool {
        for _ in 0 ..< swipes where !(element.exists && element.isHittable) { app.swipeUp(velocity: .slow) }
        return element.exists && element.isHittable
    }

    /// Taps `element`, captures at once, after 400 ms and settled.
    private func tapSteps(_ element: XCUIElement, _ name: String) {
        element.tap()
        save("\(name)-t0")
        usleep(400_000)
        save("\(name)-t400")
        sleep(2)
        save("\(name)-settled")
    }

    private func backToRoot(_ app: XCUIApplication) {
        let back = app.navigationBars.buttons.firstMatch
        if back.waitForExistence(timeout: 3) { back.tap() }
        sleep(1)
    }

    /// Back from Safari (a site link), over the system's back-to-app control.
    private func returnFromSafari(_ app: XCUIApplication) {
        app.activate()
        sleep(2)
    }

    // MARK: In a guild, under 60

    func testInGuildUnder60() {
        let app = CaptureLaunch.signedIn(tab: "guild")
        sleep(Self.settles)
        XCTAssertTrue(app.tabBars.buttons["Guild"].isSelected, "the Guild tab is not the selected tab")
        save("leveler-01-top")
        // The guild leads and the raid head answers as at 60, with one quiet line (never a block).
        XCTAssertTrue(app.descendants(matching: .any)["guild-raids-open"].waitForExistence(timeout: Self.appears), "no Raids open line")
        let late = app.buttons.matching(NSPredicate(format: "label ==[c] %@", "Late")).firstMatch
        if late.waitForExistence(timeout: Self.appears) { tapSteps(late, "leveler-02-answer-late") }
        let answerIn = app.buttons.matching(NSPredicate(format: "label ==[c] %@", "In")).firstMatch
        if answerIn.exists { tapSteps(answerIn, "leveler-03-answer-in") }
        let roster = app.buttons.matching(NSPredicate(format: "label BEGINSWITH %@", "Roster")).firstMatch
        if reveal(app, roster) {
            save("leveler-04-roster-section")
            tapSteps(roster, "leveler-05-roster-page")
            backToRoot(app)
        }
        let progression = label(app, "Progression")
        if reveal(app, progression) { save("leveler-06-progression") }
        let calendar = label(app, "Calendar")
        if reveal(app, calendar) {
            tapSteps(calendar, "leveler-07-calendar")
            backToRoot(app)
        }
        app.swipeDown(velocity: .fast)
        app.swipeDown(velocity: .fast)
        sleep(2)
        save("leveler-08-pull-refresh")
    }

    // MARK: Officer

    func testOfficer() {
        let app = CaptureLaunch.signedIn(tab: "guild", extra: Self.officer)
        sleep(Self.settles)
        save("officer-01-top")
        let applications = app.descendants(matching: .any)["guild-officer-applications"]
        XCTAssertTrue(reveal(app, applications), "no Applications row")
        save("officer-02-officer-rows")
        tapSteps(applications, "officer-03-applications")
        backToRoot(app)
        let setup = app.descendants(matching: .any)["guild-officer-setup"]
        if reveal(app, setup) {
            tapSteps(setup, "officer-04-setup-site")
            returnFromSafari(app)
        }
        let plan = app.descendants(matching: .any)["guild-officer-plan"]
        if reveal(app, plan) {
            tapSteps(plan, "officer-05-plan-site")
            returnFromSafari(app)
        }
        app.swipeUp(velocity: .slow)
        sleep(1)
        save("officer-06-scrolled")
    }

    // MARK: No guild, under 60

    func testNoGuildUnder60() {
        let app = CaptureLaunch.signedIn(tab: "guild")
        sleep(Self.settles)
        CaptureLaunch.switchCharacter(app, to: "Yellfury")
        sleep(Self.settles)
        save("noguild-01-top")
        XCTAssertTrue(label(app, "Find your guild").waitForExistence(timeout: Self.appears), "no Find your guild section")
        let search = app.descendants(matching: .any)["guild-search"]
        if reveal(app, search) {
            tapSteps(search, "noguild-02-search-site")
            returnFromSafari(app)
        }
        let claim = app.descendants(matching: .any)["guild-claim"]
        if reveal(app, claim) {
            tapSteps(claim, "noguild-03-claim-site")
            returnFromSafari(app)
        }
    }

    // MARK: The deep link

    /// A raid link (a push's tap-through) opens the raid page inside the Guild tab; back returns to the tab's root.
    func testRaidLinkOpensInsideGuild() {
        let app = CaptureLaunch.signedIn(tab: "home", extra: ["-FSOpenLink", "https://foreversixty.gg/app/raid/2/31"])
        sleep(Self.settles)
        save("link-01-raid-page")
        XCTAssertTrue(app.tabBars.buttons["Guild"].isSelected, "the raid link did not land on the Guild tab")
        XCTAssertTrue(label(app, "Onyxia").waitForExistence(timeout: Self.appears), "the raid page did not open")
        backToRoot(app)
        save("link-02-back-to-guild-root")
        XCTAssertTrue(app.tabBars.buttons["Guild"].isSelected)
    }
}
