import XCTest

/// Lane F13 review captures that need the running app: the Reset alert is a system alert, which a
/// view snapshot cannot draw. Opens the planner on the fixtures through a debug `-FSOpenLink`
/// (a warrior guide build), presses Reset and saves the alert. Runs on the shared simulator only,
/// through the ForeverSixtyCaptures scheme; not part of `make test`.
@MainActor
final class PlannerCaptureTests: XCTestCase {
    private static let out = URL(fileURLWithPath: #filePath)
        .deletingLastPathComponent() // UITests
        .deletingLastPathComponent() // apps/ios
        .deletingLastPathComponent() // apps
        .deletingLastPathComponent() // repo root
        .appending(path: "design/reviews/2026-11-01-mobile-planner/build/ios")
    private static let plannerLink =
        "https://foreversixty.gg/app/planner/warrior?code=FS1:1.60.1.70707:warrior:human:35300003/0/250533120330001351:"

    override func setUp() async throws {
        continueAfterFailure = false
        try FileManager.default.createDirectory(at: Self.out, withIntermediateDirectories: true)
    }

    func testResetAlert() throws {
        let app = XCUIApplication()
        app.launchArguments = ["-FSOpenLink", Self.plannerLink]
        app.launch()
        let reset = app.buttons["Reset"]
        XCTAssertTrue(reset.waitForExistence(timeout: 20))
        sleep(2)
        reset.tap()
        XCTAssertTrue(app.alerts.firstMatch.waitForExistence(timeout: 5))
        sleep(1)
        try XCUIScreen.main.screenshot().pngRepresentation.write(to: Self.out.appending(path: "f13-1-reset-ios.png"))
    }
}
