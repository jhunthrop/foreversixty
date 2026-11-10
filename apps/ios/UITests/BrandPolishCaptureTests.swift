import XCTest

/// Review captures for lane app-brand-polish: the four tab roots with the head fade (the band art
/// running under the status bar and the bars and fading into the ground, as the site's page head),
/// at rest and scrolled, with the system in light and in dark appearance. Runs on the shared
/// simulator only (`make -C apps/ios captures`); the PNGs land in
/// design/reviews/2026-11-10-app-brand-polish. Not part of `make test`.
@MainActor
final class BrandPolishCaptureTests: XCTestCase {
    private static let out = URL(fileURLWithPath: #filePath)
        .deletingLastPathComponent() // UITests
        .deletingLastPathComponent() // apps/ios
        .deletingLastPathComponent() // apps
        .deletingLastPathComponent() // repo root
        .appending(path: "design/reviews/2026-11-10-app-brand-polish")
    /// How long a tab takes to draw its fixture answers.
    private static let settles: UInt32 = 5

    override func setUp() async throws {
        continueAfterFailure = false
        try FileManager.default.createDirectory(at: Self.out, withIntermediateDirectories: true)
    }

    override func tearDown() async throws {
        XCUIDevice.shared.appearance = .dark
    }

    private func save(_ name: String) throws {
        try XCUIScreen.main.screenshot().pngRepresentation.write(to: Self.out.appending(path: "\(name)-ios.png"))
    }

    func testHeadFadeOnTabRoots() throws {
        for (appearance, label) in [(XCUIDevice.Appearance.dark, "dark"), (.light, "light")] {
            XCUIDevice.shared.appearance = appearance
            let app = CaptureLaunch.signedIn(tab: "home")
            sleep(Self.settles)
            try save("fade-home-\(label)")
            for tab in ["Raid", "Gear", "Activity"] {
                app.tabBars.buttons[tab].tap()
                sleep(Self.settles)
                try save("fade-\(tab.lowercased())-\(label)")
            }
            app.swipeUp()
            sleep(2)
            try save("fade-activity-scrolled-\(label)")
            app.terminate()
        }
    }
}
