import XCTest

/// Review captures for lanes M1i to M5i (round 5 on iOS): the four tab roots in the fixture app with
/// their real bars (the band clear over the art at rest, solid once scrolled) and the profile drawer.
/// Runs on the shared simulator only (`make -C apps/ios captures`); the PNGs land in
/// design/reviews/2026-11-10-restyle-app. Not part of `make test`.
@MainActor
final class RestyleCaptureTests: XCTestCase {
    private static let out = URL(fileURLWithPath: #filePath)
        .deletingLastPathComponent() // UITests
        .deletingLastPathComponent() // apps/ios
        .deletingLastPathComponent() // apps
        .deletingLastPathComponent() // repo root
        .appending(path: "design/reviews/2026-11-10-restyle-app")
    /// How long a tab takes to draw its fixture answers.
    private static let settles: UInt32 = 5

    override func setUp() async throws {
        continueAfterFailure = false
        try FileManager.default.createDirectory(at: Self.out, withIntermediateDirectories: true)
    }

    private func save(_ name: String) throws {
        try XCUIScreen.main.screenshot().pngRepresentation.write(to: Self.out.appending(path: "\(name)-ios.png"))
    }

    func testTabRootsAndDrawer() throws {
        let app = CaptureLaunch.signedIn(tab: "home")
        sleep(Self.settles)
        try save("home")
        app.swipeUp()
        sleep(2)
        try save("home-scrolled")
        for tab in ["Raid", "Gear", "Activity"] {
            app.tabBars.buttons[tab].tap()
            sleep(Self.settles)
            try save(tab.lowercased())
        }
        app.tabBars.buttons["Home"].tap()
        sleep(2)
        CaptureLaunch.openMenu(app)
        sleep(1)
        try save("menu")
    }
}
