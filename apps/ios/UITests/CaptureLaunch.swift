import XCTest

/// The review captures' shared launch and shell moves (flows spec 2026-11-08: four tabs, the
/// profile menu). `signedIn` signs in on the fixtures once (Activity shows `Sign in with
/// Battle.net` signed out; fixture mode signs in at once), then relaunches on the asked tab with
/// the asked menu row pushed through the debug `-FSTab` and `-FSMenu` arguments.
@MainActor
enum CaptureLaunch {
    /// The tab Activity's signed-out state shows the sign-in button on.
    private static let signInTab = "activity"
    private static let signInButton = NSPredicate(format: "label CONTAINS[c] 'Battle.net'")
    /// The profile button's label (`menu.open`, or `menu.a11y.open` should the shell move to it).
    private static let profileButton = NSPredicate(format: "label IN %@", ["Open menu", "Open the menu"])
    private static let signInAppears: TimeInterval = 8
    private static let signInCompletes: TimeInterval = 20
    private static let shellLoads: TimeInterval = 15
    /// How long a character switch takes to re-answer the tab in place.
    private static let switchSettles: UInt32 = 3

    /// The app signed in on the fixtures, on `tab` (an AppTab raw value) with the `menu` row (a
    /// MenuDestination raw value) pushed, plus `extra` launch arguments.
    static func signedIn(tab: String = "home", menu: String? = nil, extra: [String] = []) -> XCUIApplication {
        let app = XCUIApplication()
        signIn(app, extra: extra)
        app.launchArguments = ["-FSTab", tab] + (menu.map { ["-FSMenu", $0] } ?? []) + extra
        app.launch()
        return app
    }

    /// Opens the profile menu from the current tab root.
    static func openMenu(_ app: XCUIApplication) {
        let button = app.buttons.matching(profileButton).firstMatch
        XCTAssertTrue(button.waitForExistence(timeout: shellLoads), "no profile button on the tab root")
        button.tap()
        XCTAssertTrue(app.staticTexts["Switch character"].waitForExistence(timeout: 5), "the profile menu did not open")
    }

    /// Picks `name` in the open menu's Switch character list.
    static func pick(_ app: XCUIApplication, _ name: String) {
        let row = app.buttons.matching(NSPredicate(format: "label CONTAINS %@", name)).firstMatch
        XCTAssertTrue(row.waitForExistence(timeout: 5), "no \(name) row in the profile menu")
        row.tap()
        sleep(switchSettles)
    }

    /// Opens the menu and switches to `name`.
    static func switchCharacter(_ app: XCUIApplication, to name: String) {
        openMenu(app)
        pick(app, name)
    }

    /// Launches on Activity and taps sign in when it shows (a kept session shows none), then
    /// waits for the button to go.
    private static func signIn(_ app: XCUIApplication, extra: [String]) {
        app.launchArguments = ["-FSTab", signInTab] + extra
        app.launch()
        let button = app.buttons.matching(signInButton).firstMatch
        if button.waitForExistence(timeout: signInAppears) {
            button.tap()
            let gone = XCTNSPredicateExpectation(predicate: NSPredicate(format: "exists == false"), object: button)
            XCTAssertEqual(XCTWaiter().wait(for: [gone], timeout: signInCompletes), .completed, "sign in did not finish")
        }
        app.terminate()
    }
}
